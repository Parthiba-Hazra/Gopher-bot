package bot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Parthiba-Hazra/Gopher-bot/gopher-discord/config"
	"github.com/bwmarrin/discordgo"
)

var BotID string
var gofi_bot *discordgo.Session

type Gopher struct {
	Name string `json: "name"`
}

const Kute_go_APIURL = "http://localhost:8080"

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

	err = bot.Open()

	if err != nil {
		fmt.Println("error while opening the connection, ", err)
		return
	}
	fmt.Println("Bot is running... to exit press CTRL-c")

}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages on the channel created by the bot itself
	if m.Author.ID == BotID {
		return
	}

	if m.Content == "!gopher" {

		//Calling the KuteGo api and get the Dr Who Gopher
		response, err := http.Get(Kute_go_APIURL + "/gopher/" + "back-to-the-future-v2")
		if err != nil {
			fmt.Println(err)
		}
		defer response.Body.Close()

		if response.StatusCode == 200 {
			_, err = s.ChannelFileSend(m.ChannelID, "dr-who.png", response.Body)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Error! - Can't get the dr-who gopher :( ")
		}
	}

	if m.Content == "!random" {

		//Calling the KuteGoapi and get a random Gopher
		response, err := http.Get(Kute_go_APIURL + "/gopher/random/")
		if err != nil {
			fmt.Println(err)
		}
		defer response.Body.Close()

		if response.StatusCode == 200 {
			_, err = s.ChannelFileSend(m.ChannelID, "random-gopher.png", response.Body)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Error! - Can't get random Gophers :(")
		}
	}

	if m.Content == "!gophers" {

		//Calling the KuteGo api and get the list of all Gophers
		response, err := http.Get(Kute_go_APIURL + "/gophers/")
		if err != nil {
			fmt.Println(err)
		}
		defer response.Body.Close()

		if response.StatusCode == 200 {

			// convert the response to a []byte
			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				fmt.Println(err)
			}

			// keep needed informations from the JSON response in our array of Gopher.
			var data []Gopher
			err = json.Unmarshal(body, &data)
			if err != nil {
				fmt.Println(err)
			}

			// merge all the gophers name in a string
			var gophers strings.Builder
			for _, gopher := range data {
				gophers.WriteString(gopher.Name + "\n")
			}

			// Send the text message contain all Gophers
			_, err = s.ChannelMessageSend(m.ChannelID, gophers.String())
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Error! - Can't get list of Gophers :(")
		}
	}
}
