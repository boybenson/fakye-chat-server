package helpers

import (
	"bytes"
	"encoding/json"
	"fakye-server/types"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
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

func ModerateTexts (text string) bool{
	apiKey := os.Getenv("OPENAI_API_KEY")

	 openAIEndpoint := "https://api.openai.com/v1/moderations"

	reqBody := types.ModerationRequest{
		Model: "text-moderation-latest", 
		Input: text,
	}

	jsonData, err := json.Marshal(reqBody)

	if err != nil {
		panic(err)
	}


	req, err := http.NewRequest("POST", openAIEndpoint, bytes.NewBuffer(jsonData))

	if err != nil {
		panic(err)
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	fmt.Println("Moderation Response:")
	fmt.Println(string(body))
	
	return true
}