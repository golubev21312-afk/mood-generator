package models

type MoodRequest struct {
	ID        int    `json:"id"`
	UserInput string `json:"user_input"`
	MoodLabel string `json:"mood_label"`
	Energy    int    `json:"energy"`
	CreatedAt string `json:"created_at"`
}

type Color struct {
	Hex  string `json:"hex"`
	Name string `json:"name"`
	Role string `json:"role"`
}

type Track struct {
	Title      string `json:"title"`
	Artist     string `json:"artist"`
	SpotifyURL string `json:"spotify_url"`
	Cover      string `json:"cover"`
}

type MoodResult struct {
	ID          int     `json:"id"`
	RequestID   int     `json:"request_id"`
	MoodLabel   string  `json:"mood_label"`
	Energy      int     `json:"energy"`
	Palette     []Color `json:"palette"`
	Quote       string  `json:"quote"`
	QuoteAuthor string  `json:"quote_author"`
	Tracks      []Track `json:"tracks"`
	CreatedAt   string  `json:"created_at"`
}

// Ответ Claude
type ClaudeAnalysis struct {
	MoodLabel   string  `json:"mood_label"`
	Energy      int     `json:"energy"`
	Palette     []Color `json:"palette"`
	Quote       string  `json:"quote"`
	QuoteAuthor string  `json:"quote_author"`
}
