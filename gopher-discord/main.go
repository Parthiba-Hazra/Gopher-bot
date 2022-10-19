package main

import (
	"fmt"

	"github.com/Parthiba-Hazra/Gopher-bot/gopher-discord/bot"
	"github.com/Parthiba-Hazra/Gopher-bot/gopher-discord/config"
)

func main() {

	err := config.ReadConfigFile()

	if err != nil {
		fmt.Println("error reading config file,", err)
		return
	}

	bot.Start()

	// sc := make(chan os.Signal, 1)
	// signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	// <-sc

	<-make(chan struct{})
	// close down the discord session
	return
}
