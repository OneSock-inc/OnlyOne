package db

import (
	utils "backend/utils"
	"context"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"cloud.google.com/go/firestore"
	"github.com/sjwhitworth/golearn/kdtree"
	"github.com/sjwhitworth/golearn/metrics/pairwise"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/api/iterator"
)

var projectID string = "onlyone-cb08e"
var dbClient *firestore.Client
var ctx *context.Context

const UsersCollection = "users"
const SocksCollection = "socks"

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
	ID       string `json:"id"`
	ShoeSize uint8  `firestore:"shoeSize" json:"shoeSize"`
	//is it a high or low profile sock
	Type         Profile  `firestore:"type" json:"type"`
	Color        string   `firestore:"color" json:"color"`
	Description  string   `firestore:"description" json:"description"`
	Picture      string   `firestore:"picture" json:"picture"`
	Owner        string   `firestore:"owner" json:"owner"`
	RefusedList  []string `firestore:"refusedList" json:"refusedList"`
	AcceptedList []string `firestore:"acceptedList" json:"acceptedList"`
	Match        string   `firestore:"match" json:"match"`
}

//return all the socks of a user identified by it's cookie session

func GetUserSocks(userID string) ([]Sock, error) {
	client, err := GetDBConnection()
	if err != nil {
		return nil, err
	}

	query := client.Collection(SocksCollection).Query.Where("owner", "==", userID)
	iter := query.Documents(context.Background())
	var socks []Sock
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("Failed to iterate: %v", err)
			return nil, err
		}
		var s Sock
		doc.DataTo(&s)
		s.ID = doc.Ref.ID
		socks = append(socks, s)
	}
	return socks, nil
}

func GetUser(username string) (*firestore.DocumentSnapshot, error) {
	db, err := GetDBConnection()
	if err != nil {
		return nil, err
	}
	query := db.Collection(UsersCollection).Where("username", "==", username)
	users, err := query.Documents(*ctx).GetAll()
	if err != nil {
		log.Printf("error : %v\n", err)
		return nil, err
	}
	if len(users) == 0 {
		return nil, fmt.Errorf("user `%s` not found", username)
	}
	if len(users) > 1 {
		return nil, fmt.Errorf("multiple users `%s` found", username)
	}
	return users[0], nil
}

func GetUserFromID(id string) (User, error) {
	db, err := GetDBConnection()
	if err != nil {
		return User{}, err
	}
	doc, err := db.Collection(UsersCollection).Doc(id).Get(context.Background())
	if err != nil {
		return User{}, err
	}

	var user User
	doc.DataTo(&user)
	return user, nil
}

func EditMatchingSock(sock Sock, otherSock Sock, accept bool) error {
	otherSock, err := GetSockInfo(otherSock.ID)
	if err != nil {
		return err
	}

	if sock.Owner == otherSock.Owner {
		return fmt.Errorf("User cannot accept or refuse a sock he owns")
	}
	if utils.Contains(sock.AcceptedList, otherSock.ID) {
		return fmt.Errorf("Sock already in the accepted list")
	}
	if utils.Contains(sock.RefusedList, otherSock.ID) {
		return fmt.Errorf("Sock already in the refused list")
	}

	for _, s := range []Sock{sock, otherSock} {
		if s.Match != "" {
			return fmt.Errorf("Sock `" + s.ID + "` is already in a happy and fulfilling pair")
		}
	}

	db, err := GetDBConnection()
	if err != nil {
		return err
	}

	if accept {
		sock.AcceptedList = append(sock.AcceptedList, otherSock.ID)

		//if the other sock already accepted us and we are accepting it now, then we got a match
		if utils.Contains(otherSock.AcceptedList, sock.ID) {
			sock.Match = otherSock.ID
			otherSock.Match = sock.ID
			_, err = db.Collection(SocksCollection).Doc(otherSock.ID).Set(context.Background(), otherSock)
			if err != nil {
				return err
			}
			// TODO: alert user there is a match
		}
	} else {
		sock.RefusedList = append(sock.RefusedList, otherSock.ID)
	}

	_, err = db.Collection(SocksCollection).Doc(sock.ID).Set(context.Background(), sock)
	if err != nil {
		return err
	}

	return nil
}

// get a sock struct from the database
func GetSockInfo(sockId string) (Sock, error) {
	client, err := GetDBConnection()
	if err != nil {
		return Sock{}, err
	}
	ref, err := client.Collection(SocksCollection).Doc(sockId).Get(context.Background())
	if err != nil {
		return Sock{}, err
	}
	if !ref.Exists() {
		return Sock{}, fmt.Errorf("the given sock id doesn't exist")
	}
	var s Sock
	if err := ref.DataTo(&s); err != nil {
		return Sock{}, fmt.Errorf("corrupted data, unable to read database")
	}

	//in order to have an empty json array and not a null when converting from the go struct to the json repr
	// related to TestGetSockInfo@db_test.go
	//{..
	//"refusedList": [],
	//"acceptedList": [],...}
	if s.AcceptedList == nil {
		s.AcceptedList = make([]string, 0)
	}
	if s.RefusedList == nil {
		s.RefusedList = make([]string, 0)
	}
	s.ID = ref.Ref.ID
	return s, nil
}

func NewSock(shoeSize uint8, type_ Profile, color string, desc string, Pictureb64 string, owner string) (Sock, error) {
	if shoeSize > 75 {
		return Sock{}, fmt.Errorf("show size `%d` is giant ! Are you a giant ? I don't think so", shoeSize)
	}
	if shoeSize <= 5 {
		return Sock{}, fmt.Errorf("show size `%d` is very small ! Are you a dwarf ? I don't think so", shoeSize)
	}
	if type_ >= count {
		return Sock{}, fmt.Errorf("type `%d` is invalid", count)
	}
	_, err := utils.ParseHexColor(color)
	if err != nil {
		return Sock{}, err
	}
	if strings.TrimSpace(desc) == "" {
		return Sock{}, fmt.Errorf("description is empty")
	}
	if strings.TrimSpace(Pictureb64) == "" {
		return Sock{}, fmt.Errorf("picture is empty")
	}
	// TODO: validate base64 + image data
	client, err := GetDBConnection()
	if err != nil {
		return Sock{}, err
	}
	userSnapShot, err := client.Collection(UsersCollection).Doc(owner).Get(context.Background())
	if err != nil {
		return Sock{}, fmt.Errorf("user doesn't exist %s", err.Error())
	}
	if !userSnapShot.Exists() {
		return Sock{}, fmt.Errorf("document doesn't exist")
	}

	s := Sock{
		ShoeSize:     shoeSize,
		Type:         type_,
		Color:        color,
		Description:  desc,
		Picture:      Pictureb64,
		Owner:        owner,
		RefusedList:  make([]string, 0),
		AcceptedList: make([]string, 0),
		Match:        "",
	}
	docRef, _, err := client.Collection(SocksCollection).Add(*ctx, s)
	s.ID = docRef.ID
	return s, err
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
	doc, err := GetUser(username)
	if err != nil {
		return "", fmt.Errorf("password not correct for user `%s`", username)
	}
	var user User
	doc.DataTo(&user)
	log.Printf("trying to log with %s/%s", username, password)
	//CompareHashAndPassword take the salt part from the hash and verify using it
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", fmt.Errorf("password not correct for user `%s`", username)
	}
	return doc.Ref.ID, nil
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
	query := client.Collection(UsersCollection).Query.Where("username", "==", u.Username)
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
	docRef, _, err := client.Collection(UsersCollection).Add(*ctx, user)

	if err != nil {
		log.Printf("error : %v\n", err)
		return nil, err
	}
	return docRef, nil
}

func DeleteCollection(ctx context.Context, client *firestore.Client,
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

func getFeaturesFromSock(s *Sock) []float64 {
	rgb, _ := utils.ParseHexColor(s.Color)
	return []float64{
		float64(s.ShoeSize) * 125 * 4,
		float64(s.Type) * 250 * 4,
		float64(rgb.A),
		float64(rgb.R),
		float64(rgb.G),
		float64(rgb.B),
	}
}

/*
GetCompatibleSocks returns the most similar sock in the collection
*/
func GetCompatibleSocks(sockId string, limit uint16) ([]Sock, error) {
	tree := kdtree.New()

	client, err := GetDBConnection()
	if err != nil {
		return nil, err
	}
	doc, err := client.Collection(SocksCollection).Doc(sockId).Get(context.Background())
	if err != nil {
		return nil, err
	}
	var originalSock Sock
	doc.DataTo(&originalSock)
	socks := make([]Sock, 0, limit)

	it := client.Collection(SocksCollection).DocumentRefs(context.Background())
	//matrix of sock's features each row is an array of the sock's feature
	datas := make([][]float64, 0)
	//Query.Where("shoeSize", "==", s.ShoeSize).Where("type", "==", s.Type).Where("isMatched", "==", false).Documents(context.Background())
	var i uint16 = 0
	for {
		//if we are done
		doc, err := it.Next()
		if err == iterator.Done {
			break
		}
		ref, err := doc.Get(context.Background())
		if err != nil {
			return nil, err
		}
		if !ref.Exists() || ref.Ref.ID == sockId {
			//sock doesn't exist or it's the sock we are looking at
			continue
		}
		dockSnapShot, err := ref.Ref.Get(context.Background())
		if err != nil {
			return nil, fmt.Errorf("while iterating compatible sock: %v", err)
		}

		var currentSock Sock
		dockSnapShot.DataTo(&currentSock)
		//we don't want to match two socks from the same owner or a sock already matched
		if currentSock.Owner == originalSock.Owner || currentSock.Match != "" {
			continue
		}
		log.Printf("%+v", currentSock)
		datas = append(datas, getFeaturesFromSock(&currentSock))

		currentSock.ID = dockSnapShot.Ref.ID
		socks = append(socks, currentSock)
		i++
	}

	if err = tree.Build(datas); err != nil {
		return nil, nil
	}
	// i := len(socks)
	if i > limit {
		i = limit
	}

	/*
		This is how to datas in arranged in datas
		datas = {
			sock1 : [feature array]
			sock2 : [feature array]
			sock3 : [feature array]
			sock4 : [feature array]
			sock5 : [feature array]
			sock6 : [feature array]
			sock7 : [feature array]
		}
	*/

	euclide := pairwise.NewEuclidean()
	rgb, _ := utils.ParseHexColor(originalSock.Color)
	//search limit sock similar (euclide) to s.attribut

	//rows contains the indexes of the most similar socks, fetching socks[rows[0]] gives the best matching sock
	//fetching datas[rows[0]] gives the feature of the best matching sock
	rows, _ /*pairwise distance from sockOP to currentSock_i*/ /*err*/, _ := tree.Search(int(i), euclide, []float64{
		float64(originalSock.ShoeSize) * 125 * 4,
		float64(originalSock.Type) * 250 * 4,
		float64(rgb.A),
		float64(rgb.R),
		float64(rgb.G),
		float64(rgb.B),
	})

	//take the limit socks
	res := make([]Sock, 0, limit)
	for _, s := range rows {
		//add the k best socks (limits to limit)
		res = append(res, socks[s])
	}

	return res, nil
}
