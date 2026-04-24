package models

import "time"

const (
	MediaTypeMovie   = "movie"
	MediaTypeTVShow  = "tv_show"
	MediaTypeAnime   = "anime"
	MediaTypeOVA     = "ova"
	MediaTypeONA     = "ona"
	MediaTypeSpecial = "special"
)

type Media struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	MediaType   string    `json:"media_type"`
	TmdbID      int       `json:"tmdb_id"`
	PosterURL   string    `json:"poster_url"`
	Overview    string    `json:"overview"`
	ReleaseDate string    `json:"release_date"`
	Rating      float64   `json:"rating"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
