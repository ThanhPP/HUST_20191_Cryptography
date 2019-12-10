package main

import (
	"crypto/dsa"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"hash"
	"io"
	"math/big"
	"os"
)

func main() {
	params := new(dsa.Parameters)
	//fmt.Println(params)

	if err := dsa.GenerateParameters(params, rand.Reader, dsa.L1024N160); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//Generates private key and public key pair
	//private key
	privateKey := new(dsa.PrivateKey)
	privateKey.PublicKey.Parameters = *params
	dsa.GenerateKey(privateKey, rand.Reader)
	fmt.Printf("Private key = %x \n", privateKey)

	//public key
	var publicKey dsa.PublicKey
	publicKey = privateKey.PublicKey
	fmt.Printf("\nPublic key = %x \n", publicKey)

	//Sign
	var h hash.Hash
	h = sha256.New()
	r := big.NewInt(0)
	s := big.NewInt(0)

	message := "This is a message need to verify !!!"

	io.WriteString(h, message)
	signHash := h.Sum(nil)

	r, s, err := dsa.Sign(rand.Reader, privateKey, signHash)
	if err != nil {
		fmt.Println(err)
	}

	signature := r.Bytes()
	signature = append(signature, s.Bytes()...)
	fmt.Printf("Signature = %x \n", signature)

	//Verify
	verifyStatus := dsa.Verify(&publicKey, signHash, r, s)
	fmt.Println()
	fmt.Println("Message = ", message)
	fmt.Println("Message Hash = ", signHash)
	fmt.Println("Verify status : ", verifyStatus)

	//Verify wrong message
	message2 := "This is a message need to verify !!! - but modified"
	io.WriteString(h, message2)
	signHash2 := h.Sum(nil)

	verifyStatus2 := dsa.Verify(&publicKey, signHash2, r, s)
	fmt.Println()
	fmt.Println("Message = ", message2)
	fmt.Println("Message Hash = ", signHash2)
	fmt.Println("Verify status : ", verifyStatus2)
}
