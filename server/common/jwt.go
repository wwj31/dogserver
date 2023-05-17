package common

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"reflect"
	"server/common/log"
)

var secretKey = []byte("1234321")

func JWTSignedToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func JWTParseToken[CLAIM jwt.Claims](token string, claim CLAIM) (CLAIM, error) {
	// 解析并验证 JWT
	getSecretKey := func(token *jwt.Token) (interface{}, error) { return secretKey, nil }
	parsedToken, err := jwt.ParseWithClaims(token, claim, getSecretKey)
	if err != nil {
		return claim, err
	}

	if !parsedToken.Valid {
		return claim, fmt.Errorf("token is invalid ")
	}

	claims, good := parsedToken.Claims.(CLAIM)
	if !good {
		log.Warnw("asset token failed ", "parsedToken", reflect.TypeOf(parsedToken.Claims))
		return claim, fmt.Errorf("jwt parsed token claims assert failed")
	}
	claim = claims
	return claim, nil
}
