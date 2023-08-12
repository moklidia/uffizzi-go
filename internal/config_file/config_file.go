package configfile

import (
	"UffizziCloud/uffizzi-go/internal/global"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/afero"
	"github.com/spf13/viper"
)

var (
	FS  afero.Fs     = getOs()
	AFS *afero.Afero = &afero.Afero{Fs: FS}
)

type AccountConfig struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Config struct {
	Account   AccountConfig `json:"account"`
	Host      string        `json:"host"`
	Project   string        `json:"project"`
	cookie    string        `json:"cookie"`
	oidcToken string        `json:"oidc_token"`
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

	configDir := fmt.Sprintf("%s/%s", homeDir, global.Settings.ConfigFile.Dir)
	configPath := fmt.Sprintf("%s/%s.%s", configDir, global.Settings.ConfigFile.Name, global.Settings.ConfigFile.Ext)
	configExists := true
	if _, err := AFS.Stat(configPath); err != nil {
		configExists = false
		createConfigFile(configDir, configPath)
	}

	v := viper.New()
	v.SetConfigName(global.Settings.ConfigFile.Name)
	v.SetConfigType(global.Settings.ConfigFile.Ext)
	v.AddConfigPath(configDir)
	v.AllowEmptyEnv(true)
	v.SetFs((AFS))
	if !configExists {
		v.WriteConfig()
	}
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
	AFS.MkdirAll(dir, 0700)
	var config Config

	fileContents, err := json.Marshal(&config)
	if err != nil {
		log.Fatal(err)
	}
	AFS.WriteFile(path, fileContents, 0644)

	fmt.Println(fmt.Sprintf("Config file created successfully at path %s", path))

	return nil
}

func getOs() afero.Fs {
	env := os.Getenv("ENV")
	if env == "test" {
		return afero.NewMemMapFs()
	}

	return afero.NewOsFs()
}
