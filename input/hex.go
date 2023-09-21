package input

import "encoding/hex"

func decodeHexString(input string) ([]byte, error) {
	if has0xPrefix(input) {
		input = input[2:]
	}
	return hex.DecodeString(input)
}

func has0xPrefix(input string) bool {
	return len(input) >= 2 && input[0] == '0' && (input[1] == 'x' || input[1] == 'X')
}
