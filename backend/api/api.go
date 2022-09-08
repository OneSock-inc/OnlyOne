package api

import (
	//import gin

	"backend/db"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "net/http/pprof"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	jwtgo "github.com/golang-jwt/jwt/v4"
)

var router *gin.Engine

const msg = "message"

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

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

func getIdentity(jwtstr string) (string, error) {
	secret := []byte("This is the secret key used to sign the identity, hope it doesn't leak ;)")
	token, err := jwtgo.Parse(jwtstr, func(token *jwtgo.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwtgo.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return secret, nil
	})
	if err != nil {
		return "", err
	}

	claims, _ := token.Claims.(jwtgo.MapClaims)
	s := claims[jwt.IdentityKey].(string)
	return s, nil
}

func Setup() *gin.Engine {
	router = gin.Default()
	router.Use(CORSMiddleware())
	auth := jwtSetup()
	router.GET("/userid/:userid", auth.MiddlewareFunc(), showUserFromUserId)
	user := router.Group("/user")
	{
		user.PATCH("/update", auth.MiddlewareFunc(), updateUser)
		user.PATCH("/update/", auth.MiddlewareFunc(), updateUser)
		user.POST("/login", auth.LoginHandler)
		user.POST("/register", register)
		user.GET("/:username", auth.MiddlewareFunc(), showUser)
		user.GET("/:username/sock", auth.MiddlewareFunc(), listSocksOfUser)
	}
	router.NoRoute(CORSMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			msg: "404 not found",
		})
	})
	router.NoMethod(CORSMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			msg: fmt.Sprintf("%s is not allowed on %s", c.Request.Method, c.Request.URL.Path),
		})
	})

	//all these routes need a valide jwt
	sock := router.Group("/sock").Use(CORSMiddleware(), auth.MiddlewareFunc())
	{
		sock.POST("", addSock)
		sock.POST("/", func(c *gin.Context) {
			c.Redirect(http.StatusTemporaryRedirect, "/sock")
		})
		sock.GET("/:sockId/match", listMatchesOfSock)
		sock.PATCH("/:sockId", patchAcceptListOfSock)
		sock.GET("/:sockId", getSockInfo)
	}

	return router
}

func updateUser(c *gin.Context) {
	claim := jwt.ExtractClaims(c)
	userId, _ := claim[jwt.IdentityKey].(string)
	var user db.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			msg: err.Error(),
		})
		return
	}

	err = db.UpdateUser(userId, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			msg: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		msg: "updated",
	})

}

func getSockInfo(c *gin.Context) {

	sockId := c.Param("sockId")
	s, err := db.GetSockInfo(sockId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			msg: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, s)
}

func patchAcceptListOfSock(c *gin.Context) {
	claim := jwt.ExtractClaims(c)
	// checks already made by the middleware
	userID, _ := claim[jwt.IdentityKey].(string)

	sock, err := db.GetSockInfo(c.Param("sockId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			msg: err.Error(),
		})
		return
	}

	if sock.Owner != userID {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			msg: "User does not own sock ID `" + sock.ID + "`",
		})
		return
	}

	type TmpPatchReq struct {
		Status      string `json:"status"`
		OtherSockID string `json:"otherSockID"`
	}

	tmpPatch := TmpPatchReq{}
	err = c.BindJSON(&tmpPatch)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			msg: err.Error(),
		})
		return
	}
	var status bool
	if tmpPatch.Status == "accept" {
		status = true
	} else if tmpPatch.Status == "refuse" {
		status = false
	} else {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			msg: "Status is incorrect",
		})
		return
	}

	otherSock, err := db.GetSockInfo(tmpPatch.OtherSockID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			msg: "Other sock ID doesn't exist",
		})
		return
	}
	err = db.EditMatchingSock(sock, otherSock, status)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			msg: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		msg: "Success",
	})
}
func showUserFromUserId(c *gin.Context) {
	user, err := db.GetUserFromID(c.Param("userid"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			msg: err.Error(),
		})
		return
	}

	user.Password = ""
	c.JSON(http.StatusOK, user)

}
func showUser(c *gin.Context) {
	doc, err := db.GetUser(c.Param("username"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			msg: err.Error(),
		})
		return
	}
	var user db.User
	doc.DataTo(&user)
	// Do not show the user's hash !
	user.Password = ""
	c.JSON(http.StatusOK, user)
}

func listSocksOfUser(c *gin.Context) {
	claim := jwt.ExtractClaims(c)
	userID, ok := claim[jwt.IdentityKey].(string)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			msg: "User is not authentificated",
		})
		return
	}

	socks, err := db.GetUserSocks(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{
			msg: err.Error(),
		})
		return
	}
	if len(socks) == 0 {
		c.JSON(http.StatusNoContent, "[]")
		return
	}
	c.JSON(http.StatusOK, socks)
}

func addSock(c *gin.Context) {

	type TmpSock struct {
		ID          string     `json:"id"`
		ShoeSize    uint8      `json:"shoeSize"`
		Type        db.Profile `json:"type"`
		Color       string     `json:"color"`
		Description string     `json:"description"`
		Picture     string     `json:"picture"`
		Owner       string     `json:"owner"`
	}

	tmpSock := TmpSock{}
	err := c.BindJSON(&tmpSock)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			msg: err.Error(),
		})
		return
	}

	claim := jwt.ExtractClaims(c)
	userID, ok := claim[jwt.IdentityKey].(string)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			msg: "User not authentificated",
		})
		return
	}

	doc, err := db.NewSock(tmpSock.ShoeSize, tmpSock.Type, tmpSock.Color, tmpSock.Description, tmpSock.Picture, userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			msg: err.Error(),
		})

		return
	}
	tmpSock.ID = doc.ID
	tmpSock.Owner = userID
	log.Printf("user %s added sock %s \n", userID, tmpSock.ID)

	c.JSON(http.StatusCreated, gin.H{
		"id": tmpSock.ID,
	})
}

func listMatchesOfSock(c *gin.Context) {
	sockId := c.Param("sockId")
	socks, err := db.GetCompatibleSocks(sockId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			msg: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, socks)

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
		log.Printf("login failed for %s using %s\n", tmpLogin.Username, tmpLogin.Password)
		return "", err
	}
	log.Printf("user %s logged in\n", tmpLogin.Username)
	return id, nil
}

// create the register function
func register(c *gin.Context) {

	tmpUser := db.User{}
	if err := c.BindJSON(&tmpUser); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			msg: err.Error(),
		})
		return
	}

	_, err := db.RegisterUser(tmpUser)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			msg: err.Error(),
		})
		return
	}
	log.Printf("%s signed up", tmpUser.Username)
	c.JSON(http.StatusCreated, gin.H{msg: "registration successful"})
}
