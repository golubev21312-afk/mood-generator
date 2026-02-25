package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"mood-generator/internal/models"
)

var moodToTag = map[string]string{
	"грустный":      "sad",
	"радостный":     "happy",
	"тревожный":     "ambient",
	"спокойный":     "chill",
	"злой":          "rock",
	"вдохновлённый": "epic",
	"усталый":       "lo-fi",
}

func GetTracks(moodLabel string) ([]models.Track, error) {
	tag, ok := moodToTag[moodLabel]
	if !ok {
		tag = "chill"
	}

	apiKey := os.Getenv("LASTFM_API_KEY")
	apiURL := fmt.Sprintf(
		"https://ws.audioscrobbler.com/2.0/?method=tag.gettoptracks&tag=%s&api_key=%s&format=json&limit=6",
		tag, apiKey,
	)

	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Tracks struct {
			Track []struct {
				Name   string `json:"name"`
				URL    string `json:"url"`
				Artist struct {
					Name string `json:"name"`
				} `json:"artist"`
				Image []struct {
					Text string `json:"#text"`
					Size string `json:"size"`
				} `json:"image"`
			} `json:"track"`
		} `json:"tracks"`
	}

	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &result)

	tracks := make([]models.Track, 0, len(result.Tracks.Track))
	for _, t := range result.Tracks.Track {
		cover := ""
		for _, img := range t.Image {
			if img.Size == "large" {
				cover = img.Text
				break
			}
		}
		tracks = append(tracks, models.Track{
			Title:      t.Name,
			Artist:     t.Artist.Name,
			SpotifyURL: t.URL,
			Cover:      cover,
		})
	}
	return tracks, nil
}
