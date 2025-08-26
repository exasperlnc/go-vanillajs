package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/exasperlnc/go-vanillajs/data"
	"github.com/exasperlnc/go-vanillajs/logger"
)

type MovieHandler struct {
	Storage data.MovieStorage
	Logger  *logger.Logger
}

func (h *MovieHandler) GetTopMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := h.Storage.GetTopMovies()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		h.Logger.Error("Failed to get top movies", err)
	}
	h.writeJSONResponse(w, movies)
}

func (h *MovieHandler) writeJSONResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		h.Logger.Error("Failed to write JSON response", err)
	}
}

func (h *MovieHandler) GetRandomMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := h.Storage.GetRandomMovies()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		h.Logger.Error("Failed to get random movies", err)
	}
	h.writeJSONResponse(w, movies)
}