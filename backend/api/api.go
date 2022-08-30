package api

import (
	//import gin

	"backend/db"
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
				log.Println("OK")

				return jwt.MapClaims{
					jwt.IdentityKey: v,
				}
			}
			log.Println("not ok")

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
	c.Next()
}

func addSock(c *gin.Context) {
	log.Printf("")

	if c.Keys["docID"] == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		log.Printf("Are we here ")

		return
	}
	log.Printf("")

	type TmpSock struct {
		ShoeSize    uint8      `json:"shoeSize"`
		Type        db.Profile `json:"type"`
		Color       string     `json:"color"`
		Description string     `json:"description"`
		Picture     string     `json:"picture"`
	}
	log.Printf("")

	tmpSock := TmpSock{}
	err := c.BindJSON(&tmpSock)
	if err != nil {
		log.Printf("")

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	log.Printf("")

	claim := jwt.ExtractClaims(c)
	log.Printf("%+v", claim)

	if userID, ok := claim[jwt.IdentityKey].(string); ok {
		log.Printf("%s", userID)

		_, err = db.NewSock(tmpSock.ShoeSize, tmpSock.Type, tmpSock.Color, tmpSock.Description, tmpSock.Picture, userID)
		if err != nil {
			log.Printf("")

			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			log.Printf("")

			return
		}
	}
	log.Printf("")

	c.JSON(http.StatusCreated, gin.H{
		"message": "Sock successfully added !",
	})
	log.Printf("")

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
		log.Printf("%s", err.Error())
		return "", err
	}
	//check if the username and password are correct

	id, err := db.VerifyLogin(tmpLogin.Username, tmpLogin.Password)
	if err != nil {
		return "", err
	}

	return id, nil
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
