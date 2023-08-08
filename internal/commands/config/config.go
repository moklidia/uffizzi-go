package config

import (
	"encoding/json"
	"fmt"
	"log"

	configFile "UffizziCloud/uffizzi-go/internal/config_file"
	global "UffizziCloud/uffizzi-go/internal/global"

	"github.com/manifoldco/promptui"
)

func Config() error {
	fmt.Println(`Configure the default properties that will be used to authenticate with your
	  Uffizzi API service and manage previews.\n`)
	promt := promptui.Prompt{
		Label:   "host",
		Default: global.Settings.DefaultHost,
	}

	host, err := promt.Run()

	if err != nil {
		fmt.Printf("Config failed %v\n", err)
		return err
	}

	configFile.SetOption("host", host)

	fmt.Printf("The host is set as %q\n", host)
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
	value := configFile.ReadOption(option)
	if value != nil {
		fmt.Println(value)
	}
	fmt.Println(fmt.Sprintf("No such option or no value set: %s", option))
}

func SetValue(option, value string) {
	configFile.SetOption(option, value)

	fmt.Println(fmt.Sprintf("The option '%s' was added with value '%s'", option, value))
}

func UnsetValue(option string) {
	configFile.UnsetOption(option)

	fmt.Println(fmt.Sprintf("The option '%s' was removed", option))
}
