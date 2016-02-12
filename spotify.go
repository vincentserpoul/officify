package officify

import (
	"github.com/everdev/mack"
)

// Play will locally play the specified song
func Play(spotifyURI string) error {
	err := mack.Tell("Spotify", `play track "`+spotifyURI+`"`)
	if err != nil {
		return err
	}

	return nil
}

// Next will skip the current song and go to next in spotify
func Next() error {
	err := mack.Tell("Spotify", `next track`)
	if err != nil {
		return err
	}

	return nil
}
