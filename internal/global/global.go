package global

import (
	"UffizziCloud/uffizzi-go/internal/config/settings"
	"log"
	"os"
)

var (
	Settings *settings.Settings
)

func init() {
	var err error
	env := os.Getenv("ENV")

	// Read and parse setting file and set package level Settings var
	Settings, err = settings.Read(env, "configs/settings.yml")
	if err != nil {
		log.Fatal(err)
	}
}
