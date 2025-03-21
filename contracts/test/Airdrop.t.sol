// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

import {Test, console} from "forge-std/Test.sol";
import {Airdrop} from "../src/airdrop.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "./mocks/MockVotingEscrow.sol";
import "@openzeppelin/contracts/utils/cryptography/MerkleProof.sol";

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

contract MockToken is ERC20 {
    constructor() ERC20("Mock Token", "MTK") {
        _mint(msg.sender, 1000000 * 10 ** 18);
    }
}

contract AirdropTest is Test {
    Airdrop private implementation;
    Airdrop private airdrop;
    MockToken private brToken;
    MockVotingEscrow private votingEscrow;
    ProxyAdmin private proxyAdmin;
    address private admin;
    bytes32 private merkleRoot;
    uint32 private activationDelay = 1 days;
    uint32 private validDuration = 30 days;

    function setUp() public {
        admin = address(this);
        brToken = new MockToken();
        votingEscrow = new MockVotingEscrow(address(brToken));
        merkleRoot = keccak256(bytes.concat(keccak256(abi.encode(msg.sender, 1000))));
        
        // Deploy proxy admin contract
        proxyAdmin = new ProxyAdmin();
        
        // Deploy implementation contract
        implementation = new Airdrop();
        
        // Prepare initialization data
        bytes memory initData = abi.encodeWithSelector(
            Airdrop.initialize.selector,
            activationDelay,
            address(votingEscrow),
            address(brToken),
            admin
        );
        
        // Deploy proxy contract
        TransparentUpgradeableProxy proxy = new TransparentUpgradeableProxy(
            address(implementation),
            address(proxyAdmin),
            initData
        );
        
        // Cast proxy contract to Airdrop interface
        airdrop = Airdrop(address(proxy));
        
        // Transfer tokens for testing
        brToken.transfer(address(airdrop), 1000000 * 10**18);
    }

    function testInitialize() public view {
        assertEq(airdrop.activationDelay(), activationDelay);
        assertEq(airdrop.currentEpoch(), 0);
        assertEq(address(airdrop.votingEscrow()), address(votingEscrow));
        assertEq(address(airdrop.brToken()), address(brToken));
    }

    function testSubmitMerkleRoot() public {
        airdrop.submitMerkleRoot(merkleRoot, validDuration);
        Airdrop.DistributionRoot memory distribution = airdrop.getDistributionRoot(1);
        assertEq(distribution.root, merkleRoot);
        assertEq(distribution.duration, validDuration);
        assertEq(distribution.disabled, false);
    }

    function testClaim() public {
        // Submit merkle root and wait for activation
        airdrop.submitMerkleRoot(merkleRoot, validDuration);
        vm.warp(block.timestamp + activationDelay);

        // Calculate leaf using the same method as in contract
        bytes32 leaf = keccak256(bytes.concat(keccak256(abi.encode(address(this), 1000))));
        
        // Use this leaf as merkleRoot (simplified Merkle tree)
        merkleRoot = leaf;
        airdrop.updateMerkleRoot(merkleRoot);
        
        // Create empty proof (since we use leaf as root directly)
        bytes32[] memory proof = new bytes32[](0);

        // Execute claim
        airdrop.claim(1000, proof);
        
        // Verify claim success
        assertTrue(airdrop.hasUserClaimed(1, address(this)));
    }

    function testPauseAndUnpause() public {
        // Submit merkle root and wait for activation
        bytes32 leaf = keccak256(bytes.concat(keccak256(abi.encode(address(this), 1000))));
        merkleRoot = leaf;
        
        airdrop.submitMerkleRoot(merkleRoot, validDuration);
        bytes32[] memory proof = new bytes32[](0);

        // Test pause functionality
        airdrop.pause();
        vm.expectRevert("Pausable: paused");
        airdrop.claim(1000, proof);

        // Test unpause and claim
        airdrop.unpause();
        vm.warp(block.timestamp + activationDelay);
        airdrop.claim(1000, proof);
        assertTrue(airdrop.hasUserClaimed(1, address(this)));
    }

    function testUpdateMerkleRoot() public {
        airdrop.submitMerkleRoot(merkleRoot, validDuration);
        bytes32 newRoot = keccak256(bytes.concat(keccak256(abi.encode(msg.sender, 2000))));
        airdrop.updateMerkleRoot(newRoot);
        Airdrop.DistributionRoot memory distribution = airdrop.getDistributionRoot(1);
        assertEq(distribution.root, newRoot);
    }

    function testUpdateDuration() public {
        airdrop.submitMerkleRoot(merkleRoot, validDuration);
        uint32 newDuration = 60 days;
        airdrop.updateDuration(newDuration);
        Airdrop.DistributionRoot memory distribution = airdrop.getDistributionRoot(1);
        assertEq(distribution.duration, newDuration);
    }

    function testSetAirdropDisabled() public {
        airdrop.submitMerkleRoot(merkleRoot, validDuration);
        airdrop.setAirdrop(true);
        Airdrop.DistributionRoot memory distribution = airdrop.getDistributionRoot(1);
        assertEq(distribution.disabled, true);
    }
}
