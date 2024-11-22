package utils

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var (
	Expire  uint64
	Refresh uint64
	Secret  *string
)

func init() {
	Expire = 86400
	Refresh = 43200
}

// 生成token
func GenerateToken(name string) (string, error) {

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Second * time.Duration(Expire)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(*Secret))
}

// 提取token
func GetToken(c *gin.Context) string {
	bearerToken := c.GetHeader("Authorization")

	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

// 验证token
func VerifyToken(c *gin.Context) error {

	tokenString := GetToken(c)
	_, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(*Secret), nil
	})
	if err != nil {
		return err
	}
	return nil
}

func GetTokenName(c *gin.Context) (string, error) {

	tokenString := GetToken(c)
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(*Secret), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		name, ok := claims["name"].(string) // 进行类型断言
		if ok {
			return name, nil
		}
		return "", errors.New("name claim is not a string")
	}
	return "", nil
}

// 提取token过期时间
func GetTokenExpire(c *gin.Context) (uint64, error) {

	tokenString := GetToken(c)
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(*Secret), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		exp, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["exp"]), 10, 64)
		if err != nil {
			return 0, err
		}
		return exp, nil
	}
	return 0, nil
}

func MakeTokenWithUppercase(token string) string {
	var sb strings.Builder
	for _, char := range token {
		// 随机决定是否将当前字符转换为大写
		if isHexLowercase(char) && shouldUppercase() {
			char = hexUppercase(char)
		}
		sb.WriteRune(char)
	}
	return sb.String()
}

func isHexLowercase(char rune) bool {
	return char >= 'a' && char <= 'f'
}

func hexUppercase(char rune) rune {
	return char - 'a' + 'A'
}

func shouldUppercase() bool {
	// 例如，这里我们有 50% 的机会将字符转换为大写
	n, err := rand.Int(rand.Reader, big.NewInt(2))
	if err != nil {
		// 在实践中，rand.Int 很少会出错，但如果出错，我们选择不进行转换
		return false
	}
	return n.Int64() == 1
}
