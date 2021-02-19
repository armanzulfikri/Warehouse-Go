package web

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

var memToken *jwt.Token

func authorization(context *gin.Context) {
	var (
		tokenFromAuth string
		method        string
		path          string
		secret        string
	)
	secret = os.Getenv("SECRET_TOKEN")
	tokenFromAuth = context.Request.Header.Get("Authorization")
	method = context.Request.Method
	path = context.FullPath()

	token, err := jwt.Parse(tokenFromAuth, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Method false $e", token.Header["alg"])
		}

		return []byte(secret), nil
	})

	if token != nil && err == nil {
		payloads := token.Claims.(jwt.MapClaims)
		if payloads["role"] == "supervisor" && path == "/auth/register" {
			result := gin.H{
				"message": "Unauthorized",
				"error":   http.StatusUnauthorized,
			}

			context.Abort()
			context.JSON(http.StatusUnauthorized, result)
		} else if payloads["role"] == "entry" && (method == "DELETE" || path == "/auth/register") {
			result := gin.H{
				"message": "Unauthorized",
				"error":   http.StatusUnauthorized,
			}

			context.Abort()
			context.JSON(http.StatusUnauthorized, result)
		} else if payloads["role"] == "manajemen" && method != "GET" {
			result := gin.H{
				"message": "Unauthorized",
				"error":   http.StatusUnauthorized,
			}

			context.Abort()
			context.JSON(http.StatusUnauthorized, result)
		} else {
			memToken = token
			context.Next()
		}
	} else {
		result := gin.H{
			"message": "Token false",
			"error":   http.StatusUnauthorized,
		}

		context.Abort()
		context.JSON(http.StatusUnauthorized, result)
	}
}

//GetRole
func GetRole() string {
	payloads := memToken.Claims.(jwt.MapClaims)

	return payloads["role"].(string)
}
