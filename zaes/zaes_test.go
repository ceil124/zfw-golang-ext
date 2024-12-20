package zaes

import (
	"encoding/base64"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestEncrypt(t *testing.T) {
	a := assert.New(t)

	// 原文
	plaintext := []byte("[{\"username\":\"test1\",\"email\":\"test1@test.com\",\"phone\":\"18511111111\",\"password\":\"test\"}]")

	// AES加密
	ciphertext, err := Encrypt(DefaultAesKey, plaintext)
	if err != nil {
		log.Fatalf("加密失败: %v", err)
	}
	// 将AES加密后的密文编码为Base64字符串
	ciphertextBase64 := base64.StdEncoding.EncodeToString(ciphertext)
	fmt.Printf("加密后的数据: %s\n", ciphertextBase64)

	// 将Base64字符串解码获取原数据密文
	ciphertext2, _ := base64.StdEncoding.DecodeString(ciphertextBase64)
	// 将数据密文使用AES解密，获取原文
	plaintext2, err := Decrypt(DefaultAesKey, ciphertext2)
	if err != nil {
		log.Fatalf("解密失败: %v", err)
	}
	fmt.Printf("解密后的数据: %s\n", plaintext2)

	a.Equal(plaintext, plaintext2)
}
