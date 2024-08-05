package main

import (
	"fmt"
	"log"
	"server/src"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	fmt.Println("Starting up server...")
	database := src.NewDatabase()

	location, err := time.LoadLocation("America/Chicago")
	if err != nil {
		log.Fatalf("Failed to load location: %v", err)
		return
	}
	scheduler := cron.New(cron.WithLocation(location))
	_, _ = scheduler.AddFunc("01 12 * * *", func() {
		database.UpdateExercises()
	})
	scheduler.Start()

	select {}
}
