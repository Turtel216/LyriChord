package main

import (
	"log"
	"os"
	"strings"

	"github.com/Turtel216/LyriChord/internal/request"
	"github.com/Turtel216/LyriChord/internal/utils"
	"github.com/bwmarrin/discordgo"
)

var logger = log.New(os.Stdout, "[BOT] ", log.Ldate|log.Ltime)

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

	if !strings.Contains(m.Content, "!lyrics") {
		logger.Println("Not a !lyrics command. Message Skipped")
		return
	}

	logger.Println("Identified !lyrics Ping. Message")
	// Create the private channel with the user who sent the message.
	channel, err := s.UserChannelCreate(m.Author.ID)
	if err != nil {
		// If an error occurred, we failed to create the channel.
		//
		// Some common causes are:
		// 1. We don't share a server with the user (not possible here).
		// 2. We opened enough DM channels quickly enough for Discord to
		//    label us as abusing the endpoint, blocking us from opening
		//    new ones.
		logger.Println("error creating channel:", err)
		s.ChannelMessageSend(
			m.ChannelID,
			"Something went wrong while sending the DM!",
		)
		return
	}

	logger.Println("User channel created")

	_, song, artist, err := utils.ParseLyricsCommand(m.Content)
	if err != nil {
		logger.Printf("Error parsing message: %s", err)
		s.ChannelMessageSend(m.ChannelID, "# Incorrect command format. The correct format is !lyrics <song> by <artist>")
		return
	}

	logger.Println("Parsed Command")

	// Get the formated lyric
	res := request.RequestLyrics(song, artist)

	logger.Println("Got API response")

	// Send the lyrics
	_, err = s.ChannelMessageSend(channel.ID, res)
	if err != nil {
		// If an error occurred, we failed to send the message.
		//
		// It may occur either when we do not share a server with the
		// user (highly unlikely as we just received a message) or
		// the user disabled DM in their settings (more likely).
		logger.Println("error sending DM message:", err)
		s.ChannelMessageSend(
			m.ChannelID,
			"Failed to send you a DM. "+
				"Did you disable DM in your privacy settings?",
		)
	}

	logger.Println("Replyed to message")
}
