package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetMongoClient creates a mongo client
func GetMongoClient(host, port string) (*mongo.Client, error) {

	var c *mongo.Client
	var e error

	//Perform connection creation operation only once.
	mongoOnce.Do(func() {
		// Set client options
		clientOptions := options.Client().ApplyURI(host + ":" + port)

		// Connect to MongoDB
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			e = err
		}

		// Check the connection
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			e = err
		}
		c = client
	})
	return c, e
}