# sel2sig

**sel2sig** simplifies [EVM](https://ethereum.org/en/developers/docs/evm/) calldata/error decoding using project-generated artifacts with contracts' [ABIs](https://docs.soliditylang.org/en/latest/abi-spec.html).

## Features

- Find the signature of a [selector](https://docs.soliditylang.org/en/latest/abi-spec.html#function-selector).
- Decode ABI-encoded [calldata](https://docs.soliditylang.org/en/latest/abi-spec.html#argument-encoding).
- Decode ABI-encoded [custom error](https://docs.soliditylang.org/en/latest/abi-spec.html#errors).
- Decode error logs from supported libraries.

## Supported Libraries

- [ethers.js](https://docs.ethers.org/v5/)

## Getting Started

To utilize the decoding features of **sel2sig**, you'll need to provide the following:

1. Path to the artifacts folder within your project.
2. The calldata you want to decode.

You can provide calldata in one of the following ways:

- 4-byte selector
- ABI-encoded calldata/error
- Path to logs containing an error

## Building the CLI Tool

To build **sel2sig**, follow these steps:

1. Clone this repository to [GOPATH](https://pkg.go.dev/cmd/go#hdr-GOPATH_environment_variable)/src/:

   ```
   cd "$(go env GOPATH)/src" && git clone https://github.com/golden-expiriensu/sel2sig
   ```

2. Execute install script:

   ```
   ./install.sh
   ```

## Example Usage

Below are examples of how to use the tool. Please note that all these commands were executed within the folder containing 'artifacts' folder with compilation results of [OpenZeppelin](https://github.com/OpenZeppelin/openzeppelin-contracts/tree/master)'s smart contracts.

1. Find origin of a selector:

   ```
   $ sel2sig ./artifacts 0x70a08231
   
    function balanceOf(address account) view returns(uint256)
   ```

2. Decode calldata:

   ```
   $ sel2sig ./artifacts 0xa9059cbb000000000000000000000000e058fBFD6e991320C1Eb1B17808F742DE92410bC00000000000000000000000000000000000000000000000d8d726b7177a80000

    function transfer(address to, uint256 amount) returns(bool)

    decoded args: [
      0xe058fBFD6e991320C1Eb1B17808F742DE92410bC,
      250000000000000000000
    ]
   ```
3. Decode received custom error:

   ```
   $ sel2sig ./artifacts 0xe450d38c000000000000000000000000e058fBFD6e991320C1Eb1B17808F742DE92410bC00000000000000000000000000000000000000000000000d8d726b7177a8000000000000000000000000000000000000000000000000001043561a8829300000

    error ERC20InsufficientBalance(address sender, uint256 balance, uint256 needed)

    decoded args: [
      0xe058fBFD6e991320C1Eb1B17808F742DE92410bC,
      250000000000000000000,
      300000000000000000000
    ]
   ```
4. Decode error logs:

   ```
   $ sel2sig ./artifacts ./file_with_error_logs

    error ERC20InsufficientBalance(address sender, uint256 balance, uint256 needed)

    decoded args: [
      0xe058fBFD6e991320C1Eb1B17808F742DE92410bC,
      250000000000000000000,
      300000000000000000000
    ]
   ```
   Where file_with_error_logs may include stack traces, error messages, transaction details etc

## License

**sel2sig** is open-source software released under the [MIT License](LICENSE). Feel free to use and modify it according to your needs.
