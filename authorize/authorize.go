package authorize

import (
	"reminder/core/jwt"
	"reminder/core/service"
	"reminder/model"
)

func CheckAuthorization(token string) (isValid bool, userResult model.User) {
	var user = model.User{}

	jwt.ParseJWT(token, &user)
	db := service.GetService().DB.DB
	var userDB = model.User{}
	db.Debug().Where("fb_id = ? or google_id = ?", user.FbID, user.GoogleID).Where("fb_id IS NOT NULL OR google_id IS NOT NULL").First(&userDB)
	//db.Debug().Where("fb_id = ? or google_id = ?", 1105521586288749, user.GoogleID).Where("fb_id IS NOT NULL OR google_id IS NOT NULL").First(&userDB)

	return userDB.ID != 0, userDB
}
