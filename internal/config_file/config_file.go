package configfile

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Account struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"account"`
	Host      string `json:"host"`
	Project   string `json:"project"`
	cookie    string `json:"cookie"`
	oidcToken string `json:"oidc_token"`
}

type viperConfigReader struct {
	viper *viper.Viper
}

var ConfigReader *viperConfigReader

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	configDir := fmt.Sprintf("%s/uffizzi", homeDir)
	configPath := fmt.Sprintf("%s/config.json", configDir)
	if _, err := os.Stat(configPath); err != nil {
		createConfigFile(configDir, configPath)
	}

	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("json")
	v.AddConfigPath(configDir)
	v.SetDefault("host", "https://app.uffizzi.com")
	v.AllowEmptyEnv(true)
	v.ReadInConfig()
	err = v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	ConfigReader = &viperConfigReader{
		viper: v,
	}
}

func ReadConfig() Config {
	var config Config
	ConfigReader.viper.Unmarshal(&config)

	return config
}

func ReadOption(option string) interface{} {
	return ConfigReader.viper.Get(option)
}

func SetOption(option, value string) {
	ConfigReader.viper.Set(option, value)
	err := ConfigReader.viper.WriteConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func UnsetOption(option string) {
	ConfigReader.viper.Set(option, "")
	err := ConfigReader.viper.WriteConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func createConfigFile(dir, path string) error {
	os.MkdirAll(dir, 0700)
	_, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("Config file created successfully at path %s", path))

	return nil
}
