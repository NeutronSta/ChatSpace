package middleware

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"strings"
	"time"
)

var mySigningKey = []byte("Vermont83079->0")

type MyClaims struct {
	//Username string `json:"username"`
	UserId uint `json:"UserId"`
	jwt.RegisteredClaims
}

func NewToken(id uint) string {
	claims := MyClaims{
		//name,
		id,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "NeutronStar",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		log.Println("can not create Token")
	}
	return ss
}

func ParseToken(ss string) (uint, error) {
	token, e := jwt.ParseWithClaims(ss, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if e != nil {
		return 0, e
	}
	claims := token.Claims.(*MyClaims)
	id := claims.UserId
	return id, e
}

func GetID(Authorization string) (uint, error) {
	if Authorization == "" {

		return 0, errors.New("without Authorization")
	}

	parts := strings.Fields(Authorization)
	if len(parts) != 2 || parts[0] != "Bearer" {
		// 处理 Authorization 头部不符合预期的情况
		return 0, errors.New("unexpect Authorization")
	}

	// 提取 token 值
	Token := parts[1]

	id, err := ParseToken(Token)
	if err != nil {
		return 0, err
	}
	return id, err
}
