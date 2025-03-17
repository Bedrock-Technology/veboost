// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

import {Test, console} from "forge-std/Test.sol";
import {Airdrop} from "../src/airdrop.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "./mocks/MockVotingEscrow.sol";
import "@openzeppelin/contracts/utils/cryptography/MerkleProof.sol";

contract MockToken is ERC20 {
    constructor() ERC20("Mock Token", "MTK") {
        _mint(msg.sender, 1000000 * 10**18);
    }
}

contract AirdropTest is Test {
    Airdrop public airdrop;
    MockToken public brToken;
    MockVotingEscrow public votingEscrow;
    
    bytes32 public merkleRoot;
    address public user1;
    address public user2;
    uint256 public amount = 100 * 10**18;

    function setUp() public {
        // Deploy contracts
        brToken = new MockToken();
        votingEscrow = new MockVotingEscrow(address(brToken));
        
        // Setup test accounts
        user1 = address(0x1);
        user2 = address(0x2);
        
        // Create merkle root using the same method as in the contract
        bytes32 leaf = keccak256(bytes.concat(keccak256(abi.encode(user1, amount))));
        bytes32[] memory leaves = new bytes32[](1);
        leaves[0] = leaf;
        
        merkleRoot = leaf; // For single leaf testing
        console.log("MerkleRoot (hex):", vm.toString(merkleRoot));
        
        // Deploy airdrop contract
        airdrop = new Airdrop(
            merkleRoot,
            address(votingEscrow),
            address(brToken)
        );
        
        // Transfer tokens to airdrop contract
        brToken.transfer(address(airdrop), 1000000 * 10**18);
    }

    function test_InitialState() public view {
        assertEq(airdrop.merkleRoot(), merkleRoot);
        assertEq(address(airdrop.votingEscrow()), address(votingEscrow));
        assertEq(address(airdrop.brToken()), address(brToken));
    }

    function test_ClaimTokens() public {
        // Create valid merkle proof for single leaf
        bytes32[] memory proof = new bytes32[](0); // Empty proof for single leaf
        
        vm.startPrank(user1);
        airdrop.claimTokens(amount, proof);
        
        assertTrue(airdrop.claimed(user1));
        vm.stopPrank();
    }

    function testFail_DoubleClaimTokens() public {
        bytes32[] memory proof = new bytes32[](1);
        proof[0] = keccak256(bytes.concat(abi.encodePacked("test proof")));

        vm.startPrank(user1);
        airdrop.claimTokens(amount, proof);
        airdrop.claimTokens(amount, proof); // Should fail
        vm.stopPrank();
    }

    function testFail_InvalidMerkleProof() public {
        bytes32[] memory proof = new bytes32[](1);
        proof[0] = keccak256(bytes.concat(abi.encodePacked("invalid proof")));

        vm.startPrank(user1);
        airdrop.claimTokens(amount, proof);
        vm.stopPrank();
    }

    function test_WithdrawSuccess() public {
        // Fast forward time beyond claim deadline
        vm.warp(block.timestamp + 31 days);
        
        // Record initial balances
        uint256 initialBalance = brToken.balanceOf(address(airdrop));
        uint256 ownerBalance = brToken.balanceOf(address(this));
        
        // Execute withdrawal
        airdrop.withdraw(address(this));
        
        // Verify balances after withdrawal
        assertEq(brToken.balanceOf(address(airdrop)), 0, "Airdrop contract should have 0 balance");
        assertEq(brToken.balanceOf(address(this)), ownerBalance + initialBalance, "Owner should receive all tokens");
    }

    function testFail_WithdrawBeforeDeadline() public {
        // Try to withdraw before deadline (should fail)
        airdrop.withdraw(address(this));
    }

    function testFail_WithdrawByNonOwner() public {
        // Fast forward time beyond claim deadline
        vm.warp(block.timestamp + 31 days);
        
        // Try to withdraw from non-owner account (should fail)
        vm.prank(user1);
        airdrop.withdraw(user1);
    }

    function test_WithdrawAfterClaims() public {
        // First let user1 claim their tokens
        bytes32[] memory proof = new bytes32[](0);
        vm.prank(user1);
        airdrop.claimTokens(amount, proof);

        // Fast forward time beyond claim deadline
        vm.warp(block.timestamp + 31 days);
        
        // Record remaining balance
        uint256 remainingBalance = brToken.balanceOf(address(airdrop));
        
        // Execute withdrawal
        airdrop.withdraw(address(this));
        
        // Verify final balances
        assertEq(brToken.balanceOf(address(airdrop)), 0, "Airdrop contract should have 0 balance");
        assertEq(brToken.balanceOf(address(this)), remainingBalance, "Owner should receive remaining tokens");
    }
}