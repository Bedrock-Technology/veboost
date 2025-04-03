// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

import "../../src/interface/IVotingEscrowIncreasing.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC721/ERC721.sol";

contract MockVotingEscrow is IVotingEscrowCore, ERC721 {
    address public token;
    uint256 private _nextTokenId;

    constructor(address _token) ERC721("VotingEscrow", "veNFT") {
        token = _token;
        _nextTokenId = 1;
    }

    function createLockFor(uint256 _value, address _to) external returns (uint256) {
        // Simulate token transfer from msg.sender to this contract
        require(IERC20(token).transferFrom(msg.sender, address(this), _value), "Transfer failed");
        
        // Mint NFT to _to address
        uint256 tokenId = _nextTokenId++;
        _mint(_to, tokenId);
        
        return tokenId;
    }
}