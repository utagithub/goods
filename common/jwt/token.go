/*
 * @Author: 尤_Ta
 * @Date:  17:07
 * @Last Modified by: 尤_Ta
 * @Last Modified time:  17:07
 */

package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type CustomClaims struct {
	// 可根据需要自行添加字段
	Username             string `json:"username"`
	jwt.RegisteredClaims        // 内嵌标准的声明
}

const TokenExpireDuration = time.Hour * 24

// Secret 用于加盐的字符串
var Secret = []byte("laowang")

// GenToken 生成JWT
func GenToken(username string) interface{} {
	// 创建一个我们自己的声明
	claims := CustomClaims{
		username, // 自定义字段
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
			Issuer:    "goods", // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	newToken, _ := token.SignedString(Secret)
	data := map[string]interface{}{
		"token":  newToken,
		"expire": time.Now().Add(TokenExpireDuration),
	}
	return data

}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*CustomClaims, error) {
	// 解析token
	// 如果是自定义Claim结构体则需要使用 ParseWithClaims 方法
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}
	// 对token对象中的Claim进行类型断言
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
