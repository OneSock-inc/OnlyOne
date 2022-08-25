package main

import (
	"backend/api"
	"log"
)

// [START firestore_setup_client_create]

func main() {
	log.SetFlags(log.Flags() | log.Llongfile)
	engine := api.Setup()
	engine.Run(":8080")

	// Get a Firestore client.
	// client, err := db.GetClient()
	// if err != nil {
	// 	log.Fatalf("error : %v", err)
	// }
	// defer client.Close()
	// iter := client.Collection("users").Documents(context.Background())
	// for {
	// 	doc, err := iter.Next()
	// 	if err == iterator.Done {
	// 		break
	// 	}
	// 	if err != nil {
	// 		log.Fatalf("Failed to iterate: %v", err)
	// 	}
	// 	var user db.User
	// 	doc.DataTo(&user)
	// 	log.Printf("User: %s", user.Username)
	// }
}
