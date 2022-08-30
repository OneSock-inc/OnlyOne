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

// use in tests
const pwd = "myPwd"
const usr = "myUsername"

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
	req, _ := http.NewRequest("POST", "/user/login", r)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	log.Printf("%s", w.Body)
}
func newRegisterRequest() *http.Request {
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
	req := newRegisterRequest()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
func TestIllFormedLoginFail(t *testing.T) {
	js := fmt.Sprintf(`
	{"pwd": "%s",
	"usr" : "%s"} `, pwd, usr)
	req, _ := http.NewRequest("POST", "/user/login", strings.NewReader(js))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	log.Printf("%s", w.Body)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestIllFormed2LoginFail(t *testing.T) {
	js := fmt.Sprintf(`
	{"password": "%s"}`, pwd)
	req, _ := http.NewRequest("POST", "/user/login", strings.NewReader(js))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	log.Printf("%s", w.Body)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestLoginSucess(t *testing.T) {

	//delete all users
	client, _ := db.GetDBConnection()
	_ = deleteCollection(context.Background(), client, client.Collection("users"), 64)

	//Creat a user in the db (might not exist)
	w := httptest.NewRecorder()
	req := newRegisterRequest()
	router.ServeHTTP(w, req)
	log.Printf("%s", w.Body)
	assert.Equal(t, http.StatusOK, w.Code)
	//login with said user
	w = httptest.NewRecorder()
	r := strings.NewReader(fmt.Sprintf(
		`{
			"username": "%s", 
			"password":"%s"
		}`, usr, pwd))
	req, err := http.NewRequest("POST", "/user/login", r)
	assert.Nil(t, err)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
