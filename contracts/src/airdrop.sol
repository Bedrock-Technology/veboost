// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

import "@openzeppelin/contracts/utils/cryptography/MerkleProof.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./interface/IVotingEscrowIncreasing.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract Airdrop is Ownable {
    uint256 public claimDeadline = block.timestamp + 30 days;
    bytes32 public immutable merkleRoot;
    mapping(address => bool) public claimed;
    
    IVotingEscrowCore public immutable votingEscrow;
    IERC20 public immutable brToken;

    constructor(
        bytes32 _merkleRoot, 
        address _votingEscrow, 
        address _brToken
    ) {
        merkleRoot = _merkleRoot;
        votingEscrow = IVotingEscrowCore(_votingEscrow);
        brToken = IERC20(_brToken);
    }

    function claimTokens(uint256 amount, bytes32[] calldata merkleProof) external {
        require(!claimed[msg.sender], "Already claimed");
    
        bytes32 leaf = keccak256(bytes.concat(keccak256(abi.encode(msg.sender, amount))));
        require(MerkleProof.verify(merkleProof, merkleRoot, leaf), "Invalid merkle proof");

        claimed[msg.sender] = true;

        require(brToken.approve(address(votingEscrow), amount), "Approve failed");
       
        uint256 veNFTId = votingEscrow.createLockFor(amount, msg.sender);

        emit TokensClaimed(msg.sender, amount, veNFTId);
    }

    function withdraw(address to) external onlyOwner {
        require(block.timestamp > claimDeadline, "Claim period not ended");
        uint256 balance = brToken.balanceOf(address(this));
        require(balance > 0, "No tokens to withdraw");
        require(brToken.transfer(to, balance), "Transfer failed");
        
        emit TokensWithdrawn(to, balance);
    }

    event TokensClaimed(address indexed claimant, uint256 amount, uint256 veNFTId);
    event TokensWithdrawn(address indexed to, uint256 amount);
}