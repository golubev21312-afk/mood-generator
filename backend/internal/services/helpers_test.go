package services

import (
	"net/http"
	"net/http/httptest"
)

// roundTripFunc позволяет использовать функцию как http.RoundTripper
type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}

// mockHTTP заменяет http.DefaultTransport на заглушку с заданным статусом и телом.
// Возвращает функцию-восстановитель, которую нужно вызвать через defer.
func mockHTTP(statusCode int, body string) func() {
	old := http.DefaultTransport
	http.DefaultTransport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		rec := httptest.NewRecorder()
		rec.WriteHeader(statusCode)
		rec.WriteString(body)
		return rec.Result(), nil
	})
	return func() { http.DefaultTransport = old }
}
