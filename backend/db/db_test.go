package db

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Setenv(FirestoreEmulatorHost, "localhost:8080")
	log.SetFlags(log.Flags() | log.Llongfile)

	client, err := GetDBConnection()

	if err != nil {
		log.Print(err.Error())
		os.Exit(1)
	}
	err = DeleteCollection(*ContextBd, client, client.Collection("users"), 64)
	if err != nil {
		log.Print(err.Error())
		os.Exit(1)
	}
	err = DeleteCollection(*ContextBd, client, client.Collection("socks"), 64)
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
		t.Errorf(" user should not exist %s\n", err.Error())
	}
	_, err = RegisterUser(usr)

	if err == nil {
		t.Errorf("the user should exist : %s\n", err.Error())
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

func TestInvalidGetUserSocks(t *testing.T) {
	socks, err := GetUserSocks("invalid")
	assert.Nil(t, err)
	assert.Zero(t, len(socks))
}

func TestGetUserSocks(t *testing.T) {
	user := User{Username: "james2010", Password: "123", Firstname: "James", Surname: "Wow", Address: Address{Street: "Non", Country: "CH", City: "GE", PostalCode: "1212"}}
	doc, err := RegisterUser(user)
	assert.Nil(t, err)

	userID := doc.ID
	socks, err := GetUserSocks(userID)
	assert.Nil(t, err)
	assert.Zero(t, len(socks))

	sock, err := NewSock(10, low, "#000", "Super sock !", "aGkK", userID)
	assert.Nil(t, err)

	sockID := sock.ID
	socks, err = GetUserSocks(userID)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(socks))
	assert.Equal(t, sockID, socks[0].ID)
}

func TestGetInfoSock(t *testing.T) {
	//set a user
	doc, err := RegisterUser(User{Username: "jackob", Password: "123", Firstname: "James", Surname: "Wow", Address: Address{Street: "Non", Country: "CH", City: "GE", PostalCode: "1212"}})
	assert.Nil(t, err)
	owner := doc.ID

	s := Sock{
		ShoeSize:     41,
		Type:         Profile(1),
		Color:        "#BEDEAD",
		Description:  "I tried selling it on onlyFan, but it didn't work",
		Picture:      "JHAKHSD==",
		RefusedList:  make([]string, 0), //this init the memory see GetSockInfo@db.go for further detail
		AcceptedList: make([]string, 0),
		Owner:        owner,
	}

	s2, err := NewSock(s.ShoeSize, s.Type, s.Color, s.Description, s.Picture, s.Owner)
	s.ID = s2.ID
	assert.Nil(t, err)
	sockId := s2.ID
	sockBack, err := GetSockInfo(sockId)
	assert.Nil(t, err)
	assert.Equal(t, s, sockBack)

}

func TestGetInfoSockNilLists(t *testing.T) {
	//set a user
	doc, err := RegisterUser(User{Username: "Henry", Password: "123", Firstname: "James", Surname: "Wow", Address: Address{Street: "Non", Country: "CH", City: "GE", PostalCode: "1212"}})
	assert.Nil(t, err)
	owner := doc.ID

	//don't init the slice List therefore we get a null when marshalling struct -> firestore -> struct
	s := Sock{
		ShoeSize:    41,
		Type:        Profile(1),
		Color:       "#BEDEAD",
		Description: "I tried selling it on onlyFan, but it didn't work",
		Picture:     "JHAKHSD==",
		Owner:       owner,
	}

	s2, err := NewSock(s.ShoeSize, s.Type, s.Color, s.Description, s.Picture, s.Owner)
	assert.Nil(t, err)
	sockId := s2.ID
	sockBack, err := GetSockInfo(sockId)
	assert.Nil(t, err)
	assert.NotNil(t, sockBack.AcceptedList)
	assert.NotNil(t, sockBack.AcceptedList)
	s.ID = s2.ID
}

func TestGetUser(t *testing.T) {
	user := User{Username: "jamy", Password: "123", Firstname: "Jamy", Surname: "Yuy", Address: Address{Street: "Arf", Country: "CH", City: "Lausanne", PostalCode: "1000"}}
	doc, err := RegisterUser(user)
	assert.Nil(t, err)

	userID := doc.ID
	docr, err := GetUser("jamy")
	assert.Nil(t, err)
	assert.Equal(t, userID, docr.Ref.ID)

	_, err = GetUser("invalid")
	assert.NotNil(t, err)
}

func TestGetUserFromID(t *testing.T) {
	user := User{Username: "ronron", Password: "123", Firstname: "Ron", Surname: "Ron", Address: Address{Street: "Rond", Country: "CH", City: "Lausanne", PostalCode: "1000"}}
	doc, err := RegisterUser(user)
	assert.Nil(t, err)

	userID := doc.ID
	user2, err := GetUserFromID(userID)
	assert.Nil(t, err)
	assert.Equal(t, user.Username, user2.Username)
	assert.Equal(t, user.Firstname, user2.Firstname)
	assert.Equal(t, user.Surname, user2.Surname)
	assert.Equal(t, user.Address, user2.Address)

	user2, err = GetUserFromID("invalid")
	assert.NotNil(t, err)
	assert.Equal(t, User{}, user2)
}

func TestGetCompatibleSocks(t *testing.T) {
	//delete all the socks
	//client, err := GetDBConnection()
	//assert.Nil(t, err)
	//assert.Nil(t, DeleteCollection(*ContextBd, client, client.Collection(SocksCollection), 64))
	//assert.Nil(t, err)

	//create a user
	doc, err := RegisterUser(User{Username: "Herbe", Password: "123", Firstname: "James", Surname: "Wow", Address: Address{Street: "Non", Country: "CH", City: "GE", PostalCode: "1212"}})
	assert.Nil(t, err)
	//create another User
	doc2, err := RegisterUser(User{Username: "Banane", Password: "123", Firstname: "James", Surname: "Wow", Address: Address{Street: "Non", Country: "CH", City: "GE", PostalCode: "1212"}})

	assert.Nil(t, err)
	owner := doc.ID

	//create two similar socks with their owner beeing the new user
	s := Sock{
		ShoeSize:     41,
		Type:         Profile(1),
		Color:        "#BEDEAD",
		Description:  "I tried selling it on onlyFan, but it didn't work",
		Picture:      "==JHAKHSD",
		RefusedList:  make([]string, 0), //this init the memory see GetSockInfo@db.go for further detail
		AcceptedList: make([]string, 0),
		Owner:        owner,
	}
	sd, err := NewSock(s.ShoeSize, s.Type, s.Color, s.Description, s.Picture, s.Owner)
	s.ID = sd.ID
	assert.Nil(t, err)

	s1 := Sock{
		ShoeSize:     41,
		Type:         Profile(1),
		Color:        "#FFF",
		Description:  "I tried selling it on onlyFan, but i'm now disgusted by me",
		Picture:      "==",
		RefusedList:  make([]string, 0), //this init the memory see GetSockInfo@db.go for further detail
		AcceptedList: make([]string, 0),
		Owner:        doc2.ID,
	}

	s1d, err := NewSock(s1.ShoeSize, s1.Type, s.Color, s1.Description, s1.Picture, s1.Owner)
	s1.ID = s1d.ID
	assert.Nil(t, err)
	socks, err := GetCompatibleSocks(s1.ID)
	assert.Nil(t, err)
	for _, sock := range socks {
		log.Printf("%v\n", sock)
	}
	//assert.True(t, len(socks) == 1)
	//assert.True(t, socks[0].ID == s.ID)
}

func TestGetCompatibleSocksWithManySocksAndUser(t *testing.T) {
	//client, err := GetDBConnection()
	//assert.Nil(t, err)
	//assert.Nil(t, DeleteCollection(*ContextBd, client, client.Collection(SocksCollection), 64))
	sockId := ""
	//create 10 users with two socks each
	for i := 0; i < 10; i++ {
		user := User{Username: "kokoroko" + fmt.Sprint(i), Password: "123", Firstname: "James", Surname: "Wow", Address: Address{Street: "Non", Country: "CH", City: "GE", PostalCode: "1212"}}
		doc, err := RegisterUser(user)
		assert.Nil(t, err)
		owner := doc.ID
		s := Sock{
			ShoeSize:     uint8(41 - i),
			Type:         Profile(1 - i%2),
			Color:        "#BEDEAD",
			Description:  "I tried selling it on onlyFan, but it didn't work",
			Picture:      "==JHAKHSD",
			RefusedList:  make([]string, 0), //this init the memory see GetSockInfo@db.go for further detail
			AcceptedList: make([]string, 0),
			Owner:        owner,
		}
		sd, err := NewSock(s.ShoeSize, s.Type, s.Color, s.Description, s.Picture, s.Owner)
		s.ID = sd.ID
		assert.Nil(t, err)

		s1 := Sock{
			ShoeSize:     uint8(41 + i),
			Type:         Profile(0 + i%2),
			Color:        "#FFF",
			Description:  fmt.Sprintf("i'm owned by %s", user.Username),
			Picture:      "==",
			RefusedList:  make([]string, 0), //this init the memory see GetSockInfo@db.go for further detail
			AcceptedList: make([]string, 0),
			Owner:        owner,
		}

		s1d, err := NewSock(s1.ShoeSize, s1.Type, s.Color, s1.Description, s1.Picture, s1.Owner)
		assert.Nil(t, err)

		s1.ID = s1d.ID
		//remember the last sock
		sockId = s1d.ID
	}

	//create two similar socks with their owner beeing the new user
	socks, err := GetCompatibleSocks(sockId)
	assert.Nil(t, err)
	//we created two sock by user 10 times we should get 4 of them (defined as the maximum for a sock)
	// assert.True(t, len(socks) == int(MAX))
	assert.True(t, len(socks) != 0)
	// assert.True(t, math.Abs(float64(socks[0].ShoeSize)-float64(socks[1].ShoeSize)) <= 2)
	// for i := 1; uint16(i) < MAX; i++ {
	// 	//assert than the shoesSize are similar when looking at two similar shoes
	// 	// usually for a sock size 42 type 0 we will get [40,0]
	// 	assert.True(t, math.Abs(float64(socks[i-1].ShoeSize)-float64(socks[i].ShoeSize)) <= 4)
	// }
}

func TestGetCompatibleSocksNotInListAfterMatch(t *testing.T) {
	user1 := User{Username: "Estelle", Password: "123", Firstname: "James", Surname: "Wow", Address: Address{Street: "Non", Country: "CH", City: "GE", PostalCode: "1212"}}
	user2 := User{Username: "Elisabeth", Password: "123", Firstname: "James", Surname: "Wow", Address: Address{Street: "Non", Country: "CH", City: "GE", PostalCode: "1212"}}

	doc, err := RegisterUser(user1)
	assert.Nil(t, err)
	id1 := doc.ID
	doc, err = RegisterUser(user2)
	assert.Nil(t, err)
	id2 := doc.ID

	sock1, err := NewSock(50, 2, "#BEDEAD", "Niarf", "QnVncyBkZSBtZXJkZQo=", id1)
	assert.Nil(t, err)
	sock2, err := NewSock(50, 2, "#BEDEAD", "Niarf", "QnVncyBkZSBtZXJkZQo=", id2)
	assert.Nil(t, err)

	log.Printf("sock ID1 : %+v", sock1.ID)
	log.Printf("sock ID2 : %+v", sock2.ID)

	socks, err := GetCompatibleSocks(sock1.ID)
	log.Printf("Socks : %+v\n\n", socks)
	result := false
	for _, sock := range socks {
		if sock.ID == sock2.ID {
			result = true
		}
	}
	assert.True(t, result)

	socks, err = GetCompatibleSocks(sock2.ID)
	log.Printf("Socks : %+v\n\n", socks)
	result = false
	for _, sock := range socks {
		if sock.ID == sock1.ID {
			result = true
		}
	}
	assert.True(t, result)

	// match !
	err = EditMatchingSock(sock1, sock2, true)
	assert.Nil(t, err)
	err = EditMatchingSock(sock2, sock1, true)
	assert.Nil(t, err)

	// Should not be in the compatible list
	socks, err = GetCompatibleSocks(sock1.ID)
	//log.Printf("Socks : %+v", socks)
	result = false
	for _, sock := range socks {
		if sock.ID == sock2.ID {
			result = true
		}
	}
	assert.False(t, result)

	socks, err = GetCompatibleSocks(sock2.ID)
	//log.Printf("Socks : %+v", socks)
	result = false
	for _, sock := range socks {
		if sock.ID == sock1.ID {
			result = true
		}
	}
	assert.False(t, result)
}

func TestSocksProfiles(t *testing.T) {
	user := User{Username: "kikiriki", Password: "123", Firstname: "James", Surname: "Wow", Address: Address{Street: "Non", Country: "CH", City: "GE", PostalCode: "1212"}}
	doc, err := RegisterUser(user)
	assert.Nil(t, err)
	for i := 0; i < int(Count); i++ {
		_, err := NewSock(42, Profile(i), "#ABC", "magificent !", "Qydlc3QgcGFzIHRyw6hzIHN5bXBhcyDDp2EgOigK", doc.ID)
		assert.Nil(t, err)
	}
	_, err = NewSock(42, Profile(Count), "#ABC", "magificent !", "Qydlc3QgcGFzIHRyw6hzIHN5bXBhcyDDp2EgOigK", doc.ID)
	assert.NotNil(t, err)
}
