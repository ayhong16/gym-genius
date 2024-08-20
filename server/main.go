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

	api := src.NewAPI()

	database := src.NewDatabase(api)
	defer database.Disconnect()

	scheduler := src.NewScheduler(database)

	src.StartRouter(database, api)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		scheduler.Start(ctx)
	}()

	<-sigCh
	fmt.Println("Shutting down server...")

	cancel()
	fmt.Println("Cancelled context")
	wg.Wait()
}
