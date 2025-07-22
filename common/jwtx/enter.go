package jwtx

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var (
	AUTH_ENUMS_ATOKEN = "atoken"
	AUTH_ENUMS_RTOKEN = "rtoken"
)

type JWT struct {
	Secret string
}

func NewJWT(secret string) JWT {
	return JWT{
		Secret: secret,
	}
}

type TokenData struct {
	Userid string `json:"user_id"`
	Class  string `json:"class"`
}

func (j JWT) GenToken(userid string, exp time.Duration, class string) (string, error) {
	claims := jwt.MapClaims{
		"userid": userid,
		"class":  class,
		"exp":    time.Now().Add(exp).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(j.Secret))
	return tokenString, err
}

func (j JWT) GenAtoken(userid string, exp time.Duration) (string, error) {
	return j.GenToken(userid, exp, AUTH_ENUMS_ATOKEN)
}

func (j JWT) GenRtoken(userid string, exp time.Duration) (string, error) {
	return j.GenToken(userid, exp, AUTH_ENUMS_RTOKEN)
}

func (j JWT) IdentifyToken(tokenString string) (TokenData, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("不支持的签名方法: %v", token.Header["alg"])
		}
		return []byte(j.Secret), nil
	})
	if err != nil {
		return TokenData{}, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		// 验证token是否过期
		if time.Now().Unix() > int64(claims["exp"].(float64)) {
			return TokenData{}, fmt.Errorf("token已过期")
		}
	} else {
		// 解析失败
		return TokenData{}, fmt.Errorf("无效的token")
	}
	// 解析token成功
	return TokenData{
		Userid: claims["userid"].(string),
		Class:  claims["class"].(string),
	}, nil
}
