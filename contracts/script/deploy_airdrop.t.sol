// SPDX-License-Identifier: MIT

pragma solidity ^0.8.17;

import {Script, console} from "forge-std/Script.sol";
import {Airdrop} from "../src/airdrop.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {IVotingEscrowCore} from "../src/interface/IVotingEscrowIncreasing.sol";
/*

# prepare .env file
DEPLOYER=<deployer-account-name>
DEPLOYER_ADDRESS=<deployer-address>
MERKLE_ROOT=<merkle-root-hex>
VOTING_ESCROW_ADDRESS=<voting-escrow-contract-address>
BR_TOKEN_ADDRESS=<br-token-contract-address>
EVM_RPC=<evm-rpc>
ETHERSCAN_API_KEY=<etherscan-api-key>
ETHERSCAN_API_URL=<etherscan-api-url>

# source .env
# verify source code
forge script -vvvv \
    --account $DEPLOYER \
    --sender $DEPLOYER_ADDRESS \
    -f $EVM_RPC \
    --broadcast \
    --verify \
    --verifier custom \
    --verifier-api-key $ETHERSCAN_API_KEY \
    --verifier-url $ETHERSCAN_API_URL \
    script/deploy_airdrop.t.sol:DeployAirdrop

# verify source code using flatted code
forge script -vvvv \
    --account $DEPLOYER \
    --sender $DEPLOYER_ADDRESS \
    -f $EVM_RPC \
    --broadcast \
    script/deploy_airdrop.t.sol:DeployAirdrop
*/

contract DeployAirdrop is Script {
    function run() external {
        address deployer = vm.envAddress("DEPLOYER_ADDRESS");
        

        bytes32 merkleRoot = vm.envBytes32("MERKLE_ROOT");
        address votingEscrow = vm.envAddress("VOTING_ESCROW_ADDRESS");
        address brToken = vm.envAddress("BR_TOKEN_ADDRESS");

        vm.startBroadcast(deployer);

        console.log("[Signer] deployer:", deployer);
        console.log("[Param] merkleRoot:", uint256(merkleRoot));
        console.log("[Contract] votingEscrow:", votingEscrow);
        console.log("[Contract] brToken:", brToken);

        Airdrop airdrop = new Airdrop(
            merkleRoot,
            votingEscrow,
            brToken
        );
        console.log("[Contract] Airdrop deployed at:", address(airdrop));

        vm.stopBroadcast();
    }
}
