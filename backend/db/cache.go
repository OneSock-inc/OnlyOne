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
	sockId        []string
	socks         map[string]Sock
	socksFeatures [][]float64
}

func (c *Cache) getStrIdFromIdx(idx int) string {
	return c.sockId[idx]
}

func (c *Cache) add(s Sock, sockFeature []float64) {
	c.sockId = append(c.sockId, s.ID)
	c.socks[s.ID] = s
	c.socksFeatures = append(c.socksFeatures, sockFeature)
}
func (c *Cache) update(s Sock) {
	log.Println("Updating cache")
	//don't update sockId because the data dosen't move
	c.socks[s.ID] = s
	//features don't change
}

func newCache() (*Cache, error) {
	log.Println("creating cache")
	socks := make(map[string]Sock, 0)
	socksFeatures := make([][]float64, 0)
	socksId := make([]string, 0)
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
	i := 0
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
		socks[currentSock.ID] = currentSock
		socksId = append(socksId, currentSock.ID)
		i++
	}

	return &Cache{
			socks:         socks,
			sockId:        socksId,
			socksFeatures: socksFeatures,
		},
		nil
}
