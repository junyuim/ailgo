package core_auth

import (
	"crypto/x509"
	"encoding/base64"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

func CreateToken(privateKey string, tokenClaims jwt.Claims) (string, error) {
	// 生成Token，指定签名算法和claims
	der, err := base64.StdEncoding.DecodeString(privateKey)

	if err != nil {
		fmt.Println("[aix] auth.config.private-key decode error: ", err)
		return "", err
	}

	pkey, err1 := x509.ParsePKCS1PrivateKey(der)

	if err1 != nil {
		fmt.Println("[aix] auth.config.private-key parse error: ", err1)
		return "", err1
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, tokenClaims)

	tokenString, err2 := token.SignedString(pkey)

	if err2 != nil {
		fmt.Println("[aix] auth.config.private-key sign error: ", err2)
		return "", err2
	}

	return tokenString, nil
}

func ParseToken(publicKey string, tokenString string, tokenClaims jwt.Claims) error {
	_, err := jwt.ParseWithClaims(tokenString, tokenClaims, func(t *jwt.Token) (interface{}, error) {
		der, err := base64.StdEncoding.DecodeString(publicKey)

		if err != nil {
			fmt.Println("[aix] auth.config.public-key decode error: ", err)
			return "", err
		}

		pkey, err1 := x509.ParsePKCS1PublicKey(der)

		if err1 != nil {
			fmt.Println("[aix] auth.config.public-key parse error: ", err1)
			return "", err1
		}

		return pkey, nil
	})

	// 没有进行验证，外边自行进行验证是否有效
	return err
}
