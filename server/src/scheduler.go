package src

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

type Scheduler struct {
	cron *cron.Cron
	db   *Database
}

func NewScheduler(db *Database) *Scheduler {
	location, err := time.LoadLocation("America/Chicago")
	if err != nil {
		log.Fatalf("Failed to load location: %v", err)
	}
	scheduler := cron.New(cron.WithLocation(location))
	_, _ = scheduler.AddFunc("01 12 * * *", func() {
		db.UpdateExercises()
	})
	return &Scheduler{
		cron: scheduler,
		db:   db,
	}
}

func (s *Scheduler) Start(ctx context.Context) {
	s.cron.Start()

	<-ctx.Done()
	s.cron.Stop()
	fmt.Println("Scheduler stopped")
}
