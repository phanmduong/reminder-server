package field

import (
	"encoding/json"
	"github.com/graphql-go/graphql"
	"net/http"
	"reminder/core/service"
	"reminder/graphql/gqltype"
	"reminder/model"
	"time"
)

var getProfileFacebook = func(faceID string, token string) (model.User) {
	var myClient = &http.Client{Timeout: 10 * time.Second}
	data := model.User{}

	req, err := myClient.Get("https://graph.facebook.com/v3.1/" + string(faceID) + "?access_token=" + token + "&fields=id,name,email")
	if err != nil {
		println("errror2")
	}

	defer req.Body.Close()

	json.NewDecoder(req.Body).Decode(&data)

	println(data.Name)
	return data
}

var Login = &graphql.Field{
	Type:        gqltype.UserType,
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

		var user model.User

		fb_id, ok := p.Args["fb_id"].(string)
		token, ok := p.Args["fb_id"].(string)
		if ok {
			//db.Where("fb_id = ?", fb_id).First(&user)
			return getProfileFacebook(fb_id, token), nil
		}

		google_id, ok := p.Args["google_id"].(int)
		if ok {
			db.Where("google_id = ?", google_id).First(&user)
			return user, nil
		}
		return nil, nil
	},
}
