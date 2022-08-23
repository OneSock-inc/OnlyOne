package api

import (
	//import gin

	"onlyOne/back/db"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func Setup() *gin.Engine {
	router = gin.Default()

	user := router.Group("/user")
	{
		user.POST("/login", login)
		user.POST("/register", register)
	}
	router.GET("/next", nextOne)
	router.GET("/test", testLogin)
	return router
}

func isAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// get the next data for the user
func nextOne(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func testLogin(c *gin.Context) {
	usr := c.Query("username")
	pwd := c.Query("password")
	db.CheckUser(usr, pwd)
	c.Next()

}

// create the login function
func login(c *gin.Context) {
	//retrieve the username and the password from the post data
	// username := c.PostForm("username")
	// pwd := c.PostForm("password")

	username := c.Query("username")
	pwd := c.Query("password")
	//check if the username and password are correct
	db.CheckUser(username, pwd)
	if username == "m" && pwd == "myPwd" {
		//if they are correct, return a success message
		//TODO - add a token to the response
		c.JSON(200, gin.H{
			"message": "login successful",
		})
	} else {
		//if they are not correct, return an error message
		c.JSON(401, gin.H{
			"message": "login failed",
		})
	}
}

// create the register function
func register(c *gin.Context) {
	c.Next()
}
