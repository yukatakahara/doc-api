package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	Users  []string
	Groups []string
}

func ReadConf() {
	file, _ := os.Open("conf.json")
	fmt.Println("file", file)
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err = decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("Error reading the configuration file:", err)
	}
	fmt.Println(configuration.Users) // output: [UserA, UserB]
}

// config.File(&conf, configPathFlagWhithSaneDefault).
