// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

import {Test} from "forge-std/Test.sol";
import {Airdrop} from "../src/airdrop.sol";
import {IVotingEscrowCore} from "../src/interface/IVotingEscrowIncreasing.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/utils/cryptography/MerkleProof.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

///forge test --match-contract ForkBscMerkleTreeTest --match-test testClaimAirdrop --fork-url wss://bsc-rpc.publicnode.com -vvvv
contract ForkBscMerkleTreeTest is Test {
    // Contract instances
    Airdrop private implementation;
    Airdrop private airdrop;
    IERC20 private brToken;
    IVotingEscrowCore private votingEscrow;
    ProxyAdmin private proxyAdmin;
    address private admin;

    // BSC mainnet contract addresses
    address constant PROXY_ADMIN = 0xb3f925B430C60bA467F7729975D5151c8DE26698;
    address constant VOTING_ESCROW = 0x25ab4059145E7f1A5404390f0B82d94079b3e83f;
    address constant BR_TOKEN = 0xFf7d6A96ae471BbCD7713aF9CB1fEeB16cf56B41;

    // Constants for airdrop configuration
    uint32 public constant ACTIVATION_DELAY = 1 days;
    uint32 public constant VALID_DURATION = 30 days;
    
    // Test data for merkle proof verification
    address public constant claimant = address(0x0C99B08F2233b04066fe13A0A1Bf1474416fD77F);
    uint256 public constant amount = 1802977279010487416443;
    bytes32 public constant merkleRoot = 0xb615db797d417a5b966a181cf5ce9054a777b0a31b934bd762aab1dfb75a1016;
    bytes32[] public proof = [
        bytes32(0x5d40149407495c6d34d1c4bcb99390882123cccfd290efaa6b365d34d3ba2b47),
        bytes32(0x0005a2d93a6222f40a3731f70d4120dcc18b2d574be7190835660cf8a5acdb0a),
        bytes32(0x9834ab15ebefe10540539372e156a7b1c58f78f67ff8c827141d1e09e5ee785b),
        bytes32(0xfa98761f8b22395518d310267ee9a84b77b1347e19d3553522443ce9cd3173bc),
        bytes32(0x6379ddfc98ad6cfe1b5a4d6796fb05a54ce1f0ac94f05918a92770d9ff7e00f4),
        bytes32(0xf3858d951824f6d62facbed01c8c53709a3a987f76b76e5ebf6e8be70843f6b7),
        bytes32(0xad577363479f82be6260628232f1a8d49242c0b11e93bf2eea65184f85de08a7),
        bytes32(0xd8ac78a0434eba5b81602909d937c5f7b8dbcf047aa03e033fdd61f659f7b326),
        bytes32(0x2ca334f693e5037fb9793f0058861a6e2b97a4e8ff902af6952cc62ac6acea51),
        bytes32(0x05015516f5e2c84ee44e1cae08a85e139747ec3071eb43b0e2db0d1c815b6f60)
    ];

    function setUp() public {
        // Initialize contract instances from mainnet
        admin = address(this);
        brToken = IERC20(BR_TOKEN);
        votingEscrow = IVotingEscrowCore(VOTING_ESCROW);
        proxyAdmin = ProxyAdmin(PROXY_ADMIN);
        
        // Deploy implementation contract
        implementation = new Airdrop();
        
        // Prepare initialization data
        bytes memory initData = abi.encodeWithSelector(
            Airdrop.initialize.selector,
            ACTIVATION_DELAY,
            VOTING_ESCROW,
            BR_TOKEN,
            admin
        );
        
        // Deploy proxy contract
        TransparentUpgradeableProxy proxy = new TransparentUpgradeableProxy(
            address(implementation),
            PROXY_ADMIN,
            initData
        );
        
        // Cast proxy contract to Airdrop interface
        airdrop = Airdrop(address(proxy));
        
        // Deal BR tokens to this contract for testing
        deal(BR_TOKEN, address(this), 1000000 * 10**18);
        
        // Transfer tokens to airdrop contract for distribution
        brToken.transfer(address(airdrop), 1000000 * 10**18);
        
        // Submit merkle root for the first epoch
        airdrop.submitMerkleRoot(merkleRoot, VALID_DURATION);
        
        // Fast forward time to activation period
        vm.warp(block.timestamp + ACTIVATION_DELAY);
    }

    function testClaimAirdrop() public {
        vm.startPrank(claimant);
        
        // Verify initial claim status
        assertFalse(airdrop.hasUserClaimed(1, claimant), "User should not have claimed");
        
        // Execute claim operation
        airdrop.claim(amount, proof);
        
        // Verify final claim status
        assertTrue(airdrop.hasUserClaimed(1, claimant), "Claim should be successful");
        
        vm.stopPrank();
    }

    function testClaimTwice() public {
        vm.startPrank(claimant);
        
        // First claim should succeed
        airdrop.claim(amount, proof);
        
        // Second claim should fail with duplicate claim error
        vm.expectRevert("USR005");
        airdrop.claim(amount, proof);
        
        vm.stopPrank();
    }

    function testClaimBeforeActivation() public {
        // Reset time to before activation period
        vm.warp(block.timestamp - ACTIVATION_DELAY);
        
        vm.startPrank(claimant);
        // Should fail with not activated error
        vm.expectRevert("USR007");
        airdrop.claim(amount, proof);
        vm.stopPrank();
    }

    function testClaimAfterExpiration() public {
        // Fast forward time to after expiration
        vm.warp(block.timestamp + VALID_DURATION + 1);
        
        vm.startPrank(claimant);
        // Should fail with expired error
        vm.expectRevert("USR008");
        airdrop.claim(amount, proof);
        vm.stopPrank();
    } 
}