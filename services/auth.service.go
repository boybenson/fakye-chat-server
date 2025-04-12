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
		otpCode := helpers.GenerateOtpCode()
		var updatedUser models.User

		result := db.Model(&models.User{}).Where("phone = ?", payload.Phone).Update("auth_otp", otpCode)

		db.Where("phone = ?", payload.Phone).First(&updatedUser)
		token,_:= helpers.GenerateAuthToken(updatedUser.Phone)

		if result.Error != nil {
			return nil, result.Error
		}


		response := &types.UserWithToken{
			User:  updatedUser,
			Token: token,
		}

		message := fmt.Sprintf("Hello %v, Kindly use the OTP code %s to verify your phone number. Share freely and brighten someone's day with https://benevoghana.com", existingUser.Name, otpCode)

		go DispatchSms(message, payload.Phone)

		return response, nil
	}
}

func VerifyOtp(payload types.VerifyOtpRequest, db *gorm.DB)(bool,error){
	var existingUser models.User
	userExists := db.Where("phone = ?", payload.Phone).First(&existingUser)

	if userExists.Error != nil {
		return false, fmt.Errorf("incorrect phone number %s", payload.Phone)
	}else {

		if existingUser.AuthOtp != payload.Otp {
			return false, fmt.Errorf("incorrect Otp Code %s", payload.Otp)
		}
		return true, nil

	}
}