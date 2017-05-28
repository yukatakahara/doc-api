package config

import (
	"encoding/json"
	"log"
	"os"
)

type Configuration struct {
	WebPort string
	DbPath  string
}

func ReadConf(path string) Configuration {
	file, _ := os.Open(path)
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)

	if err != nil {
		log.Fatal("Error reading the configuration file: ", err)
	}

	return configuration
}

//$ web --config /path/to/config.json
//$ cli --config /path/to/config.json
// $HOME/.config/$PROJECT/config.json
