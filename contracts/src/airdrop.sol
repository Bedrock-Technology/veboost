// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

import "@openzeppelin/contracts-upgradeable/utils/cryptography/MerkleProofUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "./interface/IVotingEscrowIncreasing.sol";

contract Airdrop is Initializable, AccessControlUpgradeable, PausableUpgradeable, ReentrancyGuardUpgradeable {
    using SafeERC20 for IERC20;

    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");
    bytes32 public constant OPERATOR_ROLE = keccak256("OPERATOR_ROLE");

    struct DistributionRoot {
        /// @notice The root of Merkle tree for airdrop distribution.
        bytes32 root;
        /// @notice The timestamp when this distribution becomes active.
        uint32 activatedAt;
        /// @notice The duration in seconds that this distribution remains active.
        uint32 duration;
        /// @notice The flag indicating if this distribution is disabled.
        bool disabled;
    }

    /// @notice Mapping of epoch to its distribution root data.
    mapping(uint256 => DistributionRoot) private merkleRoots;
    /// @notice Mapping of epoch and user address to claim status.
    mapping(uint256 => mapping(address => bool)) private hasClaimed;
    /// @notice Delay in timestamp (seconds) before a posted root can be claimed against.
    uint32 public activationDelay;
    /// @notice Current epoch number of the airdrop distribution.
    uint256 public currentEpoch;
    /// @notice Address of the voting escrow contract.
    address public votingEscrow;
    /// @notice Address of the BR token contract.
    address public brToken;

    /**
     * ======================================================================================
     *
     * CONSTRUCTOR
     *
     * ======================================================================================
     */
    constructor() {
        _disableInitializers();
    }

    /**
     * ======================================================================================
     *
     * ADMIN
     *
     * ======================================================================================
     */
    /**
     * @notice Initializes the airdrop contract with required parameters.
     * @dev Sets up roles and initializes core contract parameters.
     * @param _activationDelay The initial delay before claims can be made.
     * @param _votingEscrow The address of the voting escrow contract.
     * @param _brToken The address of the BR token contract.
     * @param _admin The address of the contract administrator.
     */
    function initialize(uint32 _activationDelay, address _votingEscrow, address _brToken, address _admin)
        public
        initializer
    {
        require(_votingEscrow != address(0), "SYS001");
        require(_brToken != address(0), "SYS001");
        require(_admin != address(0), "SYS001");

        __AccessControl_init();
        __Pausable_init();
        __ReentrancyGuard_init();

        _setupRole(DEFAULT_ADMIN_ROLE, _admin);
        _setupRole(PAUSER_ROLE, _admin);
        _setupRole(OPERATOR_ROLE, _admin);

        votingEscrow = _votingEscrow;
        brToken = _brToken;
        _setActivationDelay(_activationDelay);
        currentEpoch = 0;
    }

    /**
     * @notice Pauses all contract operations.
     * @dev Only callable by accounts with PAUSER_ROLE.
     */
    function pause() external onlyRole(PAUSER_ROLE) {
        _pause();
    }

    /**
     * @notice Unpauses all contract operations.
     * @dev Only callable by accounts with PAUSER_ROLE.
     */
    function unpause() external onlyRole(PAUSER_ROLE) {
        _unpause();
    }

    /**
     * @notice Sets the delay in timestamp before a posted root can be claimed against.
     * @dev Only callable by accounts with OPERATOR_ROLE.
     * @param _activationDelay The new value for activationDelay.
     */
    function setActivationDelay(uint32 _activationDelay) external onlyRole(OPERATOR_ROLE) {
        _setActivationDelay(_activationDelay);
    }

    /**
     * @notice Submits a new Merkle root and starts a new airdrop epoch.
     * @dev Only callable by accounts with OPERATOR_ROLE.
     * @param _newRoot The Merkle root of the new distribution.
     * @param _duration The duration in seconds for which this distribution is valid.
     */
    function submitMerkleRoot(bytes32 _newRoot, uint32 _duration) external onlyRole(OPERATOR_ROLE) {
        require(_duration > 0, "SYS002");
        require(_newRoot != bytes32(0), "SYS002");
        require(!_isActive(), "USR001");
        currentEpoch++;

        merkleRoots[currentEpoch] = DistributionRoot({
            root: _newRoot,
            activatedAt: uint32(block.timestamp) + activationDelay,
            duration: _duration,
            disabled: false
        });

        emit MerkleRootSubmit(currentEpoch, _newRoot, _duration, uint32(block.timestamp));
    }

    /**
     * @notice Updates the Merkle root for the current epoch.
     * @dev Only callable by accounts with OPERATOR_ROLE.
     * @param _newRoot The new Merkle root to replace the current one.
     */
    function updateMerkleRoot(bytes32 _newRoot) external onlyRole(OPERATOR_ROLE) {
        require(currentEpoch > 0, "USR002");
        require(_newRoot != bytes32(0), "USR003");
        emit MerkleRootUpdate(currentEpoch, merkleRoots[currentEpoch].root, _newRoot);
        merkleRoots[currentEpoch].root = _newRoot;
    }

    /**
     * @notice Updates the valid duration for the current epoch.
     * @dev Only callable by accounts with OPERATOR_ROLE.
     * @param _duration The new duration in seconds.
     */
    function updateDuration(uint32 _duration) external onlyRole(OPERATOR_ROLE) {
        require(currentEpoch > 0, "USR002");
        require(block.timestamp <= merkleRoots[currentEpoch].activatedAt + _duration, "USR004");
        emit ValidDurationUpdate(currentEpoch, merkleRoots[currentEpoch].duration, _duration);
        merkleRoots[currentEpoch].duration = _duration;
    }

    /**
     * @notice Sets the distribution status for the current epoch.
     * @dev Only callable by accounts with OPERATOR_ROLE.
     * @param _disabled The status to set (true = disabled, false = enabled).
     */
    function setAirdrop(bool _disabled) external onlyRole(OPERATOR_ROLE) {
        require(currentEpoch > 0, "USR002");
        DistributionRoot storage distribution = merkleRoots[currentEpoch];
        emit DistributionDisabledSet(currentEpoch, distribution.disabled, _disabled);
        distribution.disabled = _disabled;
    }

    /**
     * ======================================================================================
     *
     * INTERNAL FUNCTIONS
     *
     * ======================================================================================
     */
    /**
     * @notice Updates the activation delay for airdrop claims.
     * @dev Internal function to update the delay before claims can be made.
     * @param _activationDelay The new activation delay value in seconds.
     */
    function _setActivationDelay(uint32 _activationDelay) internal {
        emit ActivationDelaySet(activationDelay, _activationDelay);
        activationDelay = _activationDelay;
    }

    /**
     * @notice Checks if the current epoch's airdrop is active and valid.
     * @dev Returns false if: no active epoch, distribution disabled, or expired.
     * @return True if the current epoch's airdrop is valid and active.
     */
    function _isActive() internal view returns (bool) {
        if (currentEpoch == 0) return false;

        DistributionRoot memory distribution = merkleRoots[currentEpoch];
        if (distribution.disabled) return false;

        uint256 currentTime = block.timestamp;
        if (currentTime > distribution.activatedAt + distribution.duration) return false;

        return true;
    }

    /**
     * ======================================================================================
     *
     * EXTERNAL FUNCTIONS
     *
     * ======================================================================================
     */

    /**
     * @notice Claims airdrop tokens for the current epoch and locks them in VotingEscrow.
     * @dev Verifies Merkle proof and handles token transfer and locking.
     * @param _amount The amount of tokens to claim.
     * @param _proof The Merkle proof verifying the claim eligibility.
     */
    function claim(uint256 _amount, bytes32[] calldata _proof) external whenNotPaused nonReentrant {
        require(currentEpoch > 0, "USR002");
        require(!hasClaimed[currentEpoch][msg.sender], "USR005");

        DistributionRoot memory distribution = merkleRoots[currentEpoch];
        require(!distribution.disabled, "USR006");

        // Check if the distribution is within valid period.
        require(block.timestamp >= distribution.activatedAt, "USR007");
        require(block.timestamp <= distribution.activatedAt + distribution.duration, "USR008");

        // Verify Merkle proof.
        bytes32 leaf = keccak256(bytes.concat(keccak256(abi.encode(msg.sender, _amount))));
        require(MerkleProofUpgradeable.verify(_proof, distribution.root, leaf), "USR009");

        // Mark as claimed.
        hasClaimed[currentEpoch][msg.sender] = true;

        // Transfer tokens and lock in VotingEscrow.
        IERC20(brToken).safeApprove(votingEscrow, _amount);
        uint256 veNFTId = IVotingEscrowCore(votingEscrow).createLockFor(_amount, msg.sender);

        emit AirdropClaimed(currentEpoch, msg.sender, _amount, veNFTId);
    }

    /**
     * @notice Retrieves the distribution root information for a specific epoch.
     * @dev Returns the complete DistributionRoot struct.
     * @param _epoch The epoch number to query.
     * @return The DistributionRoot struct containing root, activatedAt, validDuration and disabled status.
     */
    function getDistributionRoot(uint256 _epoch) external view returns (DistributionRoot memory) {
        return merkleRoots[_epoch];
    }

    /**
     * @notice Checks if a user has claimed their airdrop for a specific epoch.
     * @dev Returns the claim status from hasClaimed mapping.
     * @param _epoch The epoch number to query.
     * @param _user The address of the user to check.
     * @return True if the user has claimed, false otherwise.
     */
    function hasUserClaimed(uint256 _epoch, address _user) external view returns (bool) {
        return hasClaimed[_epoch][_user];
    }

    /**
     * @notice Checks if the current epoch's airdrop is active and valid.
     * @dev Returns false if: no active epoch, distribution disabled, not activated yet, or expired.
     * @return True if the current epoch's airdrop is valid and active.
     */
    function isCurrentEpochActive() external view returns (bool) {
        return _isActive();
    }

    /**
     * ======================================================================================
     *
     * EVENTS
     *
     * ======================================================================================
     */
    /// @notice Emitted when a new Merkle root is submitted for a new epoch.
    event MerkleRootSubmit(uint256 indexed epoch, bytes32 root, uint32 rewardsValidTime, uint32 activatedAt);
    /// @notice Emitted when the Merkle root is updated for the current epoch.
    event MerkleRootUpdate(uint256 indexed epoch, bytes32 preRoot, bytes32 root);
    /// @notice Emitted when the valid duration is updated for the current epoch.
    event ValidDurationUpdate(uint256 indexed epoch, uint32 preValidDuration, uint32 validDuration);
    /// @notice Emitted when an airdrop is claimed by a user.
    event AirdropClaimed(uint256 indexed epoch, address indexed user, uint256 amount, uint256 veNFTId);
    /// @notice Emitted when the activation delay is updated.
    event ActivationDelaySet(uint32 oldActivationDelay, uint32 newActivationDelay);
    /// @notice Emitted when the distribution status is changed.
    event DistributionDisabledSet(uint256 indexed epoch, bool preStatus, bool status);
}
