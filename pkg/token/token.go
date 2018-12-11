package token

import (
	"errors"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	// ErrMissingHeader means the `Authorization` header was empty. 请求头为空则报错
	ErrMissingHeader = errors.New("The length of the `Authorization` header is zero.")
)

// Context is the context of the JSON web token. 包含 jwt 签名内容的字段
type Context struct {
	ID       uint64
	Username string
	RoleId   int64
}

// secretFunc validates the secret format. 检测 secret 的格式
func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		// Make sure the `alg` is what we except. 确保是我们需要的 `alg` 加密方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(secret), nil
	}
}

// Parse validates the token with the specified secret, Parse使用指定的密钥验证令牌
// and returns the context if the token was valid. 并在令牌有效时返回上下文
func Parse(tokenString string, secret string) (*Context, error) {
	ctx := &Context{}

	// Parse the token. 传入 token 和 secret ，secretFunc 匹配加密算法
	// jwt.Parse 用于解析和验证 token 的合法性
	token, err := jwt.Parse(tokenString, secretFunc(secret))

	// Parse error.
	if err != nil {
		return ctx, err

		// Read the token if it's valid. 如果 token 合法，则进行解析
	} else if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		ctx.ID = uint64(claims["id"].(float64))
		ctx.Username = claims["username"].(string)
		ctx.RoleId = int64(claims["roleid"].(float64))
		return ctx, nil

		// Other errors.
	} else {
		return ctx, err
	}
}

// ParseRequest gets the token from the header and    ParseRequest从头部获取令牌
// pass it to the Parse function to parses the token. 并将其传递给Parse函数以解析令牌
func ParseRequest(c *gin.Context) (*Context, error) {
	header := c.Request.Header.Get("Authorization")

	// Load the jwt secret from config 从config加载jwt secret
	secret := viper.GetString("jwt_secret")

	if len(header) == 0 {
		return &Context{}, ErrMissingHeader
	}

	var t string
	// Parse the header to get the token part. 解析 header 以获取令牌部分
	fmt.Sscanf(header, "Bearer %s", &t)
	return Parse(t, secret)
}

// Sign signs the context with the specified secret.  用指定的 secret 签署上下文
func Sign(ctx *gin.Context, c Context, secret string) (tokenString string, err error) {
	// Load the jwt secret from the Gin config if the secret isn't specified.
	// 如果未指定 secret，则从 Gin 配置加载 jwt secret
	if secret == "" {
		secret = viper.GetString("jwt_secret")
	}
	// The token content.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       c.ID,
		"username": c.Username,
		"roleid":   c.RoleId,
		"nbf":      time.Now().Unix(), // JWT Token 生效时间
		"iat":      time.Now().Unix(), // JWT Token 签发时间
	})
	// Sign the token with the specified secret. 用指定的 secret 签署 token
	tokenString, err = token.SignedString([]byte(secret))

	return
}
