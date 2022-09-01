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
		"firstname": "first",
		"surname": "surname",
		"address": {
			"street": "rue du rhone 1",
			"country": "Swiss",
			"city": "Genève",
			"postalCode": "1212"
		},
		"password": "%s"
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

func TestListSocksOfUser(t *testing.T) {
	jwtToken := makeLogedinUser1()
	log.Printf("%s", jwtToken)
	w := httptest.NewRecorder()

	req := newSockRequest(42, db.Profile(0), "#FFF", "Do not", getValidBase64Image())
	req.Header["Authorization"] = []string{fmt.Sprintf(`Bearer %s`, jwtToken)}
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var sock db.Sock

	json.Unmarshal(w.Body.Bytes(), &sock)
	w = httptest.NewRecorder()

	req = httptest.NewRequest("GET", "/user/sockMan/sock", nil)
	req.Header["Authorization"] = []string{fmt.Sprintf(`Bearer %s`, jwtToken)}
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `[{"id":"`+
		sock.ID+
		`","shoeSize":42,"type":0,"color":"#FFF","description":"Do not","picture":"aHR0cHM6Ly9kbGFuZy5vcmcK","owner":"`+
		sock.Owner+
		`","refusedList":null,"acceptedList":null,"match":""}]`, w.Body.String())
}

func TestAddSockWithoutUser(t *testing.T) {
	w := httptest.NewRecorder()
	req := newSockRequest(42, db.Profile(1), "#FFFFFF", "i used to wank in this one", getValidBase64Image())
	router.ServeHTTP(w, req)
	log.Printf("%s", w.Body.String())
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestAddSockBadShoeSize(t *testing.T) {
	jwtToken := makeLogedinUser1()
	w := httptest.NewRecorder()
	req := newSockRequest(0, 0, "#FFF", "In a very good state", getValidBase64Image())
	req.Header["Authorization"] = []string{fmt.Sprintf(`Bearer %s`, jwtToken)}
	router.ServeHTTP(w, req)
	log.Printf("%s", w.Body.String())
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestAddSockBadType(t *testing.T) {
	jwtToken := makeLogedinUser1()
	w := httptest.NewRecorder()
	req := newSockRequest(42, 2, "#FFF", "In a very good state", getValidBase64Image())
	req.Header["Authorization"] = []string{fmt.Sprintf(`Bearer %s`, jwtToken)}
	router.ServeHTTP(w, req)
	log.Printf("%s", w.Body.String())
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestAddSockBadColor(t *testing.T) {
	jwtToken := makeLogedinUser1()

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
	jwtToken := makeLogedinUser1()
	w := httptest.NewRecorder()
	req := newSockRequest(42, 0, "#FFF", "", getValidBase64Image())
	req.Header["Authorization"] = []string{fmt.Sprintf(`Bearer %s`, jwtToken)}
	router.ServeHTTP(w, req)
	log.Printf("%s", w.Body.String())
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestAddSockBadBase64(t *testing.T) {
	jwtToken := makeLogedinUser1()
	w := httptest.NewRecorder()
	req := newSockRequest(42, 0, "#FFF", "Magnificent !", "")
	req.Header["Authorization"] = []string{fmt.Sprintf(`Bearer %s`, jwtToken)}
	router.ServeHTTP(w, req)
	log.Printf("%s", w.Body.String())
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func makeLogedinUser(username string) string {
	//creat user
	res := httptest.NewRecorder()
	pwd := "onlySock"
	req := newUserRegisterRequest(username, pwd)
	router.ServeHTTP(res, req)
	log.Printf("%s", res.Body)

	//login said user
	js := fmt.Sprintf(`{"username":"%s","password":"%s"}`, username, pwd)
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

func makeLogedinUser1() string {
	return makeLogedinUser("sockMan")
}

func makeLogedinUser2() string {
	return makeLogedinUser("sockWoman")
}

func TestCreateUser_Login_AddSock(t *testing.T) {

	jwtToken := makeLogedinUser1()
	log.Printf("%s", jwtToken)
	w := httptest.NewRecorder()
	req := newSockRequest(42, 0, "#FFF", "In a very bad state. Also smells", getValidBase64Image())
	req.Header["Authorization"] = []string{fmt.Sprintf(`Bearer %s`, jwtToken)}
	router.ServeHTTP(w, req)
	log.Printf("%s", w.Body.String())
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestShowUser(t *testing.T) {
	jwtToken := makeLogedinUser1()
	log.Printf("%s", jwtToken)
	w := httptest.NewRecorder()

	req := httptest.NewRequest("GET", "/user/invalid", nil)
	req.Header["Authorization"] = []string{fmt.Sprintf(`Bearer %s`, jwtToken)}
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	w = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/user/sockMan", nil)
	req.Header["Authorization"] = []string{fmt.Sprintf(`Bearer %s`, jwtToken)}
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"username":"sockMan","firstname":"first","surname":"surname","password":"","address":{"street":"rue du rhone 1","country":"Swiss","city":"Genève","postalCode":"1212"}}`, w.Body.String())
}

func newPatchRequest(sockID string, otherSockID string, accept bool) *http.Request {
	status := "refuse"
	if accept {
		status = "accept"
	}

	r := strings.NewReader(fmt.Sprintf(`{
		"status" : "%s",
		"otherSockID" : "%s"
	  }`, status, otherSockID))
	return httptest.NewRequest("PATCH", "/sock/"+sockID, r)
}

func createSock(t *testing.T, jwt string) db.Sock {
	w := httptest.NewRecorder()
	req := newSockRequest(37, 1, "#00BEEF", "A very ugly sock. Gonna cut my eyes into pieces", getValidBase64Image())
	req.Header["Authorization"] = []string{fmt.Sprintf(`Bearer %s`, jwt)}
	router.ServeHTTP(w, req)
	var sock db.Sock
	err := json.Unmarshal(w.Body.Bytes(), &sock)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, w.Code)

	return sock
}

func TestPatchAcceptListOfSock(t *testing.T) {
	jwt1 := makeLogedinUser1()
	jwt2 := makeLogedinUser2()

	// create 4 socks
	sock1User1 := createSock(t, jwt1)
	sock2User1 := createSock(t, jwt1)
	sock1User2 := createSock(t, jwt2)
	sock2User2 := createSock(t, jwt2)

	// sock1User1 accepts itself
	w := httptest.NewRecorder()
	req := newPatchRequest(sock1User1.ID, sock1User1.ID, true)
	req.Header["Authorization"] = []string{fmt.Sprintf(`Bearer %s`, jwt1)}
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// sock1User1 refuses sock2User1
	w = httptest.NewRecorder()
	req = newPatchRequest(sock1User1.ID, sock2User1.ID, false)
	req.Header["Authorization"] = []string{fmt.Sprintf(`Bearer %s`, jwt1)}
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// sock1User1 accepts sock1User2
	w = httptest.NewRecorder()
	req = newPatchRequest(sock1User1.ID, sock1User2.ID, true)
	req.Header["Authorization"] = []string{fmt.Sprintf(`Bearer %s`, jwt1)}
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// sock1User1 accepts sock1User2 (second time)
	w = httptest.NewRecorder()
	req = newPatchRequest(sock1User1.ID, sock1User2.ID, true)
	req.Header["Authorization"] = []string{fmt.Sprintf(`Bearer %s`, jwt1)}
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// sock1User1 refuses sock1User2 (accept then refuse)
	w = httptest.NewRecorder()
	req = newPatchRequest(sock1User1.ID, sock1User2.ID, false)
	req.Header["Authorization"] = []string{fmt.Sprintf(`Bearer %s`, jwt1)}
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// User2 accepts sock2User1 with sock1User1 (doesn't own it)
	w = httptest.NewRecorder()
	req = newPatchRequest(sock1User1.ID, sock2User1.ID, true)
	req.Header["Authorization"] = []string{fmt.Sprintf(`Bearer %s`, jwt2)}
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// sock1User2 accepts sock1User1
	w = httptest.NewRecorder()
	req = newPatchRequest(sock1User2.ID, sock1User1.ID, true)
	req.Header["Authorization"] = []string{fmt.Sprintf(`Bearer %s`, jwt2)}
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// sock2User2 accepts sock1User1 (a match already exists)
	w = httptest.NewRecorder()
	req = newPatchRequest(sock2User2.ID, sock1User1.ID, true)
	req.Header["Authorization"] = []string{fmt.Sprintf(`Bearer %s`, jwt2)}
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// sock2User2 accepts sock1User1 (a match already exists)
	w = httptest.NewRecorder()
	req = newPatchRequest(sock2User2.ID, sock1User1.ID, true)
	req.Header["Authorization"] = []string{fmt.Sprintf(`Bearer %s`, jwt2)}
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
