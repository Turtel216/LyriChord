package request

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock server responses
var successResponse = `{"lyrics": "Here are the lyrics of the song"}`
var errorResponse = `{"error": "No lyrics found"}`
var invalidJSONResponse = `{"lyrics": "Incomplete`

// Mock HTTP client using testify mocks
type MockClient struct {
	mock.Mock
}

func (m *MockClient) Get(url string) (*http.Response, error) {
	args := m.Called(url)
	return args.Get(0).(*http.Response), args.Error(1)
}

// Helper function to create a response recorder
func createMockResponse(body string, statusCode int) *http.Response {
	return &http.Response{
		StatusCode: statusCode,
		Body:       io.NopCloser(strings.NewReader(body)),
		Header:     make(http.Header),
	}
}

// Test successful lyrics fetch
func TestFetchLyrics_Success(t *testing.T) {
	// Create a test server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(successResponse))
	}))
	defer mockServer.Close()

	// Call fetchLyrics with the test server URL
	resp, err := fetchLyrics(mockServer.URL, "Imagine", "John-Lennon")

	assert.NoError(t, err, "Expected no error")
	assert.NotNil(t, resp, "Response should not be nil")
	assert.Equal(t, "Here are the lyrics of the song", resp.Lyrics)
	assert.Empty(t, resp.Error, "Error field should be empty")
}

// Test API returning an error response
func TestFetchLyrics_ErrorResponse(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(errorResponse))
	}))
	defer mockServer.Close()

	resp, err := fetchLyrics(mockServer.URL, "UnknownSong", "UnknownArtist")

	assert.NoError(t, err, "Expected no error despite error message")
	assert.NotNil(t, resp, "Response should not be nil")
	assert.Equal(t, "No lyrics found", resp.Error)
	assert.Empty(t, resp.Lyrics, "Lyrics field should be empty")
}

// Test handling of network failure
func TestFetchLyrics_NetworkFailure(t *testing.T) {
	invalidURL := "http://invalid.url"

	resp, err := fetchLyrics(invalidURL, "Song", "Artist")

	assert.Error(t, err, "Expected an error due to network failure")
	assert.Nil(t, resp, "Response should be nil on network failure")
}

// Test handling of invalid JSON response
func TestFetchLyrics_InvalidJSON(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(invalidJSONResponse))
	}))
	defer mockServer.Close()

	resp, err := fetchLyrics(mockServer.URL, "CorruptSong", "CorruptArtist")

	assert.Error(t, err, "Expected an error due to invalid JSON")
	assert.Nil(t, resp, "Response should be nil on JSON parsing failure")
}
