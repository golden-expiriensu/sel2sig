package input

import "os"

func FromOSArgs() (Input, error) {
	directory := os.Args[1]

	calldata, err := extractCalldata(os.Args[2])
	if err != nil {
		return Input{}, err
	}

	return Input{directory, calldata}, nil
}
