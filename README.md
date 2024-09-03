# Plasma-DA
PlasmaDA, a configurable system that configure with a protocol ID and an adapter endpoint to fetch and store data from third-party data availability solution.

## Installation
```bash
make install
```

## Features
- [x] Configurable of plasma-da able to set from environment variable or flags.
- [x] Support Celestia data availability solution.
- [x] Support FileStore to store the data.
- [x] Support IPFS to store the data.
- [x] Support ArWeave to store the data.

## Plasma-DA able to run with the following data availability solution.

### Run the Plasma-DA use Celestia as data availability solution. 
```shell 
plasma-da start --da=celestia --da-id=0x000c  --celestia.auth_token=<api_token> --celestia.namespace=<namespace> --celestia.rpc=<rpc_url> 
--evm-rpc-url=<evm_rpc_url> --chain-id=<chain_id> --key-file=<key_file> --passphrase=<passphrase> --plasma-hub-addr=<plasma-hub-address>
```

### Run the Plasma-DA use FileStore as data availability solution. 
```shell
plasma-da start --da=filestore --filestore.path=<path to store data>
```

### Run the Plasma-DA use IPFS as data availability solution. 
```shell
plasma-da start --da=ipfs --da-id=0x000e --ipfs.url=<ipfs node url>
--evm-rpc-url=<evm_rpc_url> --chain-id=<chain_id> --key-file=<key_file> --passphrase=<passphrase> --plasma-hub-addr=<plasma-hub-address>
```

### Run the Plasma-DA use ArWeave as data availability solution. 
```shell
 plasma-da start --da=ar --da-id=0x000d --ar.client_url=http://localhost:8080 --ar.wallet_path=<path to wallet json file>
 --evm-rpc-url=<evm_rpc_url> --chain-id=<chain_id> --key-file=<key_file> --passphrase=<passphrase> --plasma-hub-addr=<plasma-hub-address>
```
