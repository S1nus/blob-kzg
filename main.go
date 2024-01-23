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
	blob := getRandBlob()
	cmt, err := kzg.BlobToKZGCommitment(blob)
	if err != nil {
		fmt.Println("Error creating KZG commitment: ", err)
		return
	}
	fmt.Printf("Commitment bytes: %x \n", cmt)
}

func getRandBlob() kzg.Blob {
	var blob kzg.Blob
	for i := 0; i < kzg.BytesPerBlob; i += kzg.BytesPerFieldElement {
		fieldElementBytes := getRandFieldElement()
		copy(blob[i:i+kzg.BytesPerFieldElement], fieldElementBytes[:])
	}
	return blob
}

func getRandFieldElement() kzg.Bytes32 {
	bytes := make([]byte, 31)
	_, err := rand.Read(bytes)
	if err != nil {
		panic("failed to get random field element")
	}

	// This leaves the first byte in fieldElementBytes as
	// zero, which guarantees it's a canonical field element.
	var fieldElementBytes kzg.Bytes32
	copy(fieldElementBytes[1:], bytes)
	return fieldElementBytes
}
