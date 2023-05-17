package jwt

import (
	"fmt"
	"hr/pkg/app"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			appG  = app.Gin{C: c}
			token = strings.Replace(appG.C.Request.Header.Get("Authorization"), "Bearer ", "", 1)
			err   = ValidateToken(token)
		)

		if err != nil {
			appG.Response(http.StatusUnauthorized, "Please check your credential and try again.")
			appG.C.Abort()
			return
		}
		appG.C.Next()
	}
}

func ValidateToken(token string) error {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil && strings.ToLower(err.Error()) == jwt.ErrTokenUsedBeforeIssued.Error() {
		return nil
	}
	return err
}

func GenerateToken() (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(60 * time.Minute).Unix(),
	})

	ss, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}
	return &ss, nil
}
