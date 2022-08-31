package api

import (
	"backend/db"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Setenv("FIRESTORE_EMULATOR_HOST", "localhost:8080")
	client, _ := db.GetDBConnection()
	err := db.DeleteCollection(context.Background(), client, client.Collection("users"), 64)
	if err != nil {
		log.Print(err.Error())
		os.Exit(1)
	}
	log.SetFlags(log.Flags() | log.Llongfile)
	Setup()
	exit := m.Run()

	os.Exit(exit)

}

func TestSetup(t *testing.T) {

	if router == nil {
		t.Errorf("engine is nil which should be impossible")
	}
}

func TestRoutes(t *testing.T) {

	routes := router.Routes()
	if routes == nil {
		t.Errorf("routes is nil which should be impossible")
	}
	// log.Printf("routes: %v\n", routes)

}
func TestLoginFail(t *testing.T) {

	w := httptest.NewRecorder()
	r := strings.NewReader(
		`{
		"username":"mojojo",
		"password":"mySecret"
	}`)
	req, err := http.NewRequest("POST", "/user/login", r)
	assert.Nil(t, err)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	log.Printf("%s", w.Body)
}
func newUserRegisterRequest(usr string, pwd string) *http.Request {
	js := fmt.Sprintf(`{
		"username": "%s",
		"firstname": "firstname",
		"surname" : "surname",
		"password" : "%s",
		"shippingAddress" : "Planet earth, 3301"
	}`, usr, pwd)
	r := strings.NewReader(js)
	req, err := http.NewRequest("POST", "/user/register", r)
	if err != nil {
		log.Fatal(err.Error())
	}
	return req
}
func TestRegisterSuccess(t *testing.T) {
	w := httptest.NewRecorder()
	req := newUserRegisterRequest("username", "passwd")
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}
func TestIllFormedLoginFail(t *testing.T) {
	js := fmt.Sprintf(`
	{"pwd": "%s"
	"usr" : "%s" `, "pwd", "usr")
	req, err := http.NewRequest("POST", "/user/login", strings.NewReader(js))
	assert.Nil(t, err)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	log.Printf("%s", w.Body)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestIllFormed2LoginFail(t *testing.T) {
	js := fmt.Sprintf(`
	{"password": "%s"}`, "pwd")
	req, err := http.NewRequest("POST", "/user/login", strings.NewReader(js))
	assert.Nil(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	log.Printf("%s", w.Body)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestLoginSucess(t *testing.T) {

	//delete all users
	client, err := db.GetDBConnection()
	assert.Nil(t, err)
	err = db.DeleteCollection(context.Background(), client, client.Collection("users"), 64)
	assert.Nil(t, err)

	const user = "hisUsername"
	const pwd = "hisPassword"
	//Creat a user in the db (might not exist)
	w := httptest.NewRecorder()
	req := newUserRegisterRequest(user, pwd)
	router.ServeHTTP(w, req)
	log.Printf("%s", w.Body)
	assert.Equal(t, http.StatusCreated, w.Code)
	//login with said user
	w = httptest.NewRecorder()
	r := strings.NewReader(fmt.Sprintf(
		`{
			"username": "%s", 
			"password":"%s"
		}`, user, pwd))
	req, err = http.NewRequest("POST", "/user/login", r)
	assert.Nil(t, err)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
func newSockRequest(shoeSize uint16, type_ db.Profile, color string, descr string, picture string) *http.Request {
	r := strings.NewReader(fmt.Sprintf(`{
	"shoeSize": %d,
	"type": %d,
	"color": "%s",
	"description": "%s",
	"picture":"%s"}`, shoeSize, type_, color, descr, picture))
	return httptest.NewRequest("POST", "/sock/", r)
}

func getValidBase64Image() string {
	return "aHR0cHM6Ly9kbGFuZy5vcmcK"
}

func TestAddSockWithoutUser(t *testing.T) {
	w := httptest.NewRecorder()
	req := newSockRequest(42, db.Profile(1), "#FFFFFF", "i used to wank in this one", getValidBase64Image())
	router.ServeHTTP(w, req)
	log.Printf("%s", w.Body.String())
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestAddSockBadShoeSize(t *testing.T) {
	jwtToken := makeLogedinUser()
	w := httptest.NewRecorder()
	req := newSockRequest(0, 0, "#FFF", "In a very good state", getValidBase64Image())
	req.Header["Authorization"] = []string{fmt.Sprintf(`Bearer %s`, jwtToken)}
	router.ServeHTTP(w, req)
	log.Printf("%s", w.Body.String())
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestAddSockBadType(t *testing.T) {
	jwtToken := makeLogedinUser()
	w := httptest.NewRecorder()
	req := newSockRequest(42, 2, "#FFF", "In a very good state", getValidBase64Image())
	req.Header["Authorization"] = []string{fmt.Sprintf(`Bearer %s`, jwtToken)}
	router.ServeHTTP(w, req)
	log.Printf("%s", w.Body.String())
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestAddSockBadColor(t *testing.T) {
	jwtToken := makeLogedinUser()

	w := httptest.NewRecorder()
	req := newSockRequest(42, 0, "", "In a very good state", getValidBase64Image())
	req.Header["Authorization"] = []string{fmt.Sprintf(`Bearer %s`, jwtToken)}
	router.ServeHTTP(w, req)
	log.Printf("%s", w.Body.String())
	assert.Equal(t, http.StatusBadRequest, w.Code)

	req = newSockRequest(42, 1, "junk", "In a very good state", getValidBase64Image())
	req.Header["Authorization"] = []string{fmt.Sprintf(`Bearer %s`, jwtToken)}
	router.ServeHTTP(w, req)
	log.Printf("%s", w.Body.String())
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestAddSockBadDescription(t *testing.T) {
	jwtToken := makeLogedinUser()
	w := httptest.NewRecorder()
	req := newSockRequest(42, 0, "#FFF", "", getValidBase64Image())
	req.Header["Authorization"] = []string{fmt.Sprintf(`Bearer %s`, jwtToken)}
	router.ServeHTTP(w, req)
	log.Printf("%s", w.Body.String())
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestAddSockBadBase64(t *testing.T) {
	jwtToken := makeLogedinUser()
	w := httptest.NewRecorder()
	req := newSockRequest(42, 0, "#FFF", "Magnificent !", "")
	req.Header["Authorization"] = []string{fmt.Sprintf(`Bearer %s`, jwtToken)}
	router.ServeHTTP(w, req)
	log.Printf("%s", w.Body.String())
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
func makeLogedinUser() string {
	//creat user
	res := httptest.NewRecorder()
	sockMan := "sockMan"
	pwd := "onlySock"
	req := newUserRegisterRequest(sockMan, pwd)
	router.ServeHTTP(res, req)
	log.Printf("%s", res.Body)

	//login said user
	js := fmt.Sprintf(`{"username":"%s","password":"%s"}`, sockMan, pwd)
	r := strings.NewReader(js)
	req = httptest.NewRequest("POST", "/user/login", r)
	loginResponse := httptest.NewRecorder()
	router.ServeHTTP(loginResponse, req)
	body := loginResponse.Body.String()
	log.Printf("%s", body)

	type result struct {
		Code   uint16 `json:"code"`
		Expire string `json:"expire"`
		Token  string `json:"token"`
	}
	//get cookie session
	fmt.Printf("%s", body)
	var jsonResult result
	if err := json.Unmarshal([]byte(body), &jsonResult); err != nil {
		return err.Error()
	}
	log.Print("token :" + jsonResult.Token)
	return jsonResult.Token

}
func TestCreateUser_Login_AddSock(t *testing.T) {

	jwtToken := makeLogedinUser()
	log.Printf("%s", jwtToken)
	w := httptest.NewRecorder()
	req := newSockRequest(42, 0, "#FFF", "In a very bad state. Also smells", getValidBase64Image())
	req.Header["Authorization"] = []string{fmt.Sprintf(`Bearer %s`, jwtToken)}
	router.ServeHTTP(w, req)
	log.Printf("%s", w.Body.String())
	assert.Equal(t, http.StatusCreated, w.Code)
}
