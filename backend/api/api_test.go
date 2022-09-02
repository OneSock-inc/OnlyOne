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
	jwtToken := loginUser1()
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
	jwtToken := loginUser1()
	w := httptest.NewRecorder()
	req := newSockRequest(0, 0, "#FFF", "In a very good state", getValidBase64Image())
	req.Header["Authorization"] = []string{fmt.Sprintf(`Bearer %s`, jwtToken)}
	router.ServeHTTP(w, req)
	log.Printf("%s", w.Body.String())
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestAddSockBadType(t *testing.T) {
	jwtToken := loginUser1()
	w := httptest.NewRecorder()
	req := newSockRequest(42, 2, "#FFF", "In a very good state", getValidBase64Image())
	req.Header["Authorization"] = []string{fmt.Sprintf(`Bearer %s`, jwtToken)}
	router.ServeHTTP(w, req)
	log.Printf("%s", w.Body.String())
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestAddSockBadColor(t *testing.T) {
	jwtToken := loginUser1()

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
	jwtToken := loginUser1()
	w := httptest.NewRecorder()
	req := newSockRequest(42, 0, "#FFF", "", getValidBase64Image())
	req.Header["Authorization"] = []string{fmt.Sprintf(`Bearer %s`, jwtToken)}
	router.ServeHTTP(w, req)
	log.Printf("%s", w.Body.String())
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestAddSockBadBase64(t *testing.T) {
	jwtToken := loginUser1()
	w := httptest.NewRecorder()
	req := newSockRequest(42, 0, "#FFF", "Magnificent !", "")
	req.Header["Authorization"] = []string{fmt.Sprintf(`Bearer %s`, jwtToken)}
	router.ServeHTTP(w, req)
	log.Printf("%s", w.Body.String())
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func registerUser(username string, pwd string) bool {
	res := httptest.NewRecorder()
	req := newUserRegisterRequest(username, pwd)
	router.ServeHTTP(res, req)
	log.Printf("%s", res.Body)
	return res.Code == http.StatusCreated
}

func loginUser(username string, pwd string) string {
	// create the user
	registerUser(username, pwd)

	// login said user
	js := fmt.Sprintf(`{"username":"%s","password":"%s"}`, username, pwd)
	r := strings.NewReader(js)
	req := httptest.NewRequest("POST", "/user/login", r)
	loginResponse := httptest.NewRecorder()
	router.ServeHTTP(loginResponse, req)
	body := loginResponse.Body.String()
	log.Printf("%s", body)

	type result struct {
		Code   uint16 `json:"code"`
		Expire string `json:"expire"`
		Token  string `json:"token"`
	}
	// get the token
	fmt.Printf("%s", body)
	var jsonResult result
	if err := json.Unmarshal([]byte(body), &jsonResult); err != nil {
		return err.Error()
	}
	log.Print("token :" + jsonResult.Token)
	return jsonResult.Token
}

func loginUser1() string {
	return loginUser("sockMan", "a12345678")
}

func loginUser2() string {
	return loginUser("sockWoman", "a12345678")
}

func loginUser3() string {
	return loginUser("sockGuy", "petit-poireau")
}

func TestCreateUser_Login_AddSock(t *testing.T) {
	jwtToken := loginUser1()
	log.Printf("%s", jwtToken)
	w := httptest.NewRecorder()
	req := newSockRequest(42, 0, "#FFF", "In a very bad state. Also smells", getValidBase64Image())
	req.Header["Authorization"] = []string{fmt.Sprintf(`Bearer %s`, jwtToken)}
	router.ServeHTTP(w, req)
	log.Printf("%s", w.Body.String())
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestShowUser(t *testing.T) {
	jwtToken := loginUser1()
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

func sendPatchRequest(jwt string, sockID string, otherSockID string, accept bool) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req := newPatchRequest(sockID, otherSockID, true)
	req.Header["Authorization"] = []string{fmt.Sprintf(`Bearer %s`, jwt)}
	router.ServeHTTP(w, req)
	return w
}

func TestPatchAcceptListOfSock_UserAcceptsTheSameSock(t *testing.T) {
	jwt1 := loginUser1()
	sock := createSock(t, jwt1)

	w := sendPatchRequest(jwt1, sock.ID, sock.ID, true)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	sock, err := db.GetSockInfo(sock.ID)
	assert.Nil(t, err)
	assert.Equal(t, make([]string, 0), sock.AcceptedList)
	assert.Equal(t, make([]string, 0), sock.RefusedList)
	assert.Empty(t, sock.Match)
}

func TestPatchAcceptListOfSock_UserAcceptsSockHeOwns(t *testing.T) {
	jwt1 := loginUser1()
	sock1 := createSock(t, jwt1)
	sock2 := createSock(t, jwt1)

	w := sendPatchRequest(jwt1, sock1.ID, sock2.ID, true)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	sock1, err := db.GetSockInfo(sock1.ID)
	assert.Nil(t, err)
	assert.Equal(t, make([]string, 0), sock1.AcceptedList)
	assert.Equal(t, make([]string, 0), sock1.RefusedList)
	assert.Empty(t, sock1.Match)

	sock2, err = db.GetSockInfo(sock2.ID)
	assert.Nil(t, err)
	assert.Equal(t, make([]string, 0), sock2.AcceptedList)
	assert.Equal(t, make([]string, 0), sock2.RefusedList)
	assert.Empty(t, sock2.Match)
}

func TestPatchAcceptListOfSock_UserRefusesSockHeOwns(t *testing.T) {
	jwt1 := loginUser1()
	sock1 := createSock(t, jwt1)
	sock2 := createSock(t, jwt1)

	w := sendPatchRequest(jwt1, sock1.ID, sock2.ID, false)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	sock1, err := db.GetSockInfo(sock1.ID)
	assert.Nil(t, err)
	assert.Equal(t, make([]string, 0), sock1.AcceptedList)
	assert.Equal(t, make([]string, 0), sock1.RefusedList)
	assert.Empty(t, sock1.Match)

	sock2, err = db.GetSockInfo(sock2.ID)
	assert.Nil(t, err)
	assert.Equal(t, make([]string, 0), sock2.AcceptedList)
	assert.Equal(t, make([]string, 0), sock2.RefusedList)
	assert.Empty(t, sock2.Match)
}

func TestPatchAcceptListOfSock_WrongUserAcceptsSock(t *testing.T) {
	jwt1 := loginUser1()
	jwt2 := loginUser2()
	sockUsr1A := createSock(t, jwt1)
	sockUsr1B := createSock(t, jwt1)
	sockUsr2 := createSock(t, jwt2)

	assertNoChanges := func() {
		for _, s := range []string{sockUsr1A.ID, sockUsr1B.ID, sockUsr2.ID} {
			sock, err := db.GetSockInfo(s)
			assert.Nil(t, err)
			assert.Equal(t, make([]string, 0), sock.AcceptedList)
			assert.Equal(t, make([]string, 0), sock.RefusedList)
			assert.Empty(t, sock.Match)
		}
	}

	// sockUsr1A accepts sockUsr2 (User2 doesn't own sockUsr1A)
	w := sendPatchRequest(jwt2, sockUsr1A.ID, sockUsr2.ID, true)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assertNoChanges()

	// sockUsr1A refuses sockUsr2 (User2 doesn't own sockUsr1A)
	w = sendPatchRequest(jwt2, sockUsr1A.ID, sockUsr2.ID, false)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assertNoChanges()

	// sockUsr1A accepts sockUsr1B (User2 doesn't own sockUsr1A)
	w = sendPatchRequest(jwt2, sockUsr1A.ID, sockUsr2.ID, true)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assertNoChanges()

	// sockUsr1A refuses sockUsr1B (User2 doesn't own sockUsr1A)
	w = sendPatchRequest(jwt2, sockUsr1A.ID, sockUsr2.ID, false)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assertNoChanges()
}

func TestPatchAcceptListOfSock_ValidAcceptThenReacceptShouldFail(t *testing.T) {
	jwt1 := loginUser1()
	jwt2 := loginUser2()

	sockUsr1 := createSock(t, jwt1)
	sockUsr2 := createSock(t, jwt2)

	assertCorrectAcceptList := func() {
		sockUsr1, err := db.GetSockInfo(sockUsr1.ID)
		assert.Nil(t, err)
		assert.Equal(t, []string{sockUsr2.ID}, sockUsr1.AcceptedList)
		assert.Equal(t, make([]string, 0), sockUsr1.RefusedList)
		assert.Empty(t, sockUsr1.Match)

		sockUsr2, err = db.GetSockInfo(sockUsr2.ID)
		assert.Nil(t, err)
		assert.Equal(t, make([]string, 0), sockUsr2.AcceptedList)
		assert.Equal(t, make([]string, 0), sockUsr2.RefusedList)
		assert.Empty(t, sockUsr2.Match)
	}

	// sockUsr1 accepts sockUsr2 (valid)
	w := sendPatchRequest(jwt1, sockUsr1.ID, sockUsr2.ID, true)
	assert.Equal(t, http.StatusOK, w.Code)
	assertCorrectAcceptList()

	// sockUsr1 accepts sockUsr2 (already in the acceptList)
	w = sendPatchRequest(jwt1, sockUsr1.ID, sockUsr2.ID, true)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assertCorrectAcceptList()

	// sockUsr1 accepts sockUsr2 (Accept then refuse)
	w = sendPatchRequest(jwt1, sockUsr1.ID, sockUsr2.ID, false)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assertCorrectAcceptList()
}

func TestPatchAcceptListOfSock_ValidAcceptThenRefuseShouldFail(t *testing.T) {
	jwt1 := loginUser1()
	jwt2 := loginUser2()

	sockUsr1 := createSock(t, jwt1)
	sockUsr2 := createSock(t, jwt2)

	assertCorrectAcceptList := func() {
		sockUsr1, err := db.GetSockInfo(sockUsr1.ID)
		assert.Nil(t, err)
		assert.Equal(t, []string{sockUsr2.ID}, sockUsr1.AcceptedList)
		assert.Equal(t, make([]string, 0), sockUsr1.RefusedList)
		assert.Empty(t, sockUsr1.Match)

		sockUsr2, err = db.GetSockInfo(sockUsr2.ID)
		assert.Nil(t, err)
		assert.Equal(t, make([]string, 0), sockUsr2.AcceptedList)
		assert.Equal(t, make([]string, 0), sockUsr2.RefusedList)
		assert.Empty(t, sockUsr2.Match)
	}

	// sockUsr1 accepts sockUsr2 (valid)
	w := sendPatchRequest(jwt1, sockUsr1.ID, sockUsr2.ID, true)
	assert.Equal(t, http.StatusOK, w.Code)
	assertCorrectAcceptList()

	// sockUsr1 accepts sockUsr2 (accept then refuse)
	w = sendPatchRequest(jwt1, sockUsr1.ID, sockUsr2.ID, false)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assertCorrectAcceptList()
}

func TestPatchAcceptListOfSock_Match(t *testing.T) {
	jwt1 := loginUser1()
	jwt2 := loginUser2()

	sockUsr1 := createSock(t, jwt1)
	sockUsr2 := createSock(t, jwt2)

	// sockUsr1 accepts sockUsr2
	w := sendPatchRequest(jwt1, sockUsr1.ID, sockUsr2.ID, true)
	assert.Equal(t, http.StatusOK, w.Code)

	sockUsr1, err := db.GetSockInfo(sockUsr1.ID)
	assert.Nil(t, err)
	assert.Equal(t, []string{sockUsr2.ID}, sockUsr1.AcceptedList)
	assert.Equal(t, make([]string, 0), sockUsr1.RefusedList)
	assert.Empty(t, sockUsr1.Match)

	sockUsr2, err = db.GetSockInfo(sockUsr2.ID)
	assert.Nil(t, err)
	assert.Equal(t, make([]string, 0), sockUsr2.AcceptedList)
	assert.Equal(t, make([]string, 0), sockUsr2.RefusedList)
	assert.Empty(t, sockUsr2.Match)

	// sockUsr2 accepts sockUsr1
	w = sendPatchRequest(jwt2, sockUsr2.ID, sockUsr1.ID, true)
	assert.Equal(t, http.StatusOK, w.Code)

	sockUsr1, err = db.GetSockInfo(sockUsr1.ID)
	assert.Nil(t, err)
	assert.Equal(t, []string{sockUsr2.ID}, sockUsr1.AcceptedList)
	assert.Equal(t, make([]string, 0), sockUsr1.RefusedList)
	assert.Equal(t, sockUsr2.ID, sockUsr1.Match)

	sockUsr2, err = db.GetSockInfo(sockUsr2.ID)
	assert.Nil(t, err)
	assert.Equal(t, []string{sockUsr1.ID}, sockUsr2.AcceptedList)
	assert.Equal(t, make([]string, 0), sockUsr2.RefusedList)
	assert.Equal(t, sockUsr1.ID, sockUsr2.Match)
}

func TestPatchAcceptListOfSock_MatchThenShouldFail(t *testing.T) {
	jwt1 := loginUser1()
	jwt2 := loginUser2()
	jwt3 := loginUser3()

	sockUsr1 := createSock(t, jwt1)
	sockUsr2 := createSock(t, jwt2)
	sockUsr3 := createSock(t, jwt3)

	assertCorrectMatch := func() {
		sockUsr1, err := db.GetSockInfo(sockUsr1.ID)
		assert.Nil(t, err)
		assert.Equal(t, []string{sockUsr2.ID}, sockUsr1.AcceptedList)
		assert.Equal(t, make([]string, 0), sockUsr1.RefusedList)
		assert.Equal(t, sockUsr2.ID, sockUsr1.Match)

		sockUsr2, err = db.GetSockInfo(sockUsr2.ID)
		assert.Nil(t, err)
		assert.Equal(t, []string{sockUsr1.ID}, sockUsr2.AcceptedList)
		assert.Equal(t, make([]string, 0), sockUsr2.RefusedList)
		assert.Equal(t, sockUsr1.ID, sockUsr2.Match)

		sockUsr3, err := db.GetSockInfo(sockUsr3.ID)
		assert.Nil(t, err)
		assert.Equal(t, make([]string, 0), sockUsr3.AcceptedList)
		assert.Equal(t, make([]string, 0), sockUsr3.RefusedList)
		assert.Empty(t, sockUsr3.Match)
	}

	// sockUsr1 accepts sockUsr2
	w := sendPatchRequest(jwt1, sockUsr1.ID, sockUsr2.ID, true)
	assert.Equal(t, http.StatusOK, w.Code)

	sockUsr1, err := db.GetSockInfo(sockUsr1.ID)
	assert.Nil(t, err)
	assert.Equal(t, []string{sockUsr2.ID}, sockUsr1.AcceptedList)
	assert.Equal(t, make([]string, 0), sockUsr1.RefusedList)
	assert.Empty(t, sockUsr1.Match)

	sockUsr2, err = db.GetSockInfo(sockUsr2.ID)
	assert.Nil(t, err)
	assert.Equal(t, make([]string, 0), sockUsr2.AcceptedList)
	assert.Equal(t, make([]string, 0), sockUsr2.RefusedList)
	assert.Empty(t, sockUsr2.Match)

	// sockUsr2 accepts sockUsr1
	w = sendPatchRequest(jwt2, sockUsr2.ID, sockUsr1.ID, true)
	assert.Equal(t, http.StatusOK, w.Code)
	assertCorrectMatch()

	// sockUsr2 accepts sockUsr1 (a match already exists between them)
	w = sendPatchRequest(jwt2, sockUsr2.ID, sockUsr1.ID, true)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assertCorrectMatch()

	// sockUsr3 accepts sockUsr1 (sockUsr1 already has a match)
	w = sendPatchRequest(jwt3, sockUsr3.ID, sockUsr1.ID, true)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assertCorrectMatch()

	// sockUsr1 accepts sockUsr3 (sockUsr1 already has a match)
	w = sendPatchRequest(jwt1, sockUsr1.ID, sockUsr3.ID, true)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assertCorrectMatch()

	// sockUsr3 refuses sockUsr1 (sockUsr1 already has a match)
	w = sendPatchRequest(jwt3, sockUsr3.ID, sockUsr1.ID, false)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assertCorrectMatch()

	// sockUsr1 refuses sockUsr3 (sockUsr1 already has a match)
	w = sendPatchRequest(jwt1, sockUsr1.ID, sockUsr3.ID, false)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assertCorrectMatch()
}
