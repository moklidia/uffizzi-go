package config

import (
	"encoding/json"
	"fmt"
	"log"

	configFile "UffizziCloud/uffizzi-cli/internal/config_file"

	"github.com/manifoldco/promptui"
)

func Config() error {
	fmt.Println(`Configure the default properties that will be used to authenticate with your
	  Uffizzi API service and manage previews.\n`)
	promt := promptui.Prompt{
		Label:   "server",
		Default: "https://app.uffizzi.com",
	}

	server, err := promt.Run()

	if err != nil {
		fmt.Printf("Config failed %v\n", err)
		return err
	}

	configFile.SetOption("server", server)

	fmt.Printf("The server is set as %q\n", server)
	fmt.Printf("To login, run: uffizzi login\n")
	return nil
}

func List() {
	config := configFile.ReadConfig()
	configOutput, err := json.MarshalIndent(config, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(configOutput))
}

func GetValue(option string) {
	fmt.Println(configFile.ReadOption(option))
}
