package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/Turtel216/LyriChord/internal/format"
)

const (
	// LyricsURL stores the base url to the lyrics api
	LyricsURL = "https://api.lyrics.ovh/v1"
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
	apiURL := fmt.Sprintf("%s/%s/%s", baseURL, url.QueryEscape(artist), url.QueryEscape(song))

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

	// Check for non-200 status codes
	if resp.StatusCode != http.StatusOK {
		strBody := string(body)

		if strings.Contains(strBody, "No lyrics found") {
			return nil, SongNotFound
		}

		return nil, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
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
		if errors.Is(err, SongNotFound) {
			log.Printf("Song not found by API")
			msg := fmt.Sprintf("Song `%s` by `%s` **not found**", title, artist)
			return format.FormatError(msg)
		}

		log.Printf("Error fetching from lyrics api: %v", err)
		return format.FormatError("Internal error, could not fetch from lyrics API")
	}

	return format.FormatSong(title, artist, apiResponse.Lyrics)
}
