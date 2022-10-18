package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	Token     string
	BotPrefix string

	config *configStruct
)

type configStruct struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
}

func ReadConfigFile() error {
	fmt.Println("reading config file...")

	// reading the config file for Token
	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		fmt.Println("error reading config file,", err)
		return err
	}
	// display the config file
	fmt.Println(string(file))

	// unmarshal the config file
	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	Token = config.Token
	BotPrefix = config.BotPrefix

	return nil
}
