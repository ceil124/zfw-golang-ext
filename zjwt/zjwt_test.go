package zjwt

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestJwtGenerateAndValidate(t *testing.T) {
	// 引入testify断言库，简化单元测试
	a := assert.New(t)

	token, err := GenerateToken(1)
	a.Nil(err)
	log.Println(token)

	err = ValidateToken(token)
	a.Nil(err)

}
