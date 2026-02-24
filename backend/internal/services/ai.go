package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"

	"mood-generator/internal/models"
)

const deepseekPrompt = `Ты — анализатор настроения.
Пользователь описал своё состояние: "%s"

Верни ТОЛЬКО JSON без пояснений:
{
  "mood_label": "грустный",
  "energy": 3,
  "palette": [
    {"hex": "#2C3E50", "name": "Глубокий синий", "role": "основной"},
    {"hex": "#8E9BA8", "name": "Туманный серый", "role": "фон"},
    {"hex": "#E8D5B0", "name": "Тёплый беж",     "role": "акцент"}
  ],
  "quote": "Даже самая тёмная ночь заканчивается рассветом.",
  "quote_author": "Виктор Гюго"
}

mood_label — один из:
радостный | грустный | тревожный | спокойный | злой | вдохновлённый | усталый

energy — число от 1 (полная апатия) до 10 (эйфория)`

func AnalyzeMood(userInput string) (*models.ClaudeAnalysis, error) {
	prompt := fmt.Sprintf(deepseekPrompt, userInput)

	body, _ := json.Marshal(map[string]any{
		"model":      "llama-3.3-70b-versatile",
		"max_tokens": 1024,
		"messages": []map[string]string{
			{"role": "user", "content": prompt},
		},
	})

	req, _ := http.NewRequest("POST", "https://api.groq.com/openai/v1/chat/completions", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("GROQ_API_KEY"))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)

	var result struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}
	json.Unmarshal(data, &result)

	if len(result.Choices) == 0 {
		return nil, fmt.Errorf("ошибка Groq API (статус %d): %s", resp.StatusCode, string(data))
	}

	content := result.Choices[0].Message.Content
	// Groq sometimes wraps JSON in markdown code block, strip it
	if idx := strings.Index(content, "{"); idx > 0 {
		content = content[idx:]
	}
	re := regexp.MustCompile("```[\\s\\S]*$")
	content = re.ReplaceAllString(content, "")
	content = strings.TrimSpace(content)

	var analysis models.ClaudeAnalysis
	if err := json.Unmarshal([]byte(content), &analysis); err != nil {
		return nil, fmt.Errorf("ошибка парсинга JSON от Groq: %w", err)
	}

	return &analysis, nil
}
