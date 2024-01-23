package main

import (
	"crypto/rand"
	"fmt"

	kzg "github.com/ethereum/c-kzg-4844/bindings/go"
)

const (
	MaxBytesForBlob = kzg.BytesPerBlob
)

func main() {
	err := kzg.LoadTrustedSetupFile("trusted_setup.txt")
	defer kzg.FreeTrustedSetup()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Max bytes per blob: %d \n", MaxBytesForBlob)
	fmt.Printf("Shares per blob: %d \n", MaxBytesForBlob/512)
	blob := getRandomBlob()
	fmt.Printf("blob: %x\n", blob)
	cmt, err := kzg.BlobToKZGCommitment(blob)
	if err != nil {
		fmt.Println("Error creating KZG commitment: ", err)
		return
	}
	fmt.Printf("Commitment bytes: %x \n", cmt)
}

func getRandomShare() [512]byte {
	randomBytes := make([]byte, 512)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic("couldn't generate shares")
	}
	return [512]byte(randomBytes)

}

func getRandomBlob() [MaxBytesForBlob]byte {
	randomBytes := make([]byte, MaxBytesForBlob)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic("couldn't generate blob")
	}
	return [MaxBytesForBlob]byte(randomBytes)

}
