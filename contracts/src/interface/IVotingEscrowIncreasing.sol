/// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

interface IVotingEscrowCore 
{
    /// @notice Deposit `_value` tokens for `_to`
    /// @param _value Amount to deposit
    /// @param _to Address to deposit
    /// @return TokenId of created veNFT
    function createLockFor(uint256 _value, address _to) external returns (uint256);

    /// @notice Address of the underying ERC20 token.
    function token() external view returns (address);
}