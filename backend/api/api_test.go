package api

import (
	"backend/db"
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"cloud.google.com/go/firestore"
	"github.com/stretchr/testify/assert"
	"google.golang.org/api/iterator"
)

func TestMain(m *testing.M) {
	os.Setenv("FIRESTORE_EMULATOR_HOST", "localhost:8080")
	client, _ := db.GetDBConnection()
	err := deleteCollection(context.Background(), client, client.Collection("users"), 64)
	if err != nil {
		log.Print(err.Error())
		os.Exit(1)
	}
	log.SetFlags(log.Flags() | log.Llongfile)
	Setup()
	exit := m.Run()

	os.Exit(exit)

}

func deleteCollection(ctx context.Context, client *firestore.Client,
	ref *firestore.CollectionRef, batchSize int) error {

	for {
		// Get a batch of documents
		iter := ref.Limit(batchSize).Documents(ctx)
		numDeleted := 0

		// Iterate through the documents, adding
		// a delete operation for each one to a
		// WriteBatch.
		batch := client.Batch()
		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return err
			}

			batch.Delete(doc.Ref)
			numDeleted++
		}

		// If there are no documents to delete,
		// the process is over.
		if numDeleted == 0 {
			return nil
		}

		_, err := batch.Commit(ctx)
		if err != nil {
			return err
		}
	}
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
	err = deleteCollection(context.Background(), client, client.Collection("users"), 64)
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
	session := makeLogedinUser()
	w := httptest.NewRecorder()
	req := newSockRequest(0, 0, "#FFF", "In a very good state", getValidBase64Image())
	req.AddCookie(session)
	router.ServeHTTP(w, req)
	log.Printf("%s", w.Body.String())
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestAddSockBadType(t *testing.T) {
	session := makeLogedinUser()
	w := httptest.NewRecorder()
	req := newSockRequest(42, 2, "#FFF", "In a very good state", getValidBase64Image())
	req.AddCookie(session)
	router.ServeHTTP(w, req)
	log.Printf("%s", w.Body.String())
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestAddSockBadColor(t *testing.T) {
	cookie := makeLogedinUser()

	w := httptest.NewRecorder()
	req := newSockRequest(42, 0, "", "In a very good state", getValidBase64Image())
	req.AddCookie(cookie)
	router.ServeHTTP(w, req)
	log.Printf("%s", w.Body.String())
	assert.Equal(t, http.StatusBadRequest, w.Code)

	req = newSockRequest(42, 1, "junk", "In a very good state", getValidBase64Image())
	req.AddCookie(cookie)
	router.ServeHTTP(w, req)
	log.Printf("%s", w.Body.String())
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestAddSockBadDescription(t *testing.T) {
	session := makeLogedinUser()
	w := httptest.NewRecorder()
	req := newSockRequest(42, 0, "#FFF", "", getValidBase64Image())
	req.AddCookie(session)
	router.ServeHTTP(w, req)
	log.Printf("%s", w.Body.String())
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestAddSockBadBase64(t *testing.T) {
	session := makeLogedinUser()
	w := httptest.NewRecorder()
	req := newSockRequest(42, 0, "#FFF", "Magnificent !", "")
	req.AddCookie(session)
	router.ServeHTTP(w, req)
	log.Printf("%s", w.Body.String())
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
func makeLogedinUser() *http.Cookie {
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
	log.Printf("%s", loginResponse.Body)

	//get cookie session
	var c *http.Cookie = nil
	//find session cookie from login
	for _, cookie := range loginResponse.Result().Cookies() {
		if cookie.Name == "session" {
			c = cookie

		}
	}
	return c

}
func TestCreateUser_Login_AddSock(t *testing.T) {

	cookie := makeLogedinUser()
	w := httptest.NewRecorder()
	req := newSockRequest(42, 0, "#FFF", "In a very bad state. Also smells", getValidBase64Image())
	req.Host = "localhost"
	req.AddCookie(cookie)
	router.ServeHTTP(w, req)
	log.Printf("%s", w.Body.String())
	assert.Equal(t, http.StatusCreated, w.Code)
}
