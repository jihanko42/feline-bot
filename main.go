package main

import (
	"fmt"
	"feline-bot/bot"
	"feline-bot/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println("CONFIG ERROR", err.Error())
		return
	}

	bot.Start()

	return
}