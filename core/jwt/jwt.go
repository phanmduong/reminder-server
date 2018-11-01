package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
	"reflect"
	"reminder/core/config"
	"time"
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

func myDecoder(val *reflect.Value, data interface{}) (interface{}, error) {
	if val.Type().String() == "time.Time" {
		value, err := time.Parse(time.RFC3339Nano, data.(string))
		val.Set(reflect.ValueOf(value))
		return nil, err
	}
	return data, nil
}

func getDecoder(result interface{}) (*mapstructure.Decoder, error) {
	return mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName:          "json",
		Result:           result,
		WeaklyTypedInput: false})
}

func ParseJWT(tokenString string, structData interface{}) (err error) {

	appConfig := config.GetAppConfig().Config
	secretJWT := []byte(appConfig["SECRET_JWT"])
	claims := myClaims{}

	_, err = jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return secretJWT, nil
	})

	decoder, err := getDecoder(structData)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = decoder.Decode(claims.Data)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
