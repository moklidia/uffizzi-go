package config

import (
	"encoding/json"
	"fmt"
	"log"

	configFile "UffizziCloud/uffizzi-go/internal/config_file"
	global "UffizziCloud/uffizzi-go/internal/global"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "The uffizzi config command lets you configure this command-line application",
	Long: `The uffizzi config command lets you configure this command-line application.
	If COMMAND is not specified, uffizzi config launches an interactive set up
	guide.

	For more information on configuration options, see:
	https://github.com/UffizziCloud/uffizzi_cli`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 && args[0] == "help" {
			cmd.Help()
			return
		}
		configure()
	},
}

func init() {
	ConfigCmd.AddCommand(listCmd)
	ConfigCmd.AddCommand(getValueCmd)
	ConfigCmd.AddCommand(setValueCmd)
	ConfigCmd.AddCommand(unsetValueCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all options and their values from the config file.",
	Run: func(cmd *cobra.Command, args []string) {
		list()
	},
}

var getValueCmd = &cobra.Command{
	Use:   "get-value",
	Short: "Displays the value of the specified option",
	Run: func(cmd *cobra.Command, args []string) {
		getValue(args[0])
	},
}

var setValueCmd = &cobra.Command{
	Use:   "set-value",
	Short: "Sets the value of the specified option.",
	Run: func(cmd *cobra.Command, args []string) {
		setValue(args[0], args[1])
	},
}

var unsetValueCmd = &cobra.Command{
	Use:   "unset-value",
	Short: "Deletes the value of the specified option.",
	Run: func(cmd *cobra.Command, args []string) {
		unsetValue(args[0])
	},
}

func configure() error {
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

func list() {
	config := configFile.ReadConfig()
	configOutput, err := json.MarshalIndent(config, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(configOutput))
}

func getValue(option string) {
	value := configFile.ReadOption(option)

	if value != nil {
		fmt.Println(value)
		return
	}
	fmt.Println(fmt.Sprintf("No such option or no value set: %s", option))
}

func setValue(option, value string) {
	configFile.SetOption(option, value)

	fmt.Println(fmt.Sprintf("The option '%s' was added with value '%s'", option, value))
}

func unsetValue(option string) {
	configFile.UnsetOption(option)

	fmt.Println(fmt.Sprintf("The option '%s' was removed", option))
}
