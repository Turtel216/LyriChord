// Package request provides functionality for fetching and formatting song lyrics
// from external APIs. It includes functions to make API requests, handle responses,
// and format the retrieved lyrics for display.
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
	// LyricsURL is the base URL for the lyrics API.
	LyricsURL = "https://api.lyrics.ovh/v1"
	// TabURL is the base URL for the tabs API (currently unimplemented).
	TabURL = "" // TODO: Define the correct URL
)

// LyricsResponse represents the response structure from the lyrics API.
type LyricsResponse struct {
	Lyrics string `json:"lyrics,omitempty"` // Lyrics text (omitted if empty).
	Error  string `json:"error,omitempty"`  // Error message (omitted if empty).
}

// fetchLyrics retrieves song lyrics from the specified API base URL.
//
// It constructs the API request URL using the provided song title and artist,
// performs an HTTP GET request, and parses the JSON response.
//
// Returns a pointer to LyricsResponse on success or an error if the request fails.
func fetchLyrics(baseURL, song, artist string) (*LyricsResponse, error) {
	// Construct the API request URL
	apiURL := fmt.Sprintf("%s/%s/%s", baseURL, url.QueryEscape(artist), url.QueryEscape(song))

	// Make an HTTP GET request
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("failed to make API request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	// Handle non-200 status codes
	if resp.StatusCode != http.StatusOK {
		strBody := string(body)
		if strings.Contains(strBody, "No lyrics found") {
			return nil, SongNotFound
		}
		return nil, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, strBody)
	}

	// Parse the JSON response
	var result LyricsResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	return &result, nil
}

// RequestLyrics fetches and formats song lyrics based on the given title and artist.
//
// If lyrics are found, they are formatted and returned as a string.
// If the song is not found, an error message is logged and a formatted error message is returned.
//
// This function ensures proper error handling and logs any issues encountered during the request.
func RequestLyrics(title, artist string) string {
	apiResponse, err := fetchLyrics(LyricsURL, title, artist)
	if err != nil {
		if errors.Is(err, SongNotFound) {
			log.Printf("Song not found by API")
			msg := fmt.Sprintf("Song `%s` by `%s` **not found**", title, artist)
			return format.FormatError(msg)
		}

		log.Printf("Error fetching from lyrics API: %v", err)
		return format.FormatError("Internal error, could not fetch from lyrics API")
	}

	return format.FormatSong(title, artist, apiResponse.Lyrics)
}
