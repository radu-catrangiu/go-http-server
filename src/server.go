package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"server/src/handlers"
	"server/src/dependencies"
	"sync"
	"syscall"
)

func getServerAndPort() (*http.Server, string) {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		log.Fatal("ERROR: No port was provided")
	}

	serverInstance := &http.Server{
		Addr: ":" + port,
	}

	return serverInstance, port
}

func registerShutdownHandlers(serverPtr *http.Server, wg *sync.WaitGroup, cleanUpCallbacks []func(*sync.WaitGroup)) {
	var serverInstance http.Server = *serverPtr

	// Register cleanup handlers
	log.Println("Registering", len(cleanUpCallbacks), "cleanup callbacks")
	wg.Add(len(cleanUpCallbacks))
	serverInstance.RegisterOnShutdown(func() {
		log.Println("Starting dependencies shutdown...")
		for _, cleanUpCallback := range cleanUpCallbacks {
			cleanUpCallback(wg)
		}
	})

	// Start listening for SIGTERM & SIGINT
	log.Println("Registering SIGTERM & SIGINT handler...")
	wg.Add(1)
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)

	// Wait for the signal -> this blocks the main thread
	<-signalChannel
	log.Println("Got Interrupt signal. Shutdown initiating...")

	shutdownCtx := context.Background()
	err := serverInstance.Shutdown(shutdownCtx)
	if err != nil {
		log.Fatalf("HTTP Shutdown error: %v", err)
	}
	log.Println("HTTP Server Shutdown complete")
	wg.Done()
}

func startServer(serverPtr *http.Server, port string) {
	var serverInstance http.Server = *serverPtr

	log.Println("Server is starting on port ", port)
	err := serverInstance.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal("Could not start server. Error: ", err)
	}
}

func main() {
	log.Println("Process started with PID:", os.Getpid())

	env, cleanupCallbacks, err := dependencies.Initialize()
	if err != nil {
		log.Fatalln("Failed to initialize dependencies", err)
	}

	serverInstance, port := getServerAndPort()

	handlers.RegisterHandlers(env)

	wg := new(sync.WaitGroup)

	go startServer(serverInstance, port)

	registerShutdownHandlers(serverInstance, wg, cleanupCallbacks)

	wg.Wait()
	log.Println("Shutdown complete. Process exiting...")
}
