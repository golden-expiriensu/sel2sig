package input

import "github.com/ethereum/go-ethereum/common/hexutil"

func decodeHexString(input string) ([]byte, error) {
	if !has0xPrefix(input) {
		input = "0x" + input
	}
	return hexutil.Decode(input)
}

func has0xPrefix(input string) bool {
	return len(input) >= 2 && input[0] == '0' && (input[1] == 'x' || input[1] == 'X')
}
