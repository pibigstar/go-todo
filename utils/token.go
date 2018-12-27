package utils

import (
	"errors"
	"time"

	"github.com/pibigstar/go-todo/config"

	"github.com/dgrijalva/jwt-go"
)

//  使用jwt 生成token 与使用

var (
	// 加密的key值
	secretKey = config.ServerConfig.SecretKey
	// TokenClaimEXP 有效期标识
	TokenClaimEXP = "exp"
	// TokenClaimOpenID 用户OpenID标识
	TokenClaimOpenID = "openID"
	// TermOfValidity 有效期7天
	TermOfValidity = 24 * 7
)

// GenOpenIDToken 根据OPenID生成token
func GenOpenIDToken(openID string) (string, error) {
	claims := make(jwt.MapClaims)
	// 有效期
	claims[TokenClaimEXP] = time.Now().Add(time.Hour * time.Duration(TermOfValidity)).Unix()
	claims[TokenClaimOpenID] = openID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secretKey))
}

// GenJwtToken 生成token
func GenJwtToken(claims jwt.MapClaims) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secretKey))
}

// CheckJwtToken 检查token是否有效
func CheckJwtToken(tokenString string) bool {
	if tokenString == "" {
		return false
	}
	if err := CheckJwtTokenExpected(tokenString); err != nil {
		return false
	}
	return true
}

// CheckJwtTokenExpected 检查token是否过期
func CheckJwtTokenExpected(tokenString string) error {
	token, err := ParseJwtToken(tokenString)
	if err != nil {
		return err
	}
	return token.Claims.Valid()
}

// ParseJwtToken 解析token
func ParseJwtToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return nil, errors.New("unexpected token claims")
		}
		return []byte(secretKey), nil
	})

	return token, err
}

// GetOpenIDFromToken 从token中拿到openid
func GetOpenIDFromToken(tokenString string) (string, error) {
	token, err := ParseJwtToken(tokenString)
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if v, ok := claims[TokenClaimOpenID]; ok {
			return v.(string), nil
		}
	}

	return "", errors.New("failed get token")
}
