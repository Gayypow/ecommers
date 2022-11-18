package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/MilliGoshant/merci/backend/db"
	"github.com/MilliGoshant/merci/backend/models"

	"github.com/MilliGoshant/merci/backend/converter"
	jwtpack "github.com/MilliGoshant/merci/backend/jwt"
	"github.com/MilliGoshant/merci/backend/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// Register hashes the password of user and saves it into Users with unique name
func Register(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := user.HashPassword(user.Password); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := db.DB.Omit("id").Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id":   user.ID,
		"name": user.Name,
	})
}

// Login finds the user with unique name, compares its password with given password and responses access_token, refresh_token
func Login(c *gin.Context) {
	var loginRequest LoginRequest
	var user models.User

	if err := c.BindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := db.DB.Where("name = ?", loginRequest.Name).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	if err := user.CheckPassword(loginRequest.Password); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	tokens, err := jwtpack.GenerateJWT(user.ID, user.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, tokens)
}

func GetUser(c *gin.Context) {
	var user models.User
	userID, hasUserID := c.Get("user_id")
	if !hasUserID {
		c.JSON(http.StatusBadRequest, "user_id is required")
		return
	}
	if err := db.DB.Omit("password").First(&user, "id=?", userID).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":   user.ID,
		"name": user.Name,
	})
}

func UpdateUser(c *gin.Context) {
	var user models.User
	userID, hasUserID := c.Get("user_id")
	if !hasUserID {
		c.JSON(http.StatusBadRequest, "user_id is required")
		return
	}
	if err := db.DB.First(&user, "id=?", userID).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	user.ID = converter.StrToUint(fmt.Sprintf("%v", userID))
	if err := db.DB.Omit("password").Save(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":   user.ID,
		"name": user.Name,
	})
}

func Logout(c *gin.Context) {
	c.JSON(http.StatusNoContent, "Logout successfully")
}

// RefreshToken
func Token(c *gin.Context) {
	claims := &middlewares.Claims{}
	token := jwtpack.Tokens{}
	if err := c.BindJSON(&token); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	tkn, err := jwt.ParseWithClaims(token.RefreshToken, claims, func(t *jwt.Token) (interface{}, error) {
		return middlewares.JWT_SECRET, nil
	})
	if err != nil {
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
	// ! Response with middleware
	tokens, err := RefreshToken(claims)
	if err != nil {
		c.JSON(http.StatusForbidden, err.Error())
		return
	}
	c.JSON(http.StatusOK, tokens)
}

func RefreshToken(claims *middlewares.Claims) (token middlewares.Tokens, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)

	claims.ExpiresAt = &jwt.NumericDate{Time: expirationTime}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token.AccessToken, err = accessToken.SignedString(middlewares.JWT_SECRET)
	if err != nil {
		return middlewares.Tokens{}, err
	}
	expirationTime = time.Now().Add(24 * time.Hour * 30)

	claims.ExpiresAt = &jwt.NumericDate{Time: expirationTime}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token.RefreshToken, err = refreshToken.SignedString(middlewares.JWT_SECRET)
	if err != nil {
		return middlewares.Tokens{}, err
	}

	return token, nil
}
