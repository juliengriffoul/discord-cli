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
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.StringVar(&ChannelID, "c", "", "Channel ID")
	flag.StringVar(&Message, "m", "", "Message")
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

	_, err = dg.ChannelMessageSend(ChannelID, Message)
	if err != nil {
		log.Fatalf("Error sending message: %v", err)
		return
	}

	fmt.Println("Message sent successfully!")
}
