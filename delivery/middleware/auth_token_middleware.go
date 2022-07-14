package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type AuthHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var (
	ApplicationName  = "ApiWmbPos"
	JwtSigningMethod = jwt.SigningMethodHS256
	JwtSignatureKey  = []byte("WMB4P1")
)

type MyClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Email    string `json:"email"`
}

func AuthTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/login" {
			c.Next()
			fmt.Println("sss")
		} else {
			h := AuthHeader{}
			if err := c.ShouldBindHeader(&h); err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{
					"message": "Unauthrorized",
				})
				c.Abort()
				return
			}
			tokenString := strings.Replace(h.AuthorizationHeader, "Bearer ", "", -1)
			fmt.Println("token", tokenString)
			if tokenString == "" {
				c.JSON(http.StatusUnauthorized, gin.H{
					"message": "Unauthrorized",
				})
				c.Abort()
				return
			}
			fmt.Println(1)

			token, err := ParseToken(tokenString)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{
					"message": "Unauthrorized",
				})
				c.Abort()
				return
			}
			fmt.Println("token:", token)
			if token["iss"] == ApplicationName {
				c.Next()
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{
					"message": "token invalid",
				})
				c.Abort()
			}
		}

	}
}

func GenerateToken(userName string, email string) (string, error) {
	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer: ApplicationName,
		},
		Username: userName,
		Email:    email,
	}
	token := jwt.NewWithClaims(
		JwtSigningMethod,
		claims,
	)
	return token.SignedString(JwtSignatureKey)
}

func ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		method, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("signing method invalid")
		} else if method != JwtSigningMethod {
			return nil, fmt.Errorf("signing method invalid")
		}
		fmt.Println(ok)

		return JwtSignatureKey, nil
	})
	fmt.Println(err)

	claims, ok := token.Claims.(jwt.MapClaims)
	fmt.Println(ok, token.Valid)
	if !ok || !token.Valid {
		return nil, err
	}
	fmt.Println(2)
	return claims, nil
}
