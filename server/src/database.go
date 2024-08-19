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

func (db *Database) FetchWorkouts() ([]types.Workout, error) {
	collection := db.Client.Database("gym_management").Collection("workouts")
	sort := bson.M{"creationDate": -1}
	opts := options.Find().SetSort(sort).SetLimit(10)

	cursor, err := collection.Find(context.Background(), bson.D{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var workouts []types.Workout
	if err := cursor.All(context.Background(), &workouts); err != nil {
		return nil, err
	}
	return workouts, nil
}

func getClient(connectionString string) *mongo.Client {
	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("Unable to connect to MongoDB: %v", err)
	}

	return client
}
