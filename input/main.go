package input

import (
	"os"

	"github.com/ethereum/go-ethereum/common"
)

type Input struct {
	directory string
	calldata  []byte
}

const (
	OperationSearch = iota
	OperationDecode
)

func FromOSArgs() (Input, error) {
	directory := os.Args[1]

	calldata, err := common.ParseHexOrString(os.Args[2])
	if err != nil {
		return Input{}, err
	}

	return Input{directory, calldata}, nil
}

func (i Input) Operation() uint {
	if len(i.calldata) > 4 {
		return OperationDecode
	}
	return OperationSearch
}

func (i Input) Directory() string {
	return i.directory
}

func (i Input) Selector() [4]byte {
	var selector [4]byte
	copy(selector[:], i.calldata)
	return selector
}

func (i Input) SelectorWithArgs() []byte {
	return i.calldata
}
