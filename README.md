# sel2sig
CLI tool finding signature of a selector in artifacts folder

## How to use
```shell
go build -o sel2sig && chmod +x sel2sig
```
example:
```
./sel2sig /home/danil/solidity/projectContainingIERC20/artifacts 0x70a08231
```
output:
```
function balanceOf(address account) view returns(uint256)
```
