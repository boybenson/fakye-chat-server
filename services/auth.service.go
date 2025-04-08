package services

import (
	"fakye-server/helpers"
	"fakye-server/models"
	"fakye-server/types"
	"fmt"

	"gorm.io/gorm"
)



func Register(payload types.RegisterRequest, db *gorm.DB)(*models.User, error) {
	user := models.User{Name: payload.Name, Phone: payload.Phone}

	var existingUser models.User
	userExists := db.Where("phone = ?", payload.Phone).First(&existingUser)

	if userExists.Error == nil {
		return nil, fmt.Errorf("user already exists with phone number %s", payload.Phone)
	} else {
		result := db.Create(&user)
		if result.Error != nil {
			return nil, result.Error
		}

		message := fmt.Sprintf("Hello %v, Welcome to Benevo. Share freely and brighten someone's day with https://benevoghana.com", payload.Name)
		
		go DispatchSms(message, payload.Phone)

		return &user, nil
	}

}

func SignIn (payload types.SignInRequest, db *gorm.DB)(*types.UserWithToken, error){
	var existingUser models.User
	userExists := db.Where("phone = ?", payload.Phone).First(&existingUser)

	if userExists.Error != nil {
		return nil, fmt.Errorf("incorrect phone number %s", payload.Phone)
	} else {
		token,_:= helpers.GenerateAuthToken(existingUser.Phone)
		response := &types.UserWithToken{
			User:  existingUser,
			Token: token,
		}

		otpCode := helpers.GenerateOtpCode()

		result := db.Model(&models.User{}).Where("phone = ?", payload.Phone).Update("auth_otp", otpCode)

		if result.Error != nil {
			return nil, result.Error
		}

		println(otpCode)

		return response, nil
	}
}