package field

import (
	"encoding/json"
	"github.com/graphql-go/graphql"
	"net/http"
	"reminder/core/jwt"
	"reminder/core/service"
	"reminder/graphql/gqltype"
	"reminder/model"
	"time"
)

type UserFb struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	ID    string `json:"id"`
}

var getProfileFacebook = func(faceID string, token string) (model.User, error) {
	var myClient = &http.Client{Timeout: 10 * time.Second}
	data := UserFb{}
	user := model.User{}

	var url = "https://graph.facebook.com/v3.2/" + faceID + "?access_token=" + token + "&fields=id,name,email"
	req, err := myClient.Get(url)
	if err != nil {
		println("errror2")
	}

	defer req.Body.Close()

	json.NewDecoder(req.Body).Decode(&data)

	user.FbID = data.ID
	user.Name = data.Name
	user.Email = data.Email
	user.AvatarURL = "https://graph.facebook.com/" + faceID + "/picture??height=500&width=500"
	return user, err
}

var Login = &graphql.Field{
	Type:        gqltype.LoginType,
	Description: "Get user by id",
	Args: graphql.FieldConfigArgument{
		"fb_id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"google_id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		db := service.GetService().DB.DB

		//var user model.User
		token, ok := p.Args["token"].(string)
		fb_id, ok := p.Args["fb_id"].(string)

		if ok {
			//db.Where("fb_id = ?", fb_id).First(&user)
			user, error := getProfileFacebook(fb_id, token);
			if error != nil {
				return nil, error
			}
			db.Where(model.User{FbID: user.FbID}).FirstOrCreate(&user, model.User{FbID: user.FbID})
			token, err := jwt.CreateJWT(&user)

			return token, err
		}

		//google_id, ok := p.Args["google_id"].(int)
		//if ok {
		//	db.Where("google_id = ?", google_id).First(&user)
		//	return user, nil
		//}
		return nil, nil
	},
}
