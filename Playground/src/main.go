package main

import (

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"net/http"
	"log"
	"context"
)

type Credentials struct {
	Cid string `json:"cid"`
	Csecret string `json:"csecret"`
}

var c Credentials
var conf_trakt *oauth2.Config
var conf_strava *oauth2.Config
var state string

func init(){
	var c_trakt  Credentials

	c_trakt.Cid = "b8e22e27eeceab630753ea058c2e7a18c69f6eef24c0fd85b3da699c14b9f70b"
	c_trakt.Csecret = "c1579194bae558083290084b5ad2f7df76c148712cd8b73e0f6359f89655b36a"


	var c_strava Credentials
	c_strava.Cid = "12447"
	c_strava.Csecret = "ed03ec4d23aa39a78f94433d5d2c62bc25b2c44e"

	var endpoint_trakt oauth2.Endpoint
	endpoint_trakt.AuthURL = "https://api-staging.trakt.tv/oauth/authorize"
	endpoint_trakt.TokenURL = "https://api-staging.trakt.tv/oauth/token"

	var endpoint_strava = oauth2.Endpoint{}
	endpoint_strava.AuthURL = "https://www.strava.com/oauth/authorize"
	endpoint_strava.TokenURL = "https://www.strava.com/oauth/token"

	conf_trakt = &oauth2.Config{
		ClientID:     c_trakt.Cid,
		ClientSecret: c_trakt.Csecret,
		RedirectURL:  "http://localhost:8000/callBack",
		Scopes: []string{

		},
		Endpoint: endpoint_trakt,
	}

	conf_strava = &oauth2.Config{
		ClientID:     c_strava.Cid,
		ClientSecret: c_strava.Csecret,
		RedirectURL:  "http://localhost:8000/callBack",
		Scopes: []string{ },
		Endpoint: endpoint_strava,
	}
}

func getLoginURL(state string) string {
	return conf_strava.AuthCodeURL(state)
}

func main(){
	router := gin.Default()

	router.Use(gin.Logger())

	router.GET("/index", gettingHandler)
	router.GET("/callBack", authHandler)
	//router := NewRouter2()
	router.Run("localhost:8000")
}

func gettingHandler(c *gin.Context){
	c.Writer.Write([]byte("<html><title>Golang Strava</title> <body> <a href='" + getLoginURL(state) + "'><button>Login with Strava!</button> </a> </body></html>"))
}

func authHandler(c *gin.Context){
	token, err := conf_strava.Exchange(context.TODO(), c.Query("code"))

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx := context.WithValue(context.Background(), oauth2.HTTPClient, &http.Client{

	})
	httpClient := conf_strava.Client(ctx, token)
	//svc, err := urlshortener.New(httpClient)
	//client := conf_strava.Client(context.TODO(), token)

//	shows, err := client.Get("https://www.strava.com/api/v3/athlete")
	atlhete, err := httpClient.Get("https://www.strava.com/api/v3/athlete")

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	log.Println(atlhete)


}


