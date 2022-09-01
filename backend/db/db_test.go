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

func TestGetInfoSock(t *testing.T) {
	//set a user
	doc, err := RegisterUser(User{Username: "jackob", Password: "123", Firstname: "James", Surname: "Wow", Address: Address{Street: "Non", Country: "CH", City: "GE", PostalCode: "1212"}})
	assert.Nil(t, err)
	owner := doc.ID

	s := Sock{
		ShoeSize:     41,
		Type:         Profile(1),
		Color:        "#BEDEAD",
		Description:  "I tried selling it on onlyFan, but it didn't work",
		Picture:      "JHAKHSD==",
		RefusedList:  make([]string, 0), //this init the memory see GetSockInfo@db.go for further detail
		AcceptedList: make([]string, 0),
		Owner:        owner,
	}

	s2, err := NewSock(s.ShoeSize, s.Type, s.Color, s.Description, s.Picture, s.Owner)
	assert.Nil(t, err)
	sockId := s2.ID
	sockBack, err := GetSockInfo(sockId)
	assert.Nil(t, err)
	assert.Equal(t, sockBack, s)

}

func TestGetInfoSockNilLists(t *testing.T) {
	//set a user
	doc, err := RegisterUser(User{Username: "Henry", Password: "123", Firstname: "James", Surname: "Wow", Address: Address{Street: "Non", Country: "CH", City: "GE", PostalCode: "1212"}})
	assert.Nil(t, err)
	owner := doc.ID

	//don't init the slice List therefore we get a null when marshalling struct -> firestore -> struct
	s := Sock{
		ShoeSize:    41,
		Type:        Profile(1),
		Color:       "#BEDEAD",
		Description: "I tried selling it on onlyFan, but it didn't work",
		Picture:     "JHAKHSD==",
		Owner:       owner,
	}

	s2, err := NewSock(s.ShoeSize, s.Type, s.Color, s.Description, s.Picture, s.Owner)
	assert.Nil(t, err)
	sockId := s2.ID
	sockBack, err := GetSockInfo(sockId)
	assert.Nil(t, err)
	assert.NotNil(t, sockBack.AcceptedList)
	assert.NotNil(t, sockBack.AcceptedList)

}
