package botplay

import (
	"fmt"
	"log"
	"strings"

	"github.com/go-chat-bot/bot"
	"github.com/vincentserpoul/officify"
)

func play(command *bot.Cmd) (msg string, err error) {
	spotifyURI := strings.Trim(strings.Trim(command.Args[0], "<"), ">")
	msg = fmt.Sprintf("playing %s for %s!", spotifyURI, command.User.RealName)
	errSpot := officify.Play(spotifyURI)
	if errSpot != nil {
		log.Println(errSpot)
	}

	return msg, nil
}

func next(command *bot.Cmd) (msg string, err error) {
	errSpot := officify.Next()
	if errSpot != nil {
		log.Println(errSpot)
	}

	return ``, nil
}

func init() {
	bot.RegisterCommand(
		"play",
		"Play a specific song in spotify",
		"",
		play,
	)
	bot.RegisterCommand(
		"next",
		"Play next song in the playlist",
		"",
		next,
	)
}
