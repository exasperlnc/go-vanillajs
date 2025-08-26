package main

import (
	"log"
	"net/http"

	"github.com/exasperlnc/go-vanillajs/logger"
)
func initializeLogger() *logger.Logger {
	logInstance, err := logger.NewLogger("movie.log")
	if err != nil {
		log.Fatalf("Could not initialize logger: %v", err)
	}
	defer logInstance.Close()
	return logInstance
}

func main() {
	logInstance := initializeLogger()	
	
	http.Handle("/", http.FileServer(http.Dir("public")))
	const addr = ":8080"
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Server Failed: %v", err)
		logInstance.Error("Server Failed", err)
	}
}