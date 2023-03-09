package util

import (
	"autotec/pkg/env"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

func Decrypt(cipherText string) (plainText string, err error) {
	// prepare cipher

	keyByte := []byte(env.Encrypt_Key)
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return
	}
	nonceSize := gcm.NonceSize()

	// process ciphertext
	ciphertextByte, _ := base64.StdEncoding.DecodeString(cipherText)
	nonce, ciphertextByteClean := ciphertextByte[:nonceSize], ciphertextByte[nonceSize:]
	plaintextByte, err := gcm.Open(
		nil,
		nonce,
		ciphertextByteClean,
		nil)
	if err != nil {
		return
	}
	plainText = string(plaintextByte)
	//
	return
}

func Encrypt(plaintext string) (chipertext string, err error) {
	block, err := aes.NewCipher([]byte(env.Encrypt_Key))
	if err != nil {
		return
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return
	}
	ciphertextByte := gcm.Seal(
		nonce,
		nonce,
		[]byte(plaintext),
		nil)
	chipertext = base64.StdEncoding.EncodeToString(ciphertextByte)
	return
}
