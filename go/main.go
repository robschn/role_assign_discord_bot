package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

var (
	// command line args
	botToken = flag.String("t", "token", "Token of the Discord bot from the developer portal.")
	appID    = flag.String("a", "appID", "Bot app ID of your Discord bot from the developer portal.")
	guildID  = flag.String("g", "guildID", "Bot guild ID from the Discord channel.")
)

func main() {
	flag.Parse()
	s, _ := discordgo.New("Bot " + *botToken)
	_, err := s.ApplicationCommandBulkOverwrite(*appID, *guildID, []*discordgo.ApplicationCommand{
		{
			Name:        "hello-world",
			Description: "Showcase of a basic slash command",
		},
	})
	if err != nil {
		// Handle the error
	}
	s.AddHandler(func(
		s *discordgo.Session,
		i *discordgo.InteractionCreate,
	) {
		data := i.ApplicationCommandData()
		switch data.Name {
		case "hello-world":
			err := s.InteractionRespond(
				i.Interaction,
				&discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "Hello world!",
					},
				},
			)
			if err != nil {
				// Handle the error
			}
		}
	})
	err = s.Open()
	if err != nil {
		// Handle the error
	}
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	err = s.Close()
	if err != nil {
		// Handle the error
	}
}
