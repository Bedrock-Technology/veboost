// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

import {Test} from "forge-std/Test.sol";
import {Airdrop} from "../src/airdrop.sol";
import "./mocks/MockVotingEscrow.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/utils/cryptography/MerkleProof.sol";

contract MockToken is ERC20 {
    constructor() ERC20("Mock Token", "MTK") {
        _mint(msg.sender, 1000000 * 10**18);
    }
}

contract MerkleTreeTest is Test {
    Airdrop public airdrop;
    MockToken public brToken;
    MockVotingEscrow public votingEscrow;
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
        // Deploy contracts
        brToken = new MockToken();
        votingEscrow = new MockVotingEscrow(address(brToken));
        airdrop = new Airdrop(merkleRoot, address(votingEscrow),
            address(brToken));
        brToken.transfer(address(airdrop), 1000000 * 10**18);
    }

    function test_ClaimTokens() public {
        vm.startPrank(claimant);
        airdrop.claimTokens(amount, proof);
        assertTrue(airdrop.claimed(claimant), "Claim should be successful");
        vm.stopPrank();
    }
}