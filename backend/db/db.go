package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"cloud.google.com/go/firestore"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/api/iterator"
)

var projectID string = "onlyone-cb08e"
var dbClient *firestore.Client
var ctx *context.Context

type User struct {
	Username        string `firestore:"username" json:"username"`
	Firstname       string `firestore:"firstname" json:"firstname"`
	Surname         string `firestore:"surname" json:"surname"`
	Hash            []byte `firestore:"hash" json:"hash"`
	Salt            string `firestore:"salt" json:"salt"`
	ShippingAddress string `firestore:"shippingAddress" json:"shippingAddress"`
	SessionCookie   string `firestore:"sessionCookie" json:"sessionCookie"`
}

func createClient(ctx context.Context) (*firestore.Client, error) {

	err := os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "./data/service-account.json")
	if err != nil {
		return nil, err
	}
	// Sets your Google Cloud Platform project ID.
	return firestore.NewClient(ctx, projectID)
}

func GetClient() (*firestore.Client, error) {

	if dbClient == nil {
		c := context.Background()
		ctx = &c
		dbClient, err := createClient(*ctx)
		return dbClient, err
	}
	return dbClient, nil
}

// here we get the document having the username, we then retrieve the salt from the doc and the hashed password from the doc,
// we check if hash(password+salt) == doc.hash
func CheckUser(username string, password string) (User, error) {
	db, err := GetClient()
	if err != nil {
		log.Printf("error : %v\n", err)
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
	fmt.Printf("User: %v", user)

	err = bcrypt.CompareHashAndPassword(user.Hash, []byte(password+user.Salt))
	if err != nil {
		return User{}, fmt.Errorf("password not correct")
	}
	return user, nil
}

func CheckCookie(cookie string) (User, error) {
	client, err := GetClient()
	now := time.Now()
	if err != nil {
		log.Printf("error : %v\n", err)
		return User{}, err
	}

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
