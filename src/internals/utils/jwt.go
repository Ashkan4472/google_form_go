package utils

import (
	"os"
	"strconv"
	"time"

	"github.com/Ashkan4472/google_form_go/src/internals/models"
	"github.com/golang-jwt/jwt/v4"
)

func JWTGenerate(user models.User) (string, error) {
	jwtExpriationDays, _ := strconv.Atoi(os.Getenv("JWT_EXPORATION_TIME"))
	jwtExpriationTime := time.Duration(jwtExpriationDays) * time.Hour * 24;
	jwtExp := time.Now().Add(jwtExpriationTime).Unix()
	
	claims := jwt.MapClaims{
		"id": user.ID,
		"email": user.Email,
		"exp": jwtExp,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtSecret := os.Getenv("JWT_USER_SECRET")
	return token.SignedString([]byte(jwtSecret))
}
