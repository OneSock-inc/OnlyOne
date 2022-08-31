package api

import (
	//import gin

	"backend/db"
	"encoding/json"
	"log"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func jwtSetup() *jwt.GinJWTMiddleware {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:      "realm doesn't make sens in the JWT context",
		Key:        []byte("This is the secret key used to sign the identity, hope it doesn't leak ;)"),
		Timeout:    time.Hour * time.Duration(8760),
		MaxRefresh: time.Hour * time.Duration(8760),
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(string); ok {
				return jwt.MapClaims{
					jwt.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		Authenticator: login,

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	return authMiddleware
}
func Setup() *gin.Engine {
	router = gin.Default()
	auth := jwtSetup()
	user := router.Group("/user")
	{
		user.POST("/login", auth.LoginHandler)
		user.POST("/register", register)
		user.GET("/:username", showUser)
		user.GET("/:username/sock", listSocksOfUser)
	}

	sock := router.Group("/sock").Use(auth.MiddlewareFunc())
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
	claim := jwt.ExtractClaims(c)
	userID, ok := claim[jwt.IdentityKey].(string)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "User is not authentificated",
		})
		return
	}

	socks, err := db.GetUserSocks(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{
			"message": err.Error(),
		})
		return
	}

	j, err := json.Marshal(socks)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, j)
}

func addSock(c *gin.Context) {

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

	claim := jwt.ExtractClaims(c)

	if userID, ok := claim[jwt.IdentityKey].(string); ok {

		_, err = db.NewSock(tmpSock.ShoeSize, tmpSock.Type, tmpSock.Color, tmpSock.Description, tmpSock.Picture, userID)
		if err != nil {

			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})

			return
		}
		log.Printf("user %s added sock %+v \n", userID, tmpSock)
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Sock successfully added !",
	})

}

func listMatchesOfSock(c *gin.Context) {
	c.Next()
}

// create the login function
func login(c *gin.Context) (interface{}, error) {
	type TmpLogin struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	//retrieve the username and the password from the post data
	tmpLogin := TmpLogin{}
	if err := c.BindJSON(&tmpLogin); err != nil {
		return "", err
	}
	//check if the username and password are correct

	id, err := db.VerifyLogin(tmpLogin.Username, tmpLogin.Password)
	if err != nil {
		log.Printf("login failed %+v\n", tmpLogin)
		return "", err
	}
	log.Printf("user %s logged in\n", id)
	return id, nil
}

// create the register function
func register(c *gin.Context) {

	tmpUser := db.User{}
	log.Printf("new user signing up :")
	if err := c.BindJSON(&tmpUser); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})
		return
	}
	log.Printf("%+v\n", tmpUser)

	_, err := db.RegisterUser(tmpUser)

	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "registration successful"})
}
