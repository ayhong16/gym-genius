package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"server/src"
	"sync"
	"syscall"
)

func main() {
	fmt.Println("Starting up server...")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	database := src.NewDatabase()
	defer func() {
		database.Disconnect(ctx)
	}()

	scheduler := src.NewScheduler(database)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		scheduler.Start(ctx)
	}()

	<-sigCh
	fmt.Println("Shutting down server...")

	cancel()
	wg.Wait()
	fmt.Println("Server shut down successfully")
}
