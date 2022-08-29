package db

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Setenv(FirestoreEmulatorHost, "localhost:8080")
	log.SetFlags(log.Flags() | log.Llongfile)
	os.Exit(m.Run())
}

const FirestoreEmulatorHost = "FIRESTORE_EMULATOR_HOST"

func TestGetClient(t *testing.T) {

	_, err := GetDBConnection()
	if err != nil {
		t.Errorf("erro : %v", err)
	}
}
func TestCheckUser(t *testing.T) {
	//this user should exist in the database
	const (
		username     = "username"
		pwd          = "myPwd"
		firstname    = "firstname"
		surname      = "surname"
		shippingAddr = "shippingAddr"
	)
	_, err := RegisterUser(username, pwd, firstname, surname, shippingAddr)
	if err != nil {
		t.Errorf("unable to create a user\n%v", err)
	}
	_, err = VerifyLogin("test", "myPassword")
	if err == nil {
		t.Errorf("error user don't exist in mockup\n%v", err)
	}

}

func TestCreateUser(t *testing.T) {
	const (
		username        = "test"
		firstname       = "Joris"
		surname         = "Schaller"
		pwd             = "myPassword"
		shippingAddress = "this is a long shipping addr 1212 grand-Lancy"
	)
	_, err := RegisterUser(username, pwd, firstname, surname, shippingAddress)
	if err != nil {
		t.Errorf(" user should not exist\n%v", err)
	}
	_, err = RegisterUser(username, pwd, firstname, surname, shippingAddress)

	if err == nil {
		log.Printf("the user should exist : %v", err)
	}
}
