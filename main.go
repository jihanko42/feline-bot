package main

import (
	"fmt"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"github.com/bwmarrin/discordgo" 
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	disc, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("ERROR ESTABLISHING DISCORD SESSION: ", err)
		return
	}
	
	disc.AddHandler(messageCreate)

	disc.Identify.Intents = discordgo.IntentsGuildMessages

	err = disc.Open()
	if err != nil {
		fmt.Println("ERROR OPENING CONNECTION TO SERVER: ", err)
		return
	}

	fmt.Println("Bot is now running")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc


	disc.Close()
}