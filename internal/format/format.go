package format

import "fmt"

func FormatSong(title, artist, lyrics string) string {
	return fmt.Sprintf("# %s \n ## By %s \n\n %s", title, artist, lyrics)
}
