package db

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/api/iterator"
)

/*
This is how to datas in arranged in the socksFeature field of Cache

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
type Cache struct {
	socks         []Sock
	socksFeatures [][]float64
}

func (c *Cache) update(s Sock, sockFeatures []float64) {
	log.Println("Updating cache")
	c.socks = append(c.socks, s)
	c.socksFeatures = append(c.socksFeatures, sockFeatures)
}

func newCache() (*Cache, error) {
	log.Println("creating cache")
	socks := make([]Sock, 0)
	socksFeatures := make([][]float64, 0)

	dbClient, err := GetDBConnection()
	if err != nil {
		return nil, err
	}
	// err := DeleteCollection(context.Background(), dbClient, dbClient.Collection(matrixCollection), 64)
	// if err != nil {
	// 	return err
	// }
	it := dbClient.Collection(SocksCollection).DocumentRefs(context.Background())
	//matrix of sock's features each row is an array of the sock's feature

	//Query.Where("shoeSize", "==", s.ShoeSize).Where("type", "==", s.Type).Where("isMatched", "==", false).Documents(context.Background())

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
		if !ref.Exists() {
			//sock doesn't exist
			continue
		}
		dockSnapShot, err := ref.Ref.Get(context.Background())
		if err != nil {
			return nil, fmt.Errorf("while iterating compatible sock: %v", err)
		}

		var currentSock Sock
		dockSnapShot.DataTo(&currentSock)
		//we don't want an already matched sock
		if currentSock.Match != "" {
			continue
		}
		log.Printf("cached: %s\n", currentSock.ID)
		socksFeatures = append(socksFeatures, GetFeaturesFromSock(&currentSock))

		currentSock.ID = dockSnapShot.Ref.ID
		socks = append(socks, currentSock)
	}

	return &Cache{
			socks: socks,
			// owners:        owners,
			socksFeatures: socksFeatures,
		},
		nil
}
