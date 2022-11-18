package jwt

import (
	"os"
	"time"

	"github.com/MilliGoshant/merci/backend/models"

	"github.com/golang-jwt/jwt/v4"
)

var JWT_SECRET = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	User models.User `json:"user"`
	jwt.RegisteredClaims
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// GenerateJWT creates access and refresh tokens with user's id, name
func GenerateJWT(id uint, name string) (token Tokens, err error) {
	// Create access token
	accessTokenExp := time.Now().Add(3 * time.Hour)
	accessClaims := &Claims{
		User: models.User{
			ID:   id,
			Name: name,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: accessTokenExp},
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	token.AccessToken, err = accessToken.SignedString(JWT_SECRET)
	if err != nil {
		return Tokens{}, err
	}

	// Create refresh token
	refreshTokenExp := time.Now().Add(24 * time.Hour * 30)

	refreshClaims := &Claims{
		User: models.User{
			ID:   id,
			Name: name,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: refreshTokenExp},
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	token.RefreshToken, err = refreshToken.SignedString(JWT_SECRET)
	if err != nil {
		return Tokens{}, err
	}

	return token, nil
}
