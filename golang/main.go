// Copyright (c) 2023, Circle Technologies, LLC. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package main is the entrypoint for the sample code.
package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
)

// Paste your entity public key here.
var publicKeyString = "PASTE_YOUR_PUBLIC_KEY_HERE"

// Store the generated Entity Secret in a secure location rather than directly embedding it within the code.
// Paste your hex encoded Entity Secret here. The length of the hex string should be 64.
var hexEncodedEntitySecret = "PASTE_YOUR_HEX_ENCODED_ENTITY_SECRET_KEY_HERE"

// ParseRsaPublicKeyFromPem parse rsa public key from pem.
func ParseRsaPublicKeyFromPem(pubPEM []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(pubPEM)
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	switch pub := pub.(type) {
	case *rsa.PublicKey:
		return pub, nil
	default:
	}
	return nil, errors.New("key type is not rsa")
}

// EncryptOAEP rsa encrypt oaep.
func EncryptOAEP(pubKey *rsa.PublicKey, message []byte) (ciphertext []byte, err error) {
	random := rand.Reader
	ciphertext, err = rsa.EncryptOAEP(sha256.New(), random, pubKey, message, nil)
	if err != nil {
		return nil, err
	}
	return
}

// GetEntropyCSPRNG creates a random entropy of 32 byte length.
// This function generates a 32 byte random Entity Secret.
// The generation of Entity Secret only need to be executed once unless you need to rotate Entity Secret.
// Store the generated Entity Secret in a secure location rather than directly embedding it within the code.
func GetEntropyCSPRNG() []byte {
	mainBuff := make([]byte, 32)
	_, err := io.ReadFull(rand.Reader, mainBuff)
	if err != nil {
		panic("reading from crypto/rand failed: " + err.Error())
	}
	return mainBuff
}

func main() {
	entitySecret, err := hex.DecodeString(hexEncodedEntitySecret)
	if err != nil {
		panic(err)
	}
	if len(entitySecret) != 32 {
		panic("invalid entity secret")
	}
	pubKey, err := ParseRsaPublicKeyFromPem([]byte(publicKeyString))
	if err != nil {
		panic(err)
	}
	cipher, err := EncryptOAEP(pubKey, entitySecret)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hex encoded entity secret: %x\n", entitySecret)
	fmt.Printf("Entity secret ciphertext: %s\n", base64.StdEncoding.EncodeToString(cipher))
}
