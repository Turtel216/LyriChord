package request

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Turtel216/LyriChord/internal/format"
)

const (
	// LyricsURL stores the base url to the lyrics api
	LyricsURL = "https://api.lyrics.ovh/v1/"
	// TabURL stores the base url to the tabs api
	TabURL = "" //TODO
)

// Struct to hold API response
type LyricsResponse struct {
	Lyrics string `json:"lyrics,omitempty"` // `omitempty` ensures it is omitted if empty
	Error  string `json:"error,omitempty"`
}

func fetchLyrics(baseURL, song, artist string) (*LyricsResponse, error) {
	// Format the URL
	apiURL := fmt.Sprintf("%s/%s/%s", baseURL, song, artist)

	// Make HTTP GET request
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("failed to make API request: %v", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	// Parse JSON response
	var result LyricsResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	return &result, nil
}

func RequestLyrics(title, artist string) string {
	apiResponse, err := fetchLyrics(LyricsURL, title, artist)
	if err != nil {
		log.Printf("Error fetching from lyrics api: %v", err)
		return format.FormatError("Internal error, could not fetch from lyrics API")
	}

	return format.FormatSong(title, artist, apiResponse.Lyrics)
}
