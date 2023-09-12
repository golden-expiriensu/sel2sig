package search

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

type Result interface {
	String() string
	Unpack([]byte) (interface{}, error)
}

var ErrNotFound = errors.New("origin of selector is not found")

func SearchFile(filepath string, selector [4]byte) (Result, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to open artifact file: %w", err)
	}
	defer file.Close()

	var artifact struct {
		ABI abi.ABI
	}
	if err = json.NewDecoder(file).Decode(&artifact); err != nil {
		return nil, fmt.Errorf("can't decode artifact source: %w", err)
	}

	resultMethod, _ := artifact.ABI.MethodById(selector[:])
	if resultMethod != nil {
		return extendedMethod(*resultMethod), nil
	}

	resultError, _ := artifact.ABI.ErrorByID(selector)
	if resultError != nil {
		return resultError, nil
	}

	return nil, ErrNotFound
}

type extendedMethod abi.Method

func (m extendedMethod) String() string {
	return abi.Method(m).String()
}

func (m extendedMethod) Unpack(args []byte) (interface{}, error) {
	inputs := abi.Method(m).Inputs
	return inputs.Unpack(args[4:])
}
