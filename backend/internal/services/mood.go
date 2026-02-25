package services

import (
	"database/sql"
	"encoding/json"

	"mood-generator/internal/models"
)

func ProcessMood(db *sql.DB, userInput string) (*models.MoodResult, error) {
	// 1. Сохранить запрос
	var requestID int
	err := db.QueryRow(
		`INSERT INTO mood_requests (user_input) VALUES ($1) RETURNING id`,
		userInput,
	).Scan(&requestID)
	if err != nil {
		return nil, err
	}

	// 2. Анализ через Claude
	analysis, err := AnalyzeMood(userInput)
	if err != nil {
		return nil, err
	}

	// 3. Обновить mood_label и energy в запросе
	db.Exec(
		`UPDATE mood_requests SET mood_label=$1, energy=$2 WHERE id=$3`,
		analysis.MoodLabel, analysis.Energy, requestID,
	)

	// 4. Треки из Spotify
	tracks, err := GetTracks(analysis.MoodLabel)
	if err != nil {
		tracks = []models.Track{} // не ломаем ответ если Spotify недоступен
	}

	// 5. Сохранить результат
	paletteJSON, _ := json.Marshal(analysis.Palette)
	tracksJSON, _ := json.Marshal(tracks)

	var resultID int
	db.QueryRow(
		`INSERT INTO mood_results (request_id, palette, quote, quote_author, tracks)
		 VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		requestID, paletteJSON, analysis.Quote, analysis.QuoteAuthor, tracksJSON,
	).Scan(&resultID)

	return &models.MoodResult{
		ID:          resultID,
		RequestID:   requestID,
		MoodLabel:   analysis.MoodLabel,
		Energy:      analysis.Energy,
		Palette:     analysis.Palette,
		Quote:       analysis.Quote,
		QuoteAuthor: analysis.QuoteAuthor,
		Tracks:      tracks,
	}, nil
}
