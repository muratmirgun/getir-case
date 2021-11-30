package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/muratmirgun/getir-case/internal/router/datastore"
	"github.com/muratmirgun/getir-case/internal/router/search"
	"github.com/muratmirgun/getir-case/internal/store/inmemory"
	"github.com/muratmirgun/getir-case/pkg"
)

func main() {
	// Create Mongo Server
	mongoServer := new(search.MongoDB)

	// Initialize InMemory Store
	holder := inmemory.New()
	dataHandler := datastore.New(holder)

	// Initialize the Logger
	pkg.Init()

	// InitHandlers
	http.HandleFunc("/holder", dataHandler.InMemory)
	http.HandleFunc("/search", mongoServer.ServeMongo)

	pkg.Info("started listening on port 8080")
	pkg.Info("Waiting request...")

	httpServer := &http.Server{
		Addr: ":8080",
	}
	go func() {
		if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("HTTP server ListenAndServe: %v", err)
		}
	}()

	signalChan := make(chan os.Signal, 1)

	signal.Notify(
		signalChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
	)

	<-signalChan
	log.Print("os.Interrupt - shutting down...\n")

	go func() {
		<-signalChan
		log.Fatal("os.Kill - terminating...\n")
	}()

	gracefullCtx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()

	if err := httpServer.Shutdown(gracefullCtx); err != nil {
		log.Printf("shutdown error: %v\n", err)
		defer os.Exit(1)
		return
	} else {
		log.Printf("gracefully stopped\n")
	}

	defer os.Exit(0)
}
