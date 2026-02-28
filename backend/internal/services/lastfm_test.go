package services

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetTracks_Success(t *testing.T) {
	body := `{
		"tracks": {
			"track": [
				{
					"name": "Skinny Love",
					"url": "https://www.last.fm/music/Bon+Iver/_/Skinny+Love",
					"artist": {"name": "Bon Iver"},
					"image": [
						{"#text": "https://img.example.com/small.jpg", "size": "small"},
						{"#text": "https://img.example.com/large.jpg", "size": "large"}
					]
				},
				{
					"name": "Holocene",
					"url": "https://www.last.fm/music/Bon+Iver/_/Holocene",
					"artist": {"name": "Bon Iver"},
					"image": [
						{"#text": "https://img.example.com/large2.jpg", "size": "large"}
					]
				}
			]
		}
	}`
	restore := mockHTTP(200, body)
	defer restore()

	tracks, err := GetTracks("грустный")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tracks) != 2 {
		t.Fatalf("expected 2 tracks, got %d", len(tracks))
	}
	if tracks[0].Title != "Skinny Love" {
		t.Errorf("title: got %q, want %q", tracks[0].Title, "Skinny Love")
	}
	if tracks[0].Artist != "Bon Iver" {
		t.Errorf("artist: got %q, want %q", tracks[0].Artist, "Bon Iver")
	}
	if tracks[0].Cover != "https://img.example.com/large.jpg" {
		t.Errorf("cover: got %q, want large image URL", tracks[0].Cover)
	}
	if tracks[0].SpotifyURL != "https://www.last.fm/music/Bon+Iver/_/Skinny+Love" {
		t.Errorf("url: got %q", tracks[0].SpotifyURL)
	}
}

func TestGetTracks_UnknownMoodFallsBackToChill(t *testing.T) {
	var capturedURL string
	old := http.DefaultTransport
	http.DefaultTransport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		capturedURL = r.URL.String()
		rec := httptest.NewRecorder()
		rec.WriteHeader(200)
		rec.WriteString(`{"tracks":{"track":[]}}`)
		return rec.Result(), nil
	})
	defer func() { http.DefaultTransport = old }()

	_, err := GetTracks("неизвестное_настроение")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(capturedURL, "tag=chill") {
		t.Errorf("expected fallback tag=chill in URL, got: %s", capturedURL)
	}
}

func TestGetTracks_EmptyTrackList(t *testing.T) {
	restore := mockHTTP(200, `{"tracks":{"track":[]}}`)
	defer restore()

	tracks, err := GetTracks("грустный")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tracks) != 0 {
		t.Errorf("expected 0 tracks, got %d", len(tracks))
	}
}

func TestGetTracks_MoodToTagMapping(t *testing.T) {
	cases := []struct {
		mood        string
		expectedTag string
	}{
		{"грустный", "sad"},
		{"радостный", "happy"},
		{"тревожный", "ambient"},
		{"спокойный", "chill"},
		{"злой", "rock"},
		{"вдохновлённый", "epic"},
		{"усталый", "lo-fi"},
	}

	for _, tc := range cases {
		t.Run(tc.mood, func(t *testing.T) {
			var capturedURL string
			old := http.DefaultTransport
			http.DefaultTransport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
				capturedURL = r.URL.String()
				rec := httptest.NewRecorder()
				rec.WriteHeader(200)
				rec.WriteString(`{"tracks":{"track":[]}}`)
				return rec.Result(), nil
			})
			defer func() { http.DefaultTransport = old }()

			GetTracks(tc.mood)

			if !strings.Contains(capturedURL, "tag="+tc.expectedTag) {
				t.Errorf("mood %q: expected tag=%s in URL, got: %s", tc.mood, tc.expectedTag, capturedURL)
			}
		})
	}
}
