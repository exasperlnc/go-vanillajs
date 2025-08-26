package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/exasperlnc/go-vanillajs/data"
	"github.com/exasperlnc/go-vanillajs/handlers"
	"github.com/exasperlnc/go-vanillajs/logger"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
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

	err := godotenv.Load()
  if err != nil {
      log.Fatal("Error loading .env file")
  }

  dbUser := os.Getenv("DB_USER")
  dbPassword := os.Getenv("DB_PASSWORD")
  dbName := os.Getenv("DB_NAME")
  dbHost := os.Getenv("DB_HOST")
  dbPort := os.Getenv("DB_PORT")

  connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
      dbUser, dbPassword, dbHost, dbPort, dbName)

  db, err := sql.Open("postgres", connStr)
  if err != nil {
      log.Fatal(err)
  }
  defer db.Close()

	movieRepo, err := data.NewMovieRepository(db, logInstance)
	if err != nil {
		log.Fatalf("Failed to create movie repository: %v", err)
	}

  // Test the connection
  err = db.Ping()
  if err != nil {
      log.Fatal(err)
  }
  fmt.Println("Connected to the database!")

	movieHandler := handlers.MovieHandler{Storage: movieRepo, Logger: logInstance}

	http.HandleFunc("/api/movies/top", movieHandler.GetTopMovies)
	http.HandleFunc("/api/movies/random", movieHandler.GetRandomMovies)


	http.Handle("/", http.FileServer(http.Dir("public")))
	const addr = ":8080"
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Server Failed: %v", err)
		logInstance.Error("Server Failed", err)
	}
}