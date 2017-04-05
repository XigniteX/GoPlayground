package main

import (

	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()
	router.GET("/someGet", gettingHandler)
	//router := NewRouter2()
	router.Run()
}

func gettingHandler(c *gin.Context){
	c.Writer.Write([]byte ("HI"))
}



