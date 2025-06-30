# Anamd

Official client for Anam Chain

## Environment Setup

### Prerequisites

### **Node.js**

To install Node.js, we recommend using NVM (Node Version Manager)—a handy tool that makes it easy to install and switch between different Node.js versions from the command line. You can get started by running the following command:

```
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.3/install.sh | bash
```

Example
```
$ nvm use 16
Now using node v16.9.1 (npm v7.21.1)
$ node -v
v16.9.1
$ nvm use 14
Now using node v14.18.0 (npm v6.14.15)
$ node -v
v14.18.0
$ nvm install 12
Now using node v12.22.6 (npm v6.14.5)
$ node -v
v12.22.6
```
### **Go**

Anam Chain requires a Go version later than 1.19. To manage different Go versions on your system effortlessly, we recommend using GVM (Golang Version Manager). You can install it by running the following command:

```
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
```

**Install Go using GVM**

After you've installed GVM, you can easily install go by running this:

```
gvm install go1.19.1 -B
gvm use go1.19.1
```

### **Rust**

To build on the Anam Chain, Rust is required. The easiest way to install it is via Rustup, the official installer and version manager for Rust. For a more comprehensive setup guide, feel free to explore the Rust [Getting Started](https://www.rust-lang.org/learn/get-started) page.

**Install Rust**

```
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
```

**Configuring the PATH environment variable**

All Rust tools are installed to the ~/.cargo/bin directory. There you will find the rustc, cargo and rustup binaries. To add them to your PATH, open in your desired editor and add the directory to it.

```
sudo nano ~/.profile
```

Add to the end of ~/.bashrc the following line:

```
export PATH="$HOME/.cargo/bin:$PATH"
```

Then save it and check the current cargo version and update rustup:

```
rustup default stable
cargo version
# If this is lower than 1.55.0, update
rustup update stable
```

You also need to install cargo-generate and cargo-run-script. You can do so with the following commands:

```
cargo install cargo-generate --features vendored-openssl
cargo install cargo-run-script
```

### **Wasm**

Wasm (wasmd) is the backbone of the CosmWasm platform. It is the implementation of a Cosmos zone with wasm smart contracts enabled.

**Install the wasm32 target**

```
# Check the target toolchains of rustup
rustup target list --installed
# Add the wasm toolchain
rustup target add wasm32-unknown-unknown
```

### **Install Anam CLI**

The anamd binary is the official client for Anam Chain. To install it, you need to clone the anam repo:

```
https://github.com/buzz-space-labs/vietchain-alpha1.git
```

And install it using makefile:

```
cd vietchain-alpha1
make
```

## Anamd Guild Testing

### 1. Install Anamd CLI

### 2. Create a Wallet

```
anamd keys add alice --keyring-backend test
```

Get wallet's address

```
anamd keys show alice -a --keyring-backend test
```

### 3. Faucet

Copy your wallet's address to [this page](https://faucet.anam.foundation/).

### 4. Check your balance

```
anamd query bank balances <your-wallet-address> --node <anam-chain-rpc>
```

[`anam-chain-rpc`](./infra/rpc.md) is the official RPC endpoint for Anam Chain.

### 💸 Transfer Token

```
anamd tx bank send alice <recipient-address> 1000000uanam \
  --from alice \
  --chain-id anam \
  --fees 5000uanam \
  --gas auto \
  --gas-adjustment 1.3 \
  --keyring-backend test \
  --node <anam-chain-rpc>
```