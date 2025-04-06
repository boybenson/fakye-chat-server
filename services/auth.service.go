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

		return response, nil
	}
}