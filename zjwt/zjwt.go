package zjwt

import (
	_ "embed"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

func GenerateToken(prKeyByteArr []byte, hourNum int) (string, error) {
	now := time.Now()

	key, err := jwt.ParseRSAPrivateKeyFromPEM(prKeyByteArr)
	if err != nil {
		return "", errors.Join(errors.New("私钥文件加载失败"), err)
	}

	// 生成claims对象
	claims := jwt.RegisteredClaims{
		IssuedAt:  jwt.NewNumericDate(now),
		ID:        uuid.NewString(),
		ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(hourNum) * time.Hour)),
	}
	// 使用密钥文件签名claims，生成jwt token
	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		return "", errors.Join(errors.New("claims生成失败"), err)
	}

	return token, nil
}

func ValidateToken(pubKeyByteArr []byte, token string) error {
	key, err := jwt.ParseRSAPublicKeyFromPEM(pubKeyByteArr)
	if err != nil {
		return errors.Join(errors.New("公钥文件加载失败"), err)
	}

	// 解析claims文件
	parsed, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method:%s", jwtToken.Header["alg"])
		}
		return key, nil
	})
	if err != nil {
		return errors.Join(errors.New("claims解密失败"), err)
	}

	_, ok := parsed.Claims.(jwt.MapClaims)
	if !ok || !parsed.Valid {
		return errors.New("令牌无效")
	}
	return nil
}
