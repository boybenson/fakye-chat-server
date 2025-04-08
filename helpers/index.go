package helpers

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateAuthToken (phone string) (string, error){

	claims := jwt.MapClaims{
		"phone": phone,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), 
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("TEST"))

	if err != nil {
		return "", err
	}
	return signedToken, nil
}


func GenerateOtpCode () string{
rand.Seed(time.Now().UnixNano())

otp := rand.Intn(9000) + 1000 


return fmt.Sprintf("%d", otp)
}