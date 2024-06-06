# Plasma-DA
PlasmaDA, a configurable system that configure with a protocol ID and an adapter endpoint to fetch and store data from third-party data availability solution.

## Installation
```bash
make install
```

## How to use
### Step 1: Create config file. The config file use tomal format.
```shell
plasma-da init --da <data availability>
```
Example:
```shell
plasma-da init --da celestia
```
If you want to change the config file, you can edit config.toml file in the $HOME/.plasma-da/config/config.toml for default.
The config file will look like this:
```toml
[server]
http_host = "localhost"
http_port = 3128
da = "celestia"

[celestia]
rpc_port = "http://localhost:7980"
auth_token = ""
namespace = ""
max_block_size = 2000
gas_price = 0
eth_fallback_disabled = false

[filestore]
path = ".plasma-da/data/filestore"
```
### Step 2: Start the Plasma-DA
```shell
plasma-da start
```


