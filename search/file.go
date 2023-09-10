package search

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

type Artifact struct {
	ABI abi.ABI
}

var ErrNotFound = errors.New("origin of selector is not found")

func SearchFile(filepath string, selector [4]byte) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", fmt.Errorf("failed to open artifact file: %w", err)
	}
	defer file.Close()

	var artifact Artifact
	if err = json.NewDecoder(file).Decode(&artifact); err != nil {
		return "", fmt.Errorf("can't decode artifact source: %w", err)
	}

	resultMethod, _ := artifact.ABI.MethodById(selector[:])
	if resultMethod != nil {
		return resultMethod.String(), nil
	}

	resultError, _ := artifact.ABI.ErrorByID(selector)
	if resultError != nil {
		return resultError.String(), nil
	}

	return "", ErrNotFound
}
