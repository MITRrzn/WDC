package main

import (
	"WDC/internal/database"
	"WDC/internal/endpoint"
	endpointRepository "WDC/internal/repository/endpoint"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	db, err := database.GetClient()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if dbCloseErr := db.Close(); dbCloseErr != nil {
			log.Printf("db close error: %v", dbCloseErr)
		}
	}()

	mux := http.NewServeMux()

	endpointRepo := endpointRepository.NewRepository(db)
	endpointSvc := endpoint.NewService(endpointRepo)
	endpointHandler := endpoint.NewHandler(endpointSvc)

	mux.HandleFunc("POST /endpoints", endpointHandler.CreateEndpoint)
	//mux.HandleFunc("GET /endpoints/{id}", endpointHandler.GetEndpoint)

	//mux.HandleFunc("POST /events", eventsHandler.CreateEvent)

	port := os.Getenv("APP_PORT")
	log.Println("Starting server at port", port)
	server := &http.Server{
		Addr:              fmt.Sprint(":", port),
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	go func() {
		listenErr := server.ListenAndServe()
		if listenErr != nil && !errors.Is(listenErr, http.ErrServerClosed) {
			log.Fatal(listenErr)
		}
	}()

	<-ctx.Done()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("Shutting down server")
	shutdownErr := server.Shutdown(shutdownCtx)
	if shutdownErr != nil {
		log.Printf("server shutdown error: %v", shutdownErr)
	}
	log.Println("Server stopped")
}
