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
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Error opening the configuration file: ", err)
	}

	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err = decoder.Decode(&configuration)
	if err != nil {
		log.Fatal("Error decoding the configuration file: ", err)
	}

	return configuration
}
