package main

import (
	"fmt"
	"log"
	"os"

	"github.com/golden-expiriensu/sel2sig/input"
	"github.com/golden-expiriensu/sel2sig/search"
)

func main() {
	in, err := input.FromOSArgs()
	if err != nil {
		log.Fatalf("failed to parse args: %v", err)
	}

	result, err := search.SearchDirectory(in.Directory(), in.Selector())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n%s\n", result)
	if in.Operation() == input.OperationSearch {
		os.Exit(0)
	}

	unpacked, err := result.Unpack(in.SelectorWithArgs())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("args: %v\n", unpacked)
}
