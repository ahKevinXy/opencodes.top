package jwt

import (
	"errors"
	"net/http"
	config2 "opencodes/pkg/config"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/dgrijalva/jwt-go"
	goJwt "github.com/dgrijalva/jwt-go"
)

type Jwt struct {
	SigningKey []byte
}

//载负内容
type CustomClaims struct {
	Id int `json:"id"`
	goJwt.StandardClaims
}

// 错误内容
var (
	TokenExpired     error = errors.New("Token is expired")
	TokenNotValidYet error = errors.New("Token not active yet")
	TokenMalformed   error = errors.New("That's not even a token")
	TokenInvalid     error = errors.New("Couldn't handle this token:")
	//SignKey          string = config2.JwtConfig.Secret
)

//新建一个jwt 实例

func NewJwt() *Jwt {
	return &Jwt{
		[]byte(GetSignKey()),
	}
}

func GetSignKey() string {
	return config2.JwtConfig.Secret
}

//创建 token
func CreateToken(id int) (string, error) {
	expireSeconds := config2.JwtConfig.Timeout
	mySignKey := []byte(GetSignKey())
	expireAt := time.Now().Add(time.Second * time.Duration(expireSeconds)).Unix()
	claims := CustomClaims{
		id,
		goJwt.StandardClaims{
			Audience:  "",
			ExpiresAt: expireAt,
			Id:        "1",
			IssuedAt:  0,
			Issuer:    "",
			NotBefore: 0,
			Subject:   "",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(mySignKey)
	if err != nil {
		return "", err
	}
	return tokenStr, err
}

// 获取过期时间
func GetExpireTime() int64 {
	return 10000
}

//解析Token
func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(GetSignKey()), nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

func Refresh() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 500,
				"msg":  "请求未携带token，无权限访问",
			})
			c.Abort()
			return
		}
		j := NewJwt()

		data, err := j.RefreshToken(token)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 500,
				"msg":  err,
			})
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "OK",
			"data": data,
		})
		c.Abort()
		return
	}
}

// 更新token
func (j *Jwt) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return CreateToken(claims.Id)
	}
	return "", TokenInvalid
}

func GetUid(c *gin.Context) int {
	//获取header
	token := c.Request.Header.Get("token")
	if token == "" {
		return 0
	}
	claims, err := ParseToken(token)
	if err != nil {
		return 0
	}
	return claims.Id
}
