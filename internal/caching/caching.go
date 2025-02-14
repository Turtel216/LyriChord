// Package caching provides a simple in-memory caching mechanism for storing lyrics
// with automatic expiration. It includes a thread-safe cache and a request deduplication
// mechanism to prevent redundant API calls.
package caching

import (
	"strings"
	"sync"
	"time"

	"golang.org/x/sync/singleflight"
)

// CacheItem represents an individual cache entry, storing the lyrics
// and their expiration time.
type CacheItem struct {
	Lyrics     string    // Cached lyrics content
	Expiration time.Time // Time when the cache entry expires
}

// LyricsCache is a thread-safe map used for storing cached lyrics.
// The sync.Map structure ensures safe concurrent access.
var LyricsCache sync.Map

// RequestGroup prevents multiple simultaneous requests for the same lyrics,
// reducing redundant API calls using the singleflight package.
var RequestGroup singleflight.Group

// StartCacheCleanup initiates a background goroutine that periodically removes
// expired cache entries. The cleanup runs at the specified interval.
func StartCacheCleanup(interval time.Duration) {
	go func() {
		for {
			time.Sleep(interval) // Wait for the next cleanup cycle

			// Iterate over cache entries and remove expired ones
			LyricsCache.Range(func(key, value interface{}) bool {
				item := value.(CacheItem)
				if time.Now().After(item.Expiration) {
					LyricsCache.Delete(key) // Remove expired item
				}
				return true
			})
		}
	}()
}

// GetCacheKey generates the cache key for a given song + artist combination
func GetCacheKey(song, artist string) string {
	return strings.ToLower(song + ":" + artist)
}
