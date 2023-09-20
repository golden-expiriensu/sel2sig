package search

import "github.com/ethereum/go-ethereum/accounts/abi"

type method abi.Method

func (m method) String() string {
	return abi.Method(m).String()
}

func (m method) Unpack(args []byte) (interface{}, error) {
	inputs := abi.Method(m).Inputs
	return inputs.Unpack(args[4:])
}
