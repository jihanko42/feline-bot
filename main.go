package main

import (
	"fmt"
	"encoding/json"
	"flag"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"github.com/bwmarrin/discordgo" 
)

var Token string

const CAT_API_ENDPOINT = "https://api.thecatapi.com"

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	discord, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("ERROR ESTABLISHING DISCORD SESSION: ", err)
		return
	}
	
}