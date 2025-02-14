package main

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/Turtel216/LyriChord/internal/caching"
	"github.com/Turtel216/LyriChord/internal/request"
	"github.com/Turtel216/LyriChord/internal/utils"
	"github.com/bwmarrin/discordgo"
)

var logger = log.New(os.Stdout, "[BOT] ", log.Ldate|log.Ltime)

// lyricsCacheDuration holds the lifetime of lyrics in the in memory cache
const lyricsCacheDuration = 10 * time.Minute

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
//
// It is called whenever a message is created but only when it's sent through a
// server as we did not request IntentsDirectMessages.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	logger.Println("Message received")

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		logger.Println("Message created by bot. Message Skipped")
		return
	}

	// Check if the message contains a recognized command
	if !strings.Contains(m.Content, "!lyrics") {
		logger.Println("Not a !lyrics command. Message Skipped")
		return
	}

	logger.Println("Identified !lyrics Ping. Message")

	// Parse lyrics request queries
	_, song, artist, err := utils.ParseLyricsCommand(m.Content)
	if err != nil {
		logger.Printf("Error parsing message: %s", err)
		s.ChannelMessageSend(m.ChannelID, "# Incorrect command format. The correct format is !lyrics <song> by <artist>")
		return
	}

	logger.Println("Parsed Command")

	// Cache key (normalized)
	cacheKey := caching.GetCacheKey(song, artist)

	// Check cache first
	if cached, found := caching.LyricsCache.Load(cacheKey); found {
		logger.Println("Cache hit: Sending cached lyrics")
		sendLyrics(s, m.Author.ID, m.ChannelID, cached.(caching.CacheItem).Lyrics)
		return
	}

	// Use singleflight to prevent duplicate API calls
	res, err, _ := caching.RequestGroup.Do(cacheKey, func() (interface{}, error) {
		logger.Println("Cache miss: Fetching lyrics from API")
		lyrics := request.RequestLyrics(song, artist)
		caching.LyricsCache.Store(cacheKey, caching.CacheItem{Lyrics: lyrics, Expiration: time.Now().Add(lyricsCacheDuration)})
		return lyrics, nil
	})

	if err != nil {
		logger.Println("Error fetching lyrics:", err)
		s.ChannelMessageSend(m.ChannelID, "Failed to retrieve lyrics.")
		return
	}

	// Send lyrics
	sendLyrics(s, m.Author.ID, m.ChannelID, res.(string))
}

// sendLyrics is used to send lyrics to the user
func sendLyrics(s *discordgo.Session, userID, channelID, lyrics string) {
	channel, err := s.UserChannelCreate(userID)
	if err != nil {
		logger.Println("Error creating DM channel:", err)
		s.ChannelMessageSend(channelID, "Something went wrong while sending the DM!")
		return
	}

	_, err = s.ChannelMessageSend(channel.ID, lyrics)
	if err != nil {
		logger.Println("Error sending DM:", err)
		s.ChannelMessageSend(channelID, "Failed to send you a DM. Did you disable DMs in privacy settings?")
	}
	logger.Println("Message send")
}
