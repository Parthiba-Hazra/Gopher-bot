package bot

import (
	"fmt"

	"github.com/Parthiba-Hazra/Gopher-bot/gopher-discord/config"
	"github.com/bwmarrin/discordgo"
)

var BotID string
var gofi_bot *discordgo.Session

const Kute_go_APIURL = "https://kutego-api-xxxxx-ew.a.run.app"

func Start() {

	// Create a new Discord session using the bot token.
	bot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println("error creating new discord session,", err)
		return
	}

	u, err := bot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	BotID = u.ID

	bot.AddHandler(messageHandler)

	// we only care about receiving message events.
	bot.Identify.Intents = discordgo.IntentsGuildMessages

	err = bot.Open()

	if err != nil {
		fmt.Println("error opening the connection,", err)
		return
	}
	fmt.Println("Bot is running! to exit press CTRL-C")

}
