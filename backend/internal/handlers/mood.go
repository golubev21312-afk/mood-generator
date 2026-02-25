package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"mood-generator/internal/services"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	DB *sql.DB
}

// POST /api/mood
func (h *Handler) PostMood(c *gin.Context) {
	var body struct {
		Input string `json:"input" binding:"required"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "поле input обязательно"})
		return
	}

	result, err := services.ProcessMood(h.DB, body.Input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// GET /api/history
func (h *Handler) GetHistory(c *gin.Context) {
	rows, err := h.DB.Query(
		`SELECT r.id, r.user_input, r.mood_label, r.energy, r.created_at,
		        res.palette, res.quote, res.quote_author, res.tracks
		 FROM mood_requests r
		 LEFT JOIN mood_results res ON res.request_id = r.id
		 ORDER BY r.created_at DESC LIMIT 20`,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var items []gin.H
	for rows.Next() {
		var id, energy int
		var userInput, moodLabel, createdAt string
		var paletteJSON, tracksJSON []byte
		var quote, quoteAuthor string

		rows.Scan(&id, &userInput, &moodLabel, &energy, &createdAt,
			&paletteJSON, &quote, &quoteAuthor, &tracksJSON)

		var palette, tracks any
		json.Unmarshal(paletteJSON, &palette)
		json.Unmarshal(tracksJSON, &tracks)

		items = append(items, gin.H{
			"id": id, "user_input": userInput, "mood_label": moodLabel,
			"energy": energy, "created_at": createdAt,
			"palette": palette, "quote": quote, "quote_author": quoteAuthor,
			"tracks": tracks,
		})
	}

	c.JSON(http.StatusOK, items)
}

// GET /api/mood/:id
func (h *Handler) GetMoodByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "невалидный id"})
		return
	}

	var userInput, moodLabel, quote, quoteAuthor, createdAt string
	var energy int
	var paletteJSON, tracksJSON []byte

	err = h.DB.QueryRow(
		`SELECT r.user_input, r.mood_label, r.energy, r.created_at,
		        res.palette, res.quote, res.quote_author, res.tracks
		 FROM mood_requests r
		 LEFT JOIN mood_results res ON res.request_id = r.id
		 WHERE r.id = $1`, id,
	).Scan(&userInput, &moodLabel, &energy, &createdAt,
		&paletteJSON, &quote, &quoteAuthor, &tracksJSON)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "не найдено"})
		return
	}

	var palette, tracks any
	json.Unmarshal(paletteJSON, &palette)
	json.Unmarshal(tracksJSON, &tracks)

	c.JSON(http.StatusOK, gin.H{
		"id": id, "user_input": userInput, "mood_label": moodLabel,
		"energy": energy, "created_at": createdAt,
		"palette": palette, "quote": quote, "quote_author": quoteAuthor,
		"tracks": tracks,
	})
}
