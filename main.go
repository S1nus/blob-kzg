package main

import (
	"fmt"

	kzg "github.com/ethereum/c-kzg-4844/bindings/go"
)

func main() {
	fmt.Println("Hello world")
	err := kzg.LoadTrustedSetupFile("trusted_setup.txt")
	defer kzg.FreeTrustedSetup()
	if err != nil {
		fmt.Println(err)
	}
}
