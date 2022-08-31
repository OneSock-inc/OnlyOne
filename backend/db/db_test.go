package db

import (
	"context"
	"log"
	"os"
	"testing"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
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
