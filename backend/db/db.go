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
	Username        string `firestore:"username" json:"username"`
	Firstname       string `firestore:"firstname" json:"firstname"`
	Surname         string `firestore:"surname" json:"surname"`
	Hash            []byte `firestore:"hash" json:"hash"`
	Salt            string `firestore:"salt" json:"salt"`
	ShippingAddress string `firestore:"shippingAddress" json:"shippingAddress"`
	SessionCookie   string `firestore:"sessionCookie" json:"sessionCookie"`
}

// Ther is two type : low and high
type profile int

const (
	low profile = iota
	high
)

type Sock struct {
	SockId   string `firestore:"sockId" json:"sockId"`
	ShoeSize string `firestore:"shoeSize" json:"shoeSize"`
	//is it a high or low profile sock
	Type           profile  `firestore:"type" json:"type"`
	Color          string   `firestore:"color" json:"color"`
	Description    string   `firestore:"description" json:"description"`
	Picture        string   `firestore:"picture" json:"picture"`
	Owner          string   `firestore:"owner" json:"owner"`
	CompatibleList []string `firestore:"compatibleList" json:"compatibleList"`
	AcceptList     []string `firestore:"acceptList" json:"acceptList"`
	IsMatched      bool     `firestore:"isMatched" json:"isMatched"`
}

func createClient(ctx context.Context) (*firestore.Client, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("unable to get the home dir %v", err)
		return nil, err
	}
	resPath := path.Join(home, ".ssh", "service-account.json")
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
		dbClient, err := createClient(*ctx)
		return dbClient, err
	}
	return dbClient, nil
}

// here we get the document having the username, we then retrieve the salt from the doc and the hashed password from the doc,
// we check if hash(password+salt) == doc.hash
func CheckUser(username string, password string) (User, error) {
	db, err := GetDBConnection()
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

	err = bcrypt.CompareHashAndPassword(user.Hash, []byte(password+user.Salt))
	if err != nil {
		return User{}, fmt.Errorf("password not correct")
	}
	return user, nil
}

func CheckCookie(cookie string) (User, error) {
	client, err := GetDBConnection()
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

func CreateUser(user User) error {
	client, err := GetDBConnection()
	if err != nil {
		log.Printf("error : %v\n", err)
		return err
	}
	query := client.Collection("users").Query.Where("username", "==", user.Username)
	docs, err := query.Documents(context.Background()).GetAll()
	if err != nil {
		log.Printf("error : %v\n", err)
		return err
	}
	if len(docs) != 0 {
		return fmt.Errorf("user already exists")
	}
	client.Collection("users").NewDoc().Set(*ctx, user)
	if err != nil {
		log.Printf("error : %v\n", err)
		return err
	}
	return nil
}
