package input

import (
	"errors"
	"os"
	"regexp"
)

var (
	regexpErrEthersV5 = regexp.MustCompile(`"code":-?[0-9]+,"data":"(0x[0-9a-fA-F]+)"`)
)

var ErrUnknownLogsStructure = errors.New("neither hex data nor known error logs")

func extractCalldata(input string) ([]byte, error) {
	if res, err := decodeHexString(input); err == nil {
		return res, nil
	}
	return extractCalldataFromFile(input)
}

func extractCalldataFromFile(filepath string) ([]byte, error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	if matches := regexpErrEthersV5.FindSubmatch(content); matches != nil {
		return decodeHexString(string(matches[1]))
	}
	return nil, ErrUnknownLogsStructure
}
