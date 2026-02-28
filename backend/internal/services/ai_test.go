package services

import (
	"encoding/json"
	"strings"
	"testing"

	"mood-generator/internal/models"
)

// deepseekResponse оборачивает анализ в формат ответа DeepSeek (OpenAI-compatible)
func deepseekResponse(analysis models.ClaudeAnalysis) string {
	analysisJSON, _ := json.Marshal(analysis)
	resp := map[string]any{
		"choices": []map[string]any{
			{"message": map[string]string{"content": string(analysisJSON)}},
		},
	}
	b, _ := json.Marshal(resp)
	return string(b)
}

func TestAnalyzeMood_Success(t *testing.T) {
	expected := models.ClaudeAnalysis{
		MoodLabel: "грустный",
		Energy:    3,
		Palette: []models.Color{
			{Hex: "#2C3E50", Name: "Глубокий синий", Role: "основной"},
			{Hex: "#8E9BA8", Name: "Туманный серый", Role: "фон"},
		},
		Quote:       "Даже самая тёмная ночь заканчивается рассветом.",
		QuoteAuthor: "Виктор Гюго",
	}
	restore := mockHTTP(200, deepseekResponse(expected))
	defer restore()

	result, err := AnalyzeMood("мне грустно и тяжело")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.MoodLabel != expected.MoodLabel {
		t.Errorf("mood_label: got %q, want %q", result.MoodLabel, expected.MoodLabel)
	}
	if result.Energy != expected.Energy {
		t.Errorf("energy: got %d, want %d", result.Energy, expected.Energy)
	}
	if len(result.Palette) != len(expected.Palette) {
		t.Errorf("palette len: got %d, want %d", len(result.Palette), len(expected.Palette))
	}
	if result.Quote != expected.Quote {
		t.Errorf("quote: got %q, want %q", result.Quote, expected.Quote)
	}
	if result.QuoteAuthor != expected.QuoteAuthor {
		t.Errorf("quote_author: got %q, want %q", result.QuoteAuthor, expected.QuoteAuthor)
	}
}

func TestAnalyzeMood_EmptyChoices(t *testing.T) {
	restore := mockHTTP(200, `{"choices": []}`)
	defer restore()

	_, err := AnalyzeMood("тест")
	if err == nil {
		t.Fatal("expected error for empty choices, got nil")
	}
	if !strings.Contains(err.Error(), "Groq") {
		t.Errorf("expected error mentioning Groq, got: %v", err)
	}
}

func TestAnalyzeMood_InvalidJSONFromDeepSeek(t *testing.T) {
	resp := map[string]any{
		"choices": []map[string]any{
			{"message": map[string]string{"content": "это не JSON, а просто текст"}},
		},
	}
	b, _ := json.Marshal(resp)
	restore := mockHTTP(200, string(b))
	defer restore()

	_, err := AnalyzeMood("тест")
	if err == nil {
		t.Fatal("expected parse error, got nil")
	}
}

func TestAnalyzeMood_AllMoodLabels(t *testing.T) {
	labels := []string{"радостный", "грустный", "тревожный", "спокойный", "злой", "вдохновлённый", "усталый"}
	for _, label := range labels {
		t.Run(label, func(t *testing.T) {
			analysis := models.ClaudeAnalysis{
				MoodLabel:   label,
				Energy:      5,
				Palette:     []models.Color{{Hex: "#000000", Name: "Чёрный", Role: "основной"}},
				Quote:       "Цитата",
				QuoteAuthor: "Автор",
			}
			restore := mockHTTP(200, deepseekResponse(analysis))
			defer restore()

			result, err := AnalyzeMood("тест")
			if err != nil {
				t.Fatalf("unexpected error for label %q: %v", label, err)
			}
			if result.MoodLabel != label {
				t.Errorf("got %q, want %q", result.MoodLabel, label)
			}
		})
	}
}
