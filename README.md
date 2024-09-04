# Optimism Alt-DA
An Optimism Alt-DA server enables OP Stack chains to use third-party data availability providers with an on-chain translation system that converts Keccak256 commitments into provider-specific CIDs (content identifier hashes).

This Alt-DA server uses Keccak256 commitments instead of generic commitments, providing a security advantage. The challenging contract for Keccak256 commitments is well-implemented, audited, and battle-tested, whereas generic commitments often lack an approved challenging logic. This Keccak256 commitments approach is **approved for joining the Superchain**, as seen in the Redstone chain.

## Installation
```bash
make install
```

## Features
- [x] Configurable of alt-da able to set from environment variable or flags.
- [x] Support Celestia data availability solution.
- [x] Support FileStore to store the data.
- [x] Support IPFS to store the data.
- [x] Support ArWeave to store the data.

## Alt-DA able to run with the following data availability solution.

### Run the Alt-DA use Celestia as data availability solution. 
```shell 
alt-da start --da=celestia --da-id=0x000c  --celestia.auth_token=<api_token> --celestia.namespace=<namespace> --celestia.rpc=<rpc_url> 
--evm-rpc-url=<evm_rpc_url> --chain-id=<chain_id> --key-file=<key_file> --passphrase=<passphrase> --plasma-hub-addr=<plasma-hub-address>
```

### Run the Alt-DA use FileStore as data availability solution. 
```shell
alt-da start --da=filestore --filestore.path=<path to store data>
```

### Run the Alt-DA use IPFS as data availability solution. 
```shell
alt-da start --da=ipfs --da-id=0x000e --ipfs.url=<ipfs node url>
--evm-rpc-url=<evm_rpc_url> --chain-id=<chain_id> --key-file=<key_file> --passphrase=<passphrase> --plasma-hub-addr=<plasma-hub-address>
```

### Run the Alt-DA use ArWeave as data availability solution. 
```shell
alt-da start --da=ar --da-id=0x000d --ar.client_url=http://localhost:8080 --ar.wallet_path=<path to wallet json file>
--evm-rpc-url=<evm_rpc_url> --chain-id=<chain_id> --key-file=<key_file> --passphrase=<passphrase> --plasma-hub-addr=<plasma-hub-address>
```
