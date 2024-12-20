package zaes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

var DefaultAesKey = []byte("OjoYRaKF1loHPNcdHGRQqp7uy0X9SH8F")

func Encrypt(key []byte, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)

	if err != nil {
		return nil, err
	}

	// PKCS7填充
	padding := block.BlockSize() - len(text)%block.BlockSize()
	padtext := append(text, bytes.Repeat([]byte{byte(padding)}, padding)...)

	// 初始化向量
	ciphertext := make([]byte, block.BlockSize()+len(padtext))
	iv := ciphertext[:block.BlockSize()]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[block.BlockSize():], padtext)

	return ciphertext, nil
}

func Decrypt(key []byte, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 初始化向量
	if len(ciphertext) < block.BlockSize() {
		return nil, errors.New("ciphertext too short")
	}
	iv := ciphertext[:block.BlockSize()]
	ciphertext = ciphertext[block.BlockSize():]

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	// PKCS7去填充
	padding := int(ciphertext[len(ciphertext)-1])
	return ciphertext[:len(ciphertext)-padding], nil
}
