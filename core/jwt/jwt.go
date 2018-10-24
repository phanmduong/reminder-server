package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
	"reminder/core/config"
)

type myClaims struct {
	Data interface{} `json:"data"`
	jwt.StandardClaims
}

const EXPIRES_AT = 10080

func CreateJWT(data interface{}) (token string, error error) {
	appConfig := config.GetAppConfig().Config
	secretJWT := []byte(appConfig["SECRET_JWT"])

	claims := myClaims{
		data,
		jwt.StandardClaims{
			ExpiresAt: EXPIRES_AT,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	token, err := tokenClaims.SignedString(secretJWT)
	return token, err
}

func ParseJWT(tokenString string, stuctData interface{}) (err error) {

	appConfig := config.GetAppConfig().Config
	secretJWT := []byte(appConfig["SECRET_JWT"])
	claims := myClaims{}

	_, err = jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return secretJWT, nil
	})
	//if err == nil && token.Valid {
	//	return &claims.Data, nil
	//}
	mapstructure.Decode(claims.Data, stuctData)
	return err

}
