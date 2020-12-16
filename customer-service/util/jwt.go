package util

import (
	"github.com/dgrijalva/jwt-go"
	"videoWeb/customer-service/config"
)

// 生成token
func GenerateToken(claims map[string]interface{}) (string, error) {
	at := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		buildClaims(claims),
	)
	token, err := at.SignedString([]byte(getSecret()))
	if err != nil {
		return "", err
	}
	return token, nil
}

// 解密token
func ParseToken(token string, item string) (interface{}, error) {
	claim, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(getSecret()), nil
	})
	if err != nil {
		return nil, err
	}
	return claim.Claims.(jwt.MapClaims)[item], nil
}

// 构建载荷
func buildClaims(claims map[string]interface{}) jwt.MapClaims {
	mapClaims := jwt.MapClaims{}
	for k, v := range claims {
		mapClaims[k] = v
	}
	return mapClaims
}

// 获取加密秘钥
func getSecret() string {
	return config.Conf.SecretCode
}
