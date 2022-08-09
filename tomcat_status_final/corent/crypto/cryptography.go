// cryptography
package crypto

import (
	"bufio"
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
)

func Encrypt(secretKey string, plaintext []byte) (string, error) {
	block, err := des.NewTripleDESCipher([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("Create Cipher Block with secret key:%s", err)
	}
	ciphertext := []byte("abcdef1234567890")
	iv := ciphertext[:des.BlockSize] // const BlockSize = 8
	mode := cipher.NewCBCEncrypter(block, iv)
	origData := PKCS5Padding(plaintext, block.BlockSize())
	encrypted := make([]byte, len(origData))
	mode.CryptBlocks(encrypted, origData)
	encoded := base64.StdEncoding.EncodeToString(encrypted)
	return encoded, nil
}

func Decrypt(secretKey string, encrypted []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher([]byte(secretKey))
	if err != nil {
		return nil, fmt.Errorf("Create Cipher Block with secret key:%s", err)
	}
	ciphertext := []byte("abcdef1234567890")
	iv := ciphertext[:des.BlockSize] // const BlockSize = 8
	blockMode := cipher.NewCBCDecrypter(block, iv)
	decoded, err := base64.StdEncoding.DecodeString(string(encrypted))
	if err != nil {
		return nil, fmt.Errorf("Decode Encrypted text:%s", err)
	}
	decrypted := make([]byte, len(decoded))
	blockMode.CryptBlocks(decrypted, decoded)
	return PKCS5UnPadding(decrypted), nil
}

func GetHash(input []byte) (string, error) {
	hasher := sha1.New()
	br := bufio.NewReader(bytes.NewBuffer(input))
	_, err := io.Copy(hasher, br)
	if err != nil {
		return "", fmt.Errorf("Copy input to Hasher:%s", err)
	}
	return fmt.Sprintf("%x", hasher.Sum(nil)), nil
}

func PKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}
