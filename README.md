# 🎵 LyriChord Discord Bot

> [!CAUTION]
> The Bot is still in early development. Many features are not tested or yet to be implemented.

A Discord bot written in Golang that fetches song lyrics and guitar tabs on request. The bot utilizes the [Lyrics.ovh API](https://lyricsovh.docs.apiary.io/#) for lyrics and the [Songsterr API](https://www.songsterr.com/a/wa/api)(See [issue #2](https://github.com/Turtel216/LyriChord/issues/2)) for guitar tabs.

## 🚀 Features
- Retrieve lyrics of a song by specifying the song title and artist.
- Fetch guitar tabs for a song.
- Lightweight and fast response times.
- Easy to deploy and configure.

## 🛠️ Installation

### Prerequisites
- Go (latest stable version)
- A Discord bot token (Get one from the [Discord Developer Portal](https://discord.com/developers/applications))

### Clone the Repository
```sh
git clone https://github.com/Turtel216/LyriChord.git
cd LyriChord
```

### Install Dependencies
```sh
go mod tidy
```

### Configuration
Create a `.env` file in the root directory and add your bot token:
```env
TOKEN=your_discord_bot_token_here
```

### Run the Bot
```sh
go run cmd/main.go
```

## 📜 Commands
| Command | Description |
|---------|-------------|
| `!lyrics <song> by <artist>` | Fetches the lyrics of the requested song. |
| `!tabs <song>` | Retrieves guitar tabs for the requested song. |

## 🔧 Deployment
You can deploy the bot using Docker:
```sh
docker build -t discord-lyrics-tabs-bot .
docker run -d --env-file .env discord-lyrics-tabs-bot
```

## 🤝 Contributing
Feel free to submit issues or pull requests if you have any improvements or bug fixes.

## 📜 License
This project is licensed under the GPL2 License.

---

### 🎶 Enjoy your music journey with this bot! 🎸
