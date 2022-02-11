package main

import (
	"fmt"
	"net/http"
	"github.com/bwmarrin/discordgo"
)

var Token string

const CAT_API_ENDPOINT = "https://api.thecatapi.com"

const GET_RANDOM_CAT = "/v1/images/search"

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!meowme" {
		response, err := http.Get(CAT_API_ENDPOINT + GET_RANDOM_CAT)
		if err != nil {
			fmt.Println("ERROR SENDING REQUEST TO CAT API: ", err)
		}
		defer response.Body.Close()

		if response.StatusCode == 200 {
			_, err = s.ChannelFileSend(m.ChannelID, "A cat wondered by...", response.Body)
			if err != nil {
				fmt.Println("ERROR SENDING FILE FROM RESPONSE: ", err)
			}
		} else {
			fmt.Println("ERROR RESPONSE FROM CAT API")
		}
	}
}