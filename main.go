package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

const CAT_API_ENDPOINT = "https://api.thecatapi.com"

const GET_RANDOM_CAT = "/v1/images/search"

var (
	Token string
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

	fmt.Println("FELINE-BOT IS NOW RUNNING: TRY !MEOWME IN YOUR DISCORD CHANNEL")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	disc.Close()
}

type CatResponse struct {
	Url string
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!meowme" {
		route := CAT_API_ENDPOINT + GET_RANDOM_CAT
		client := &http.Client{}
		req, err := http.NewRequest("GET", route, nil)
		if err != nil {
			fmt.Println("ERROR CREATING REQUEST: ", err)
		}

		req.Header.Set("x-api-key", "3a96ee5f-2244-40af-b33f-c07b7da8747e")
		fmt.Println("REQUEST: ", req)
		response, err := client.Do(req)
		if err != nil {
			fmt.Println("ERROR SENDING REQUEST TO CAT API: ", err)
		}
		defer response.Body.Close()

		if response.StatusCode == 200 {
			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				fmt.Println("ERROR GETTING FILE FROM RESPONSE: ", err)
			}

			var catImage []CatResponse
			err = json.Unmarshal(body, &catImage)
			fmt.Println(catImage)
			if err != nil {
				fmt.Println("ERROR PARSING RESPONSE: ", err)
			}

			_, err = s.ChannelMessageSend(m.ChannelID, catImage[0].Url)

			if err != nil {
				fmt.Println("ERROR SENDING IMAGE IN CHAT: ", err)
			}

		} else {
			fmt.Println("ERROR RESPONSE FROM CAT API")
		}
	}
}
