package api

import (
	//import gin

	"backend/db"
	cookie "backend/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func Setup() *gin.Engine {
	router = gin.Default()
	user := router.Group("/user")
	{
		user.POST("/login", login)
		user.POST("/register", register)
		user.GET("/:username", showUser)             //.Use(isAuthenticated())
		user.GET("/:username/sock", listSocksOfUser) //.Use(isAuthenticated())
	}

	sock := router.Group("/sock")
	{
		sock.POST("/", addSock)                        //.Use(isAuthenticated())
		sock.GET("/:sockId/match", listMatchesOfSock)  //.Use(isAuthenticated())
		sock.PATCH("/:sockId/", patchAcceptListOfSock) //.Use(isAuthenticated())
		sock.GET("/:sockId", getSockInfo)
	}

	return router
}

func getSockInfo(c *gin.Context) {
	c.Next()
}

func patchAcceptListOfSock(c *gin.Context) {
	c.Next()
}
func showUser(c *gin.Context) {
	c.Next()
}

func listSocksOfUser(c *gin.Context) {
	c.Next()
}

func addSock(c *gin.Context) {
	c.Next()

}

func listMatchesOfSock(c *gin.Context) {
	c.Next()
}

func isAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		session_cookie, err := c.Cookie("session")
		if err != nil {
			c.AbortWithStatus(401)
		}

		_, err = db.CheckCookie(session_cookie)
		if err != nil {
			c.AbortWithStatus(401)
		}
		c.Next()
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

	u, err := db.VerifyLogin(usr, pwd)
	if err != nil {
		c.JSON(401, gin.H{
			"message": "wrong username or password",
		})
		return
	}
	log.Printf("user: %v\n", u)

	c.JSON(200, gin.H{
		"message": "login successful",
	})

}

// create the login function
func login(c *gin.Context) {
	type TmpLogin struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	//retrieve the username and the password from the post data
	tmpLogin := TmpLogin{}
	err := c.BindJSON(&tmpLogin)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "login failed",
		})
		return
	}
	//check if the username and password are correct

	_, err = db.VerifyLogin(tmpLogin.Username, tmpLogin.Password)
	if err != nil {
		//if they are not correct, return an error message
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "login failed",
		})
		return
	}

	//if they are correct, return a success message
	//TODO - add a token to the response
	ck := cookie.GenSessionCookie(c)
	db.SetCookie(ck, tmpLogin.Username)

	c.Status(http.StatusOK)
	c.JSON(http.StatusOK, gin.H{
		"message": "login sucessful",
	})

}

// create the register function
func register(c *gin.Context) {
	type TmpUser struct {
		Username        string `json:"username"`
		Firstname       string `json:"firstname"`
		Surname         string `json:"surname"`
		ShippingAddress string `json:"shippingAddress"`
		Password        string `json:"password"`
	}
	tmpUser := TmpUser{}
	err := c.BindJSON(&tmpUser)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})
		return
	}

	_, err = db.RegisterUser(tmpUser.Username, tmpUser.Password, tmpUser.Firstname, tmpUser.Surname, tmpUser.ShippingAddress)

	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"message": err.Error(),
		})
		return
	}
}
