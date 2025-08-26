package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/exasperlnc/go-vanillajs/models"
)

type MovieHandler struct {
	// database and Logger
}

func (h *MovieHandler) GetTopMovies(w http.ResponseWriter, r *http.Request) {
	movies := []models.Movie{
		{
			ID: 1,
			TMDB_ID: 181,
			Title: "Logan has run",
			ReleaseYear: 2024,
			Genres: []models.Genre{{ID: 1, Name: "Action"}},
			Keywords: []string{"Logan", "Action", "Superhero"},
			Casting: []models.Actor{
				{ID: 1, FirstName: "Hugh", LastName: "Jackman"},
				{ID: 2, FirstName: "Patrick", LastName: "Stewart"},
			},
		},
		{
			ID: 2,
			TMDB_ID: 182,
			Title: "The Return of Wolverine",
			ReleaseYear: 2025,
			Genres: []models.Genre{{ID: 1, Name: "Action"}},
			Keywords: []string{"Wolverine", "Action", "Superhero"},
			Casting: []models.Actor{
				{ID: 1, FirstName: "Hugh", LastName: "Jackman"},
				{ID: 2, FirstName: "Patrick", LastName: "Stewart"},
			},
		},
	}
	h.writeJSONResponse(w, movies)
}

func (h *MovieHandler) writeJSONResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		// TODO: Log error
	}
}

func (h *MovieHandler) GetRandomMovies(w http.ResponseWriter, r *http.Request) {
	movies := []models.Movie{
		{
			ID: 1,
			TMDB_ID: 181,
			Title: "Logan has Random",
			ReleaseYear: 2024,
			Genres: []models.Genre{{ID: 1, Name: "Action"}},
			Keywords: []string{"Logan", "Action", "Superhero"},
			Casting: []models.Actor{
				{ID: 1, FirstName: "Hugh", LastName: "Jackman"},
				{ID: 2, FirstName: "Patrick", LastName: "Stewart"},
			},
		},
		{
			ID: 2,
			TMDB_ID: 182,
			Title: "The Return of Random",
			ReleaseYear: 2025,
			Genres: []models.Genre{{ID: 1, Name: "Action"}},
			Keywords: []string{"Wolverine", "Action", "Superhero"},
			Casting: []models.Actor{
				{ID: 1, FirstName: "Hugh", LastName: "Jackman"},
				{ID: 2, FirstName: "Patrick", LastName: "Stewart"},
			},
		},
	}
	h.writeJSONResponse(w, movies)
}