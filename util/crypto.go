// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

// MD5 return the MD5 checksum of the data
func MD5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

// SHA256 return the SHA256 checksum of the data
func SHA256(str string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(str)))
}

func EncodeAESWithKey(key, str string) string {
	hash := md5.New()
	hash.Write([]byte(key))
	keyData := hash.Sum(nil)
	block, err := aes.NewCipher(keyData)
	if err != nil {
		panic(err)
	}
	iv := []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}
	enc := cipher.NewCBCEncrypter(block, iv)
	content := _PKCS5Padding([]byte(str), block.BlockSize())
	crypted := make([]byte, len(content))
	enc.CryptBlocks(crypted, content)
	return base64.StdEncoding.EncodeToString(crypted)
}

func DecodeAESWithKey(key, str string) string {
	hash := md5.New()
	hash.Write([]byte(key))
	keyData := hash.Sum(nil)
	block, err := aes.NewCipher([]byte(keyData))
	if err != nil {
		panic(err)
	}
	iv := []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}
	messageData, _ := base64.StdEncoding.DecodeString(str)
	dec := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(messageData))
	dec.CryptBlocks(decrypted, messageData)
	return string(_PKCS5Unpadding(decrypted))
}

func EqualAES(key, raw, hash string) bool {
	return EncodeAESWithKey(key, raw) == hash
}

func _PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func _PKCS5Unpadding(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}
