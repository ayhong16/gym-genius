package src

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"server/initializers"
	"server/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	apiKey           string
	connectionString string
	Client           *mongo.Client
}

func NewDatabase() *Database {
	apiKey, connectionString := initializers.LoadEnvVariables()
	client := getClient(connectionString)

	return &Database{
		apiKey:           apiKey,
		connectionString: connectionString,
		Client:           client,
	}
}

func (db *Database) UpdateExercises() {
	fmt.Println("Updating MongoDB exercises...")

	collection := db.Client.Database("gym_management").Collection("exercises")

	exercises, err := db.fetchExercises()
	if err != nil {
		log.Fatalf("Failed to fetch exercises: %v", err)
		return
	}
	updateCollection(exercises, collection)
}

func (db *Database) Disconnect() {
	fmt.Println("Disconnecting from MongoDB...")
	if err := db.Client.Disconnect(context.TODO()); err != nil {
		panic(err)
	} else {
		fmt.Println("Disconnected from MongoDB")
	}
}

// func (db *Database) Disconnect() {
// 	fmt.Println("Disconnecting from MongoDB...")
// 	ctx := context.TODO()
// 	done := make(chan error, 1)
// 	go func() {
// 		done <- db.Client.Disconnect(ctx)
// 	}()

// 	select {
// 	case err := <-done:
// 		if err != nil {
// 			log.Printf("Failed to disconnect from MongoDB: %v", err)
// 		}
// 		log.Printf("Disconnected from MongoDB")
// 	case <-ctx.Done():
// 		if ctx.Err() == context.Canceled {
// 			log.Println("Context canceled while disconnecting from MongoDB")
// 		} else {
// 			log.Println("Context deadline exceeded while disconnecting from MongoDB")
// 		}
// 	}
// }

func updateCollection(exercises []types.Exercise, collection *mongo.Collection) {
	var models []mongo.WriteModel
	for _, exercise := range exercises {
		filter := bson.M{"id": exercise.ID}
		update := bson.M{"$set": exercise}
		model := mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true)
		models = append(models, model)
	}
	opts := options.BulkWrite().SetOrdered(false)
	_, err := collection.BulkWrite(context.TODO(), models, opts)
	if err != nil {
		log.Fatalf("Failed to perform bulk write: %v", err)
		return
	}
	fmt.Println("Exercises updated!")
}

func (db *Database) fetchExercises() ([]types.Exercise, error) {
	url := "https://exercisedb.p.rapidapi.com/exercises?limit=0&offset=0"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("x-rapidapi-host", "exercisedb.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", db.apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	var exercises []types.Exercise
	if err := json.NewDecoder(resp.Body).Decode(&exercises); err != nil {
		return nil, err
	}

	return exercises, nil
}

func getClient(connectionString string) *mongo.Client {
	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("Unable to connect to MongoDB: %v", err)
	}

	return client
}
