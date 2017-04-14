package main

import (

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"net/http"
	"context"
	"io/ioutil"
	"encoding/json"
	"log"
	"fitbit"
	"github.com/gin-gonic/contrib/sessions"
	"model"
	"fmt"
)

//type Credentials struct {
//	Cid string `json:"cid"`
//	Csecret string `json:"csecret"`
//}



var conf_fitbit *oauth2.Config
//var state string
var client *http.Client


func init(){
	//conf_fitbit = &fitbit.Config()

	var c_fitbit model.Credentials
	c_fitbit.Cid = "227TT7"
	c_fitbit.Csecret = "98aa8d159b62dec86958ef6a38e79115"
	//
	//
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

}

func getFitibitLoginURL() string {
	return conf_fitbit.AuthCodeURL("state")
}

func main(){

	router := gin.Default()
	store := sessions.NewCookieStore([]byte("secret"))

	router.Use(sessions.Sessions("mysession", store))

	router.StaticFS("/sports", http.Dir("C:\\Users\\OskAlb\\GIT\\GoPlayground\\Playground\\html\\sports"))
	router.Use(gin.Logger())

	router.GET("/", indexHandler)
	router.GET("/login", loginHandler)
	router.GET("/callBack", authHandler)
	router.GET("/api/user", userHandler)
	router.GET("/api/activities", activitiesHandler)
	router.GET("/api/activity/:key", activityHandler)

	router.Run("localhost:8000")
}

func indexHandler(c *gin.Context){
	c.HTML(http.StatusOK, "index.html", nil)
}


func loginHandler(c *gin.Context){
	c.Writer.Write([]byte("<html><title>Sports</title> <body> <a href='" + getFitibitLoginURL() + "'><button>Login with Fitbit!</button> </a></body></html>"))
}

func authHandler(c *gin.Context){
	session := sessions.Default(c)

	code := c.Query("code")
	token, err := conf_fitbit.Exchange(context.TODO(), code)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		c.Writer.WriteString(err.Error())
		return
	}

	ctx := context.WithValue(context.Background(), oauth2.HTTPClient, &http.Client{

	})

	client = conf_fitbit.Client(ctx, token)

	session.Set("sessionId", model.RandStringRunes(8))
	session.Save()


	c.Writer.WriteString(string("Connected to Fitbit"))

}

func userHandler(c *gin.Context){
	session := sessions.Default(c)
	sessionId := session.Get("sessionId")

	fmt.Println(sessionId)

	resp, err := client.Get("https://api.fitbit.com/1/user/-/profile.json")
	if err != nil {
		c.HTML(http.StatusServiceUnavailable, "error.html", nil)
		return
	}

	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)

	user := fitbit.UserContainer{}
	err2 := json.Unmarshal([]byte(data), &user)

	if err2 != nil {
		c.HTML(http.StatusServiceUnavailable,"error.html", nil)
		log.Fatal(err2)
	}

	c.JSON(http.StatusOK, user)

}

func activitiesHandler(c *gin.Context){

	resp, err := client.Get("https://api.fitbit.com/1/user/-/activities/list.json?beforeDate=2017-04-07&sort=asc&limit=20&offset=0")
	if err != nil {
		c.HTML(http.StatusServiceUnavailable, "error.html", nil)
		return
	}

	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)


	ac := fitbit.ActivityContainer{}

	err2 := json.Unmarshal([]byte(data), &ac)

	if err2 != nil {
		log.Fatal(err2)
	}

	c.JSON(http.StatusOK, ac)

}

func activityHandler(c *gin.Context){
	key := c.Param("key")
	resp, err := client.Get("https://api.fitbit.com/1/activities/" + key + ".json")
	if err != nil{
		c.HTML(http.StatusServiceUnavailable, "error.html", nil)
		return
	}

	defer resp.Body.Close()
	data, _:=ioutil.ReadAll(resp.Body)

	c.Writer.WriteString(string(data))

}

