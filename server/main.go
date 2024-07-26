package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"server/initializers"
	"server/types"
	"time"

	"github.com/robfig/cron/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var APIKey string
var connectionString string

func init() {
	APIKey, connectionString = initializers.LoadEnvVariables()
}

func main() {
	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("Unable to connect to MongoDB: %v", err)
		return
	}

	fmt.Println("Connected to MongoDB!")

	location, _ := time.LoadLocation("America/Central")
	scheduler := cron.New(cron.WithLocation(location))
	_, err = scheduler.AddFunc("2 12 * * *", func() {
		updateExercises(client)
	})
	if err != nil {
		log.Fatalf("Failed to schedule job: %v", err)
		return
	}
	select {}
}

func updateExercises(client *mongo.Client) {
	collection := client.Database("gym_management").Collection("exercises")

	// if collection has no data, insert new exercises (only happens the first time)
	exercises, err := fetchExercises()
	if err != nil {
		log.Fatalf("Failed to fetch exercises: %v", err)
		return
	}

	for _, exercise := range exercises {
		filter := bson.M{"id": exercise.ID}
		update := bson.M{"$set": exercise}
		options := options.Update().SetUpsert(true)

		_, err := collection.UpdateOne(context.TODO(), filter, update, options)
		if err != nil {
			log.Fatalf("Failed to update or insert exercise %s: %v", exercise.ID, err)
			return
		}
	}
	fmt.Println("Exercises stored in MongoDB!")
}

func fetchExercises() ([]types.Exercise, error) {
	url := "https://exercisedb.p.rapidapi.com/exercises?limit=0&offset=0"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("x-rapidapi-host", "exercisedb.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", APIKey)

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
