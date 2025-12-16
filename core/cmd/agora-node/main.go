// core/cmd/agora-node/main.go
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Jaime2003z/Agora/core/node"
	"github.com/gorilla/mux"
)

func main() {
	// Create node
	n, err := node.NewNode()
	if err != nil {
		log.Fatalf("Failed to create node: %v", err)
	}
	defer n.Close()

	// Setup protocols and discovery
	n.RegisterProtocols()
	n.SetupDiscovery()

	// Setup HTTP server
	router := mux.NewRouter()
	n.SetupAPIRoutes(router)

	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	// Start HTTP server
	go func() {
		log.Println("Starting HTTP server on :3000")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	// Wait for interrupt signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Println("Shutting down node...")
	if err := server.Shutdown(context.Background()); err != nil {
		log.Printf("HTTP server shutdown error: %v", err)
	}
	log.Println("Node stopped")
}
