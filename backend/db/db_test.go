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
	exit := m.Run()
	// delete collection after run
	client, err := GetDBConnection()

	if err != nil {
		log.Print(err.Error())
		os.Exit(1)
	}
	err = deleteCollection(context.Background(), client, client.Collection("users"), 64)
	if err != nil {
		log.Print(err.Error())
		os.Exit(1)
	}
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

func TestInvalidRegisterUser(t *testing.T) {
	_, err := RegisterUser("", "123", "Luc", "Skywalter", "Rue des Grands 28")
	if err == nil {
		t.Errorf("Empty username should not be allowed\n")
	}
	_, err = RegisterUser("luc", "", "Luc", "Skywalter", "Rue des Grands 28")
	if err == nil {
		t.Errorf("Empty password should not be allowed\n")
	}
	_, err = RegisterUser("luc", "123", "", "Skywalter", "Rue des Grands 28")
	if err == nil {
		t.Errorf("Empty firstname should not be allowed\n")
	}
	_, err = RegisterUser("luc", "123", "Luc", "", "Rue des Grands 28")
	if err == nil {
		t.Errorf("Empty surname should not be allowed\n")
	}
	_, err = RegisterUser("luc", "123", "Luc", "Skywalter", "")
	if err == nil {
		t.Errorf("Empty shipping address should not be allowed\n")
	}
}
