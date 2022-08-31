package db

import (
	utils "backend/utils"
	"context"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"time"

	"cloud.google.com/go/firestore"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/api/iterator"
)

var projectID string = "onlyone-cb08e"
var dbClient *firestore.Client
var ctx *context.Context

type Address struct {
	Street     string `firestore:"street" json:"street"`
	Country    string `firestore:"country" json:"country"`
	City       string `firestore:"city" json:"city"`
	PostalCode string `firestore:"postalCode" json:"postalCode"`
}

type User struct {
	Username  string  `firestore:"username" json:"username"`
	Firstname string  `firestore:"firstname" json:"firstname"`
	Surname   string  `firestore:"surname" json:"surname"`
	Password  string  `firestore:"hash" json:"password"`
	Address   Address `firestore:"address" json:"address"`
}

// Ther is two type : low and high
type Profile uint8

const (
	low Profile = iota
	high
	count
)

type Sock struct {
	ShoeSize uint8 `firestore:"shoeSize" json:"shoeSize"`
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

func NewSock(shoeSize uint8, type_ Profile, color string, desc string, Pictureb64 string, owner string) (*firestore.DocumentRef, error) {
	if shoeSize > 75 {
		return nil, fmt.Errorf("show size `%d` is giant ! Are you a giant ? I don't think so", shoeSize)
	}
	if shoeSize <= 5 {
		return nil, fmt.Errorf("show size `%d` is very small ! Are you a dwarf ? I don't think so", shoeSize)
	}
	if type_ >= count {
		return nil, fmt.Errorf("type `%d` is invalid", count)
	}
	_, err := utils.ParseHexColor(color)
	if err != nil {
		return nil, err
	}
	if strings.TrimSpace(desc) == "" {
		return nil, fmt.Errorf("description is empty")
	}
	if strings.TrimSpace(Pictureb64) == "" {
		return nil, fmt.Errorf("picture is empty")
	}
	// TODO: validate base64 + image data
	client, err := GetDBConnection()
	if err != nil {
		return nil, err
	}
	userSnapShot, err := client.Collection("users").Doc(owner).Get(context.Background())
	if !userSnapShot.Exists() {
		errMsg := ""
		if err != nil {
			errMsg = err.Error()
		}
		return nil, fmt.Errorf("user doesn't exist %s", errMsg)
	}

	s := Sock{
		ShoeSize:     shoeSize,
		Type:         type_,
		Color:        color,
		Description:  desc,
		Picture:      Pictureb64,
		Owner:        owner,
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
func VerifyLogin(username string, password string) (string, error) {
	db, err := GetDBConnection()
	if err != nil {
		return "", err
	}
	query := db.Collection("users").Where("username", "==", username)
	users, err := query.Documents(*ctx).GetAll()
	if err != nil {
		log.Printf("error : %v\n", err)
		return "", err
	}
	if len(users) == 0 {
		return "", fmt.Errorf("user `%s` not found", username)
	}
	if len(users) > 1 {
		return "", fmt.Errorf("multiple users `%s` found", username)
	}
	var user User
	users[0].DataTo(&user)
	log.Printf("trying to log with %s/%s", username, password)
	//CompareHashAndPassword take the salt part from the hash and verify using it
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", fmt.Errorf("password not correct for user `%s`", username)
	}
	return users[0].Ref.ID, nil
}

func CheckCookie(cookie string) (*firestore.DocumentRef, error) {
	client, err := GetDBConnection()
	if err != nil {
		return nil, err
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
			return nil, err
		}
		//if cookie is fresh (less than 1 day)
		if now.Sub(doc.UpdateTime).Hours() < 24 {
			var user User
			doc.DataTo(&user)
			return doc.Ref, nil
		}
	}
	return nil, nil
}

// register a user in the db, currently the hash field should be a clear password
// unfortunatly we cannot have an option as in rust
func RegisterUser(u User) (*firestore.DocumentRef, error) {
	if strings.TrimSpace(u.Username) == "" {
		return nil, fmt.Errorf("username is empty")
	}
	if strings.TrimSpace(u.Password) == "" {
		return nil, fmt.Errorf("password is empty")
	}
	if strings.TrimSpace(u.Firstname) == "" {
		return nil, fmt.Errorf("firstname is empty")
	}
	if strings.TrimSpace(u.Surname) == "" {
		return nil, fmt.Errorf("surname is empty")
	}
	if strings.TrimSpace(u.Address.City) == "" {
		return nil, fmt.Errorf("city address is empty")
	}
	if strings.TrimSpace(u.Address.Country) == "" {
		return nil, fmt.Errorf("country address is empty")
	}
	if strings.TrimSpace(u.Address.PostalCode) == "" {
		return nil, fmt.Errorf("postal address is empty")
	}
	if strings.TrimSpace(u.Address.Street) == "" {
		return nil, fmt.Errorf("street address is empty")
	}

	client, err := GetDBConnection()
	if err != nil {
		return nil, err
	}
	//query doc where username's field == `username`
	query := client.Collection("users").Query.Where("username", "==", u.Username)
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
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	log.Printf("Hashed password : %s\n", hash)
	user := u
	user.Password = string(hash)
	docRef, _, err := client.Collection("users").Add(*ctx, user)

	if err != nil {
		log.Printf("error : %v\n", err)
		return nil, err
	}
	return docRef, nil
}

func SetCookie(cookie string, username string) error {
	client, err := GetDBConnection()
	if err != nil {
		return err
	}

	query := client.Collection("users").Query.Where("username", "==", username)
	docs, err := query.Documents(context.Background()).GetAll()
	if err != nil {
		log.Printf("error : %v\n", err)
		return err
	}
	//if there are docs with this username
	if len(docs) != 1 {
		return fmt.Errorf("user already exists")
	}

	doc := docs[0]
	data := doc.Data()
	data["sessionCookie"] = cookie
	client.Collection("users").Doc(doc.Ref.ID).Set(context.Background(), data)

	return nil
}
