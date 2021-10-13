package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

var (
	privateKey *rsa.PrivateKey
	publicKey  rsa.PublicKey
)

// Encrypt information
func encryption(info []byte) []byte {
	var err error

	//-----generating the public/private key pairs-----
	privateKey, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	publicKey = privateKey.PublicKey

	//-----encryption---
	// Optimal Asymmetric Encryption Padding (OAEP)
	encryptedBytes, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		&publicKey,
		info,
		nil)

	if err != nil {
		panic(err)
	}
	return encryptedBytes
}

// Decrypt information
func decryption(encryptedBytes []byte) string {
	//-----decryption-----
	decryptedBytes, err := privateKey.Decrypt(
		nil,
		encryptedBytes,
		&rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		panic(err)
	}
	return string(decryptedBytes)
}

/*
// Generate Digital Signature
func digitalSignature() []byte {
	//=============DIGITAL SIGNATURE================

	//----generating the signature--------
	msg := []byte("Verfiable message.")
	msgHash := sha256.New()
	_, err := msgHash.Write(msg)
	if err != nil {
		panic(err)
	}
	msgHashSum = msgHash.Sum(nil)
	signature, _ = rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, msgHashSum, nil)
	return signature
}
*/

/*
// Verify Digital Signature
func verifySignature() {
	//----verifying the signature----------
	err := rsa.VerifyPSS(&publicKey, crypto.SHA256, msgHashSum, signature, nil)
	if err != nil {
		fmt.Println("Could not verify signature!")
	} else {
		fmt.Println("Signature verified!")
	}
}
*/
