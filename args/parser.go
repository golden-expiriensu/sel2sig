package args

import (
	"os"

	"github.com/ethereum/go-ethereum/common"
)

func GetArgs() (searchDirectory string, selector [4]byte, _ error) {
	searchDirectory = os.Args[1]

	selectorBytes, err := common.ParseHexOrString(os.Args[2])
	if err != nil {
		return "", [4]byte{}, err
	}

	copy(selector[:], selectorBytes)
	return searchDirectory, selector, nil
}
