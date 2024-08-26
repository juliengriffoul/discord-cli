package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

var (
	Token     string
	ChannelID string
	Message   string
	Failure   bool
)

const (
	Red   = 0xFF0000
	Green = 0x00FF00
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.StringVar(&ChannelID, "c", "", "Channel ID")
	flag.StringVar(&Message, "m", "", "Message")
	flag.BoolVar(&Failure, "f", false, "Is a failure message?")
	flag.Parse()
}

func main() {
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Fatalf("Error creating Discord session: %v", err)
		return
	}

	err = dg.Open()
	if err != nil {
		log.Fatalf("Error opening connection: %v", err)
		return
	}

	defer dg.Close()

	embed := &discordgo.MessageEmbed{
		Title:       "Success",
		Description: Message,
		Color:       Green,
	}

	if Failure {
		embed = &discordgo.MessageEmbed{
			Title:       "Failure",
			Description: Message,
			Color:       Red,
		}
	}

	_, err = dg.ChannelMessageSendEmbed(ChannelID, embed)
	if err != nil {
		log.Fatalf("Error sending message: %v", err)
		return
	}

	fmt.Println("Message sent successfully!")
}
