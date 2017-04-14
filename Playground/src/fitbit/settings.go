package fitbit

import (
	"golang.org/x/oauth2"
	"model"
)

func init(){

}
var conf_fitbit *oauth2.Config


func Config() oauth2.Config{

	var c_fitbit model.Credentials
	c_fitbit.Cid = "227TT7"
	c_fitbit.Csecret = "98aa8d159b62dec86958ef6a38e79115"

	var endpoint_fitbit = oauth2.Endpoint{}
	endpoint_fitbit.AuthURL = "https://www.fitbit.com/oauth2/authorize"
	endpoint_fitbit.TokenURL = "https://api.fitbit.com/oauth2/token"


	conf_fitbit = &oauth2.Config{
		ClientID:     c_fitbit.Cid,
		ClientSecret: c_fitbit.Csecret,
		RedirectURL:  "http://localhost:8000/callBack",
		Scopes: []string{"profile", "activity", "sleep", "weight" },
		Endpoint: endpoint_fitbit,
	}

	return  *conf_fitbit
}