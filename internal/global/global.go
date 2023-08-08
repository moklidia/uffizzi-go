package global

import (
	"UffizziCloud/uffizzi-go/internal/config/settings"
	"log"
)

var (
	Settings *settings.Settings
)

func init() {
	var err error

	// Read and parse setting file and set package level Settings var
	Settings, err = settings.Read("configs/settings.yml")
	if err != nil {
		log.Fatal(err)
	}
}
