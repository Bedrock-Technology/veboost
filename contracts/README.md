# **VeBoost Airdrop Contract**

A **Merkle Tree-based** token airdrop contract system that automatically locks distributed tokens into a **VotingEscrow** contract.

---

## **🚀 Core Features**

- ✅ **Merkle Tree Verification** – Validates user eligibility using Merkle proofs  
- ✅ **Automatic Token Locking** – Airdropped tokens are automatically locked in the **VotingEscrow** contract  
- ✅ **Duplicate Claim Prevention** – Each address can only claim once  
- ✅ **Event Tracking** – Records all claim events  

---

## **📜 Contract Architecture**

| Contract                         | Description                                 |
|----------------------------------|---------------------------------------------|
| `Airdrop.sol`                    | Main airdrop contract                      |
| `IVotingEscrowIncreasing.sol`    | VotingEscrow interface contract            |

---

## **🛠 Development Tools**

This project uses the **Foundry** framework:

- **Forge** – Ethereum testing framework  
- **Cast** – Command-line tool for interacting with EVM smart contracts  
- **Anvil** – Local Ethereum node  

---

## **🚀 Getting Started**

### **📌 Install Dependencies**
```sh
forge install
```

### **🔧 Compile Contracts**
```sh
forge build
```

### **🧪 Run Tests**
```sh
forge test
```

---

## **📤 Deployment Process**

1. Prepare **Merkle Tree root**  
2. Deploy **VotingEscrow contract**  
3. Deploy **Airdrop contract** with:
   - **Merkle Root**
   - **VotingEscrow contract address**
   - **BR Token address**

```sh
forge script script/Airdrop.s.sol:AirdropScript --rpc-url <your_rpc_url> --private-key <your_private_key>
```

