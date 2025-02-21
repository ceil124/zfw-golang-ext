package zjwt

import (
	_ "embed"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

//go:embed test_cert/private_key.pem
var privateKeyBytes []byte

//go:embed test_cert/public_key.pem
var publicKeyBytes []byte

func TestJwtGenerateAndValidate(t *testing.T) {
	// 引入testify断言库，简化单元测试
	a := assert.New(t)

	token, err := GenerateToken(privateKeyBytes, 1)
	a.Nil(err)
	log.Println(token)

	err = ValidateToken(publicKeyBytes, token)
	a.Nil(err)
}
