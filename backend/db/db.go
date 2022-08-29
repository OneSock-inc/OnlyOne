package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"cloud.google.com/go/firestore"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/api/iterator"
)

var projectID string = "onlyone-cb08e"
var dbClient *firestore.Client
var ctx *context.Context

type User struct {
	ID              string `firestore:"ID,omitempty" json:"ID"`
	Username        string `firestore:"username" json:"username"`
	Firstname       string `firestore:"firstname" json:"firstname"`
	Surname         string `firestore:"surname" json:"surname"`
	Hash            []byte `firestore:"hash" json:"hash"`
	ShippingAddress string `firestore:"shippingAddress" json:"shippingAddress"`
	SessionCookie   string `firestore:"sessionCookie" json:"sessionCookie"`
}

// Ther is two type : low and high
type Profile int

const (
	low Profile = iota
	high
)

type Sock struct {
	SockId   string `firestore:"sockId,omitempty" json:"sockId"`
	ShoeSize int    `firestore:"shoeSize" json:"shoeSize"`
	//is it a high or low profile sock
	Type         Profile  `firestore:"type" json:"type"`
	Color        string   `firestore:"color" json:"color"`
	Description  string   `firestore:"description" json:"description"`
	Picture      string   `firestore:"picture" json:"picture"`
	Owner        string   `firestore:"owner" json:"owner"`
	RefusedList  []string `firestore:"refusedList" json:"refusedList"`
	AcceptedList []string `firestore:"acceptedList" json:"acceptedList"`
	IsMatched    bool     `firestore:"isMatched" json:"isMatched"`
}

//return all the socks of a user identified by it's cookie session

func getUserSocks(userCookie string) ([]Sock, error) {
	return make([]Sock, 0), nil
}

func getUser(username string) (User, error) {
	return User{}, nil
}
func editMatchingSock(sockID string, otherSockID string, accept bool) error {
	return nil
}

func getSockInfo(sockID string) Sock {
	return Sock{}
}

func NewSock(shoeSize int, size Profile, color string, desc string, Pictureb64 string) (*firestore.DocumentRef, error) {

	client, err := GetDBConnection()
	if err != nil {
		return nil, err
	}
	s := Sock{
		ShoeSize:     shoeSize,
		Type:         size,
		Color:        color,
		Description:  desc,
		Picture:      Pictureb64,
		RefusedList:  []string{},
		AcceptedList: []string{},
		IsMatched:    false,
	}
	docRef, _, err := client.Collection("socks").Add(*ctx, s)
	return docRef, err
}

func createClient(ctx context.Context) (*firestore.Client, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("unable to get the home dir %v", err)
		return nil, err
	}
	resPath := path.Join(home, "service-account.json")
	err = os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", resPath)
	if err != nil {
		return nil, err
	}
	// Sets your Google Cloud Platform project ID.
	return firestore.NewClient(ctx, projectID)
}
func GetDBConnection() (*firestore.Client, error) {

	if dbClient == nil {
		c := context.Background()
		ctx = &c
		return createClient(*ctx)
	}
	return dbClient, nil
}

// here we get the document having the username, we then retrieve the salt from the doc and the hashed password from the doc,
// we check if hash(password+salt) == doc.hash
func VerifyLogin(username string, password string) (User, error) {
	db, err := GetDBConnection()
	if err != nil {
		return User{}, err
	}
	query := db.Collection("users").Where("username", "==", username)
	users, err := query.Documents(*ctx).GetAll()
	if err != nil {
		log.Printf("error : %v\n", err)
		return User{}, err
	}
	if len(users) == 0 {
		return User{}, fmt.Errorf("user not found")
	}
	var user User
	users[0].DataTo(&user)
	//CompareHashAndPassword take the salt part from the hash and verify using it
	err = bcrypt.CompareHashAndPassword(user.Hash, []byte(password))
	if err != nil {
		return User{}, fmt.Errorf("password not correct")
	}
	return user, nil
}

func CheckCookie(cookie string) (User, error) {
	client, err := GetDBConnection()
	if err != nil {
		return User{}, err
	}
	now := time.Now()

	query := client.Collection("users").Query.Where("sessionCookie", "==", cookie)
	iter := query.Documents(context.Background())
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("Failed to iterate: %v", err)
			return User{}, err
		}
		//if cookie is fresh (less than 1 day)
		if now.Sub(doc.UpdateTime).Hours() < 24 {
			var user User
			doc.DataTo(&user)
			return user, nil
		}
	}
	return User{}, nil
}

func RegisterUser(username string, pwd string, firstname string, surname string, shippingAddr string) (*firestore.DocumentRef, error) {
	client, err := GetDBConnection()
	if err != nil {
		return nil, err
	}
	//query doc where username's field == `username`
	query := client.Collection("users").Query.Where("username", "==", username)
	docs, err := query.Documents(context.Background()).GetAll()
	if err != nil {
		log.Printf("error : %v\n", err)
		return nil, err
	}
	//if there are docs with this username
	if len(docs) != 0 {
		return nil, fmt.Errorf("user already exists")
	}

	//bcrypt's GenerateFromPassword generate the password with a salt !!
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := User{Username: username, Firstname: firstname, Surname: surname, Hash: hash, ShippingAddress: shippingAddr, SessionCookie: ""}
	docRef, _, err := client.Collection("users").Add(*ctx, user)

	if err != nil {
		log.Printf("error : %v\n", err)
		return nil, err
	}
	return docRef, nil
}
