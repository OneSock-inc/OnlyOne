package api

import (
	//import gin

	"backend/db"
	"backend/utils"
	"fmt"
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
		user.GET("/:username", showUser)
		user.GET("/:username/sock", listSocksOfUser)
	}

	sock := router.Group("/sock").Use(isAuthenticated())
	{
		sock.POST("/", addSock)
		sock.GET("/:sockId/match", listMatchesOfSock)
		sock.PATCH("/:sockId/", patchAcceptListOfSock)
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
	if c.Keys["docID"] == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	type TmpSock struct {
		ShoeSize    uint8      `json:"shoeSize"`
		Type        db.Profile `json:"type"`
		Color       string     `json:"color"`
		Description string     `json:"description"`
		Picture     string     `json:"picture"`
	}
	tmpSock := TmpSock{}
	err := c.BindJSON(&tmpSock)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	userID := fmt.Sprintf("%s", c.Keys["docID"])
	_, err = db.NewSock(tmpSock.ShoeSize, tmpSock.Type, tmpSock.Color, tmpSock.Description, tmpSock.Picture, userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Sock successfully added !",
	})
}

func listMatchesOfSock(c *gin.Context) {
	c.Next()
}

func isAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		session_cookie, err := c.Cookie("session")
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		userRef, err := db.CheckCookie(session_cookie)
		if err != nil || userRef == nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set("docID", userRef.ID)
		c.Next()
	}
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
			"message": err.Error(),
		})
		return
	}
	//check if the username and password are correct

	_, err = db.VerifyLogin(tmpLogin.Username, tmpLogin.Password)
	if err != nil {
		//if they are not correct, return an error message
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "login failed 2",
		})
		return
	}

	//if they are correct, return a success message
	//TODO - add a token to the response
	ck := utils.GenSessionCookie(c)
	db.SetCookie(ck, tmpLogin.Username)

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
	c.JSON(http.StatusCreated, gin.H{"message": "registration successful"})
}
