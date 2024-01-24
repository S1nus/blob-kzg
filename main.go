package main

import (
	"crypto/rand"
	"fmt"

	cblob "github.com/celestiaorg/celestia-app/pkg/blob"
	inc "github.com/celestiaorg/celestia-app/pkg/inclusion"
	namespace "github.com/celestiaorg/celestia-app/pkg/namespace"
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
	namespaceId, err := generateRandomBytes(namespace.NamespaceIDSize)
	if err != nil {
		fmt.Println("Error generating namespace id: ", err)
		return
	}
	namespace := namespace.Namespace{Version: namespace.NamespaceVersionMax, ID: namespaceId}
	b := cblob.New(namespace, blob[:], 0)
	cc, err := inc.CreateCommitment(b)
	if err != nil {
		fmt.Println("Error creating celestia blob commitment: ", err)
		return
	}
	fmt.Println("Blob commitment: ", cc)
	fmt.Println("KZG commitment: ", cmt)
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

func generateRandomBytes(length int) ([]byte, error) {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}
	return randomBytes, nil
}
