package middlewares

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/MilliGoshant/merci/backend/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var JWT_SECRET = []byte(os.Getenv("JWT_SECRET"))

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Claims struct {
	User models.User `json:"user"`
	jwt.RegisteredClaims
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := &Claims{}
		var token string

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "Token is required")
			return
		}
		splitToken := strings.Split(authHeader, "Bearer ")
		if len(splitToken) > 1 {
			token = splitToken[1]
		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, "Invalid token")
			return
		}
		tkn, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
			return JWT_SECRET, nil
		})
		if err != nil {

			if strings.Contains(err.Error(), "token is expired") {
				c.AbortWithStatusJSON(http.StatusForbidden, "Token expired")
				return
			}

			c.AbortWithStatusJSON(http.StatusBadRequest, "Cannot parse token")
			return

		}

		if claims.ExpiresAt.Unix() < time.Now().Local().Unix() {
			c.AbortWithStatusJSON(http.StatusForbidden, "Token expired")
			return
		}

		if !tkn.Valid {
			c.AbortWithStatusJSON(http.StatusBadRequest, "Invalid token")
			return
		}
		c.Set("user_id", claims.User.ID)
		c.Next()
	}
}
