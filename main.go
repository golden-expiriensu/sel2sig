package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/golden-expiriensu/sel2sig/args"
	"github.com/golden-expiriensu/sel2sig/search"
)

func main() {
	searchFolder, selector, err := args.GetArgs()
	if err != nil {
		log.Fatalf("failed to parse args: %v", err)
	}

	result, err := search.SearchDirectory(searchFolder, selector)
	if errors.Is(err, search.ErrNotFound) {
		log.Fatal("selector origin is not found")
	}
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
