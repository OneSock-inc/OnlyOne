package db

import (
	"log"
	"os"
	"testing"
)

func Setup() {
	err := os.Chdir("../")
	if err != nil {
		log.Fatalf("error : %v", err)
	}

}

func TestGetClient(t *testing.T) {
	Setup()

	_, err := GetClient()
	if err != nil {
		t.Errorf("error : %v", err)
	}
}
func TestCheckUser(t *testing.T) {
	Setup()
	//this user should exist in the database
	_, err := CheckUser("test", "myPassword")
	if err != nil {
		t.Errorf("error : %v", err)
	}
}

func TestCreateUserAlreadyExist(t *testing.T) {
	Setup()
	user := User{
		Username:        "test",
		Firstname:       "Joris",
		Surname:         "Schaller",
		Hash:            []byte("myPassword"),
		Salt:            "salt",
		ShippingAddress: "this is a long shipping addr 1212 grand-Lancy",
		SessionCookie:   "",
	}

	err := CreateUser(user)
	if err == nil {
		t.Error("error user should exist")
	}
	if err != nil {
		log.Printf("the user should exist : %s", err.Error())
	}

}
