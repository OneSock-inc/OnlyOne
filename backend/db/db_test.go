package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Setenv(FirestoreEmulatorHost, "localhost:8080")
	log.SetFlags(log.Flags() | log.Llongfile)

	// delete collection after run
	client, err := GetDBConnection()

	if err != nil {
		log.Print(err.Error())
		os.Exit(1)
	}
	err = DeleteCollection(context.Background(), client, client.Collection("users"), 64)
	if err != nil {
		log.Print(err.Error())
		os.Exit(1)
	}
	err = DeleteCollection(context.Background(), client, client.Collection("socks"), 64)
	if err != nil {
		log.Print(err.Error())
		os.Exit(1)
	}
	exit := m.Run()
	os.Exit(exit)
}

const FirestoreEmulatorHost = "FIRESTORE_EMULATOR_HOST"

func TestGetClient(t *testing.T) {

	_, err := GetDBConnection()
	if err != nil {
		t.Errorf("erro : %v", err)
	}
}
func TestCheckUser(t *testing.T) {
	username := "A username"
	pwd := `compl3 x Pwd!`
	usr := User{username,
		"first",
		"surname",
		pwd,
		Address{
			"Street name",
			"country",
			"City",
			"zip code",
		},
	}
	_, err := RegisterUser(usr)
	if err != nil {
		t.Errorf("unable to create a user\n%v", err)
	}
	_, err = VerifyLogin(username, pwd)
	if err != nil {
		t.Errorf("error user don't exist in mockup\n%v", err)
	}

}

func TestCreateUser(t *testing.T) {

	usr := User{"test",
		"Joris",
		"jsch",
		"myPassword", Address{
			"Street name",
			"country",
			"City",
			"zip code",
		},
	}

	_, err := RegisterUser(usr)
	if err != nil {
		t.Errorf(" user should not exist\n%v", err)
	}
	_, err = RegisterUser(usr)

	if err == nil {
		log.Printf("the user should exist : %v", err)
	}
}

func TestInvalidRegisterUser(t *testing.T) {
	usr := User{"test",
		"Joris",
		"Schaller",
		"myPassword", Address{
			"Street name",
			"country",
			"City",
			"zip code",
		},
	}
	fakeUsr := usr
	fakeUsr.Username = ""
	_, err := RegisterUser(fakeUsr)
	if err == nil {
		t.Errorf("Empty username should not be allowed\n")
	}
	fakeUsr = usr
	fakeUsr.Password = ""
	_, err = RegisterUser(fakeUsr)
	if err == nil {
		t.Errorf("Empty password should not be allowed\n")
	}
	fakeUsr = usr
	fakeUsr.Firstname = ""
	_, err = RegisterUser(fakeUsr)
	if err == nil {
		t.Errorf("Empty firstname should not be allowed\n")
	}
	fakeUsr = usr
	fakeUsr.Surname = ""
	_, err = RegisterUser(fakeUsr)
	if err == nil {
		t.Errorf("Empty surname should not be allowed\n")
	}
	fakeUsr = usr
	fakeUsr.Address = Address{}
	_, err = RegisterUser(fakeUsr)
	if err == nil {
		t.Errorf("Empty shipping address should not be allowed\n")
	}
}

func TestInvalidGetUserSocks(t *testing.T) {
	socks, err := GetUserSocks("invalid")
	assert.Nil(t, err)
	assert.Zero(t, len(socks))
}

func TestGetUserSocks(t *testing.T) {
	user := User{Username: "james2010", Password: "123", Firstname: "James", Surname: "Wow", Address: Address{Street: "Non", Country: "CH", City: "GE", PostalCode: "1212"}}
	doc, err := RegisterUser(user)
	assert.Nil(t, err)

	userID := doc.ID
	socks, err := GetUserSocks(userID)
	assert.Nil(t, err)
	assert.Zero(t, len(socks))

	doc, err = NewSock(10, low, "#000", "Super sock !", "aGkK", userID)
	assert.Nil(t, err)

	sockID := doc.ID
	socks, err = GetUserSocks(userID)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(socks))
	assert.Equal(t, sockID, socks[0].ID)
}

func TestGetUser(t *testing.T) {
	user := User{Username: "jamy", Password: "123", Firstname: "Jamy", Surname: "Yuy", Address: Address{Street: "Arf", Country: "CH", City: "Lausanne", PostalCode: "1000"}}
	doc, err := RegisterUser(user)
	assert.Nil(t, err)

	userID := doc.ID
	docr, err := GetUser("jamy")
	assert.Nil(t, err)
	assert.Equal(t, userID, docr.Ref.ID)

	docr, err = GetUser("invalid")
	assert.NotNil(t, err)
}

func TestGetUserFromID(t *testing.T) {
	user := User{Username: "ronron", Password: "123", Firstname: "Ron", Surname: "Ron", Address: Address{Street: "Rond", Country: "CH", City: "Lausanne", PostalCode: "1000"}}
	doc, err := RegisterUser(user)
	assert.Nil(t, err)

	userID := doc.ID
	user2, err := GetUserFromID(userID)
	assert.Nil(t, err)
	assert.Equal(t, user.Username, user2.Username)
	assert.Equal(t, user.Firstname, user2.Firstname)
	assert.Equal(t, user.Surname, user2.Surname)
	assert.Equal(t, user.Address, user2.Address)

	user2, err = GetUserFromID("invalid")
	assert.NotNil(t, err)
	assert.Equal(t, User{}, user2)
}
