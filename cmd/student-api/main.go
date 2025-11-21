package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	config "github.com/Rishabhsingh78/student-api/internal"
)

func main() {
	cfg := config.MustLoad() // load config here
	// setup router
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to student api"))
	})
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}
	slog.Info("Server Started", slog.String("address", cfg.Addr))
	done := make(chan os.Signal, 1)
	// yaha pe jaise ctrl+c press kiya to done k channel me signal mil gya 
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGABRT) 

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("failed to start server")
		}
	}()
	<-done // yaah jaise hi usko yeh receive hua to server shutdown ho gya 
	slog.Info("Shutting down the server")
	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancle()
	if err := server.Shutdown(ctx); err != nil {
		slog.Error("failed to shutdown server", slog.String("error", err.Error()))
	}
	slog.Info("Server shutdown successfully")
}
