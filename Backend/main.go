package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"

	middleware "social-network/Middleware"
	routes "social-network/Routes"
	utils "social-network/Utils"
)

func init() {
	args := os.Args
	if len(args) != 2 {
		return
	}

	if strings.ToLower(args[1]) == "--loaddata" || strings.ToLower(args[1]) == "-l" {
		db, err := utils.OpenDb("sqlite3", "./Database/Database.sqlite")
		if err != nil {
			fmt.Println(err)
		}
		defer db.Close()

		start := time.Now()
		if err = utils.LoadData(db); err != nil {
			fmt.Println(err)
		}
		end := time.Now()
		fmt.Println(end.Sub(start))
	}
}

func main() {
	// We create a log file and redirect the stdout to the new file
	logFile, _ := os.Create("./Log/" + time.Now().Format("2006-01-02__15-04-05") + ".log")
	defer logFile.Close()

	log.SetOutput(logFile)

	// We launch the server
	mux := http.NewServeMux()

	// Enchaîner les middlewares
	handler := middleware.SetHeaderAccessControll(
		middleware.LookMethod(mux),
	)

	// We set all the endpoints
	routes.Routes(mux)

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	// We set the time out limit
	srv := &http.Server{
		Handler:      handler,
		Addr:         "localhost:8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
		TLSConfig:    tlsConfig,
	}

	go func() {
		certFile := "Key/cert.pem"
		keyFile := "Key/key.pem"

		log.Printf("Server listening on https://%s", srv.Addr)
		fmt.Printf("\033[96mServer started at: https://%s\033[0m\n", srv.Addr)

		if err := srv.ListenAndServeTLS(certFile, keyFile); err != nil {
			log.Fatalf("Error starting TLS server: %v", err)
		}
	}()

	// Signal to capture a clean shutdown (SIGINT/SIGTERM)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	fmt.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// We close correctly the server
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %s\n", err)
	}

	// We reset the stdout to is normal status
	fmt.Println("Server exiting")
	log.Println("Server exiting")

}
