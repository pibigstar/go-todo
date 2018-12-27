package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pibigstar/go-todo/config"
	"github.com/pibigstar/go-todo/utils"
)

var (
	// 加密的key值
	secretKey = config.ServerConfig.SecretKey
	// TokenClaimEXP 有效期
	TokenClaimEXP = "exp"
	// TokenClaimOpenID 将用户OpenID存放到token中
	TokenClaimOpenID = "openID"
)

func TestToken(t *testing.T) {
	claims := make(jwt.MapClaims)
	// 有效期
	claims[TokenClaimEXP] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims[TokenClaimOpenID] = "this is user id"

	token, err := utils.GenJwtToken(claims)
	if err != nil {
		fmt.Printf("generate jwt token failed:%e", err)
	}

	fmt.Println("token:", token)

	isToken := utils.CheckJwtToken(token)
	fmt.Println("isToken:", isToken)

	if uid, err := utils.GetOpenIDFromToken(token); err == nil {
		fmt.Println("用户id：", uid)
	}
}
