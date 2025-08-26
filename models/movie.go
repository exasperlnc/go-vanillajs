package models

type Movie struct {
	ID          int
	TMDB_ID     int
	Title       string
	ReleaseYear int
	Tagline     string
	Genres			[]Genre
	Overview    *string
	Score       *float32
	Popularity  *float32
	Keywords    []string
	Language    *string
	PosterURL   *string
	TrailerURL  *string
	Casting     []Actor
}
