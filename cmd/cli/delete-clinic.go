package main

import (
	"flag"
	"log"
	"os"

	"github.com/oren/doc-api/bolt"
	"github.com/oren/doc-api/config"
)

func DeleteClinic(cmd *flag.FlagSet) {
	configPath := cmd.String("config", "", "Config file (Optional)")
	adminJWT := cmd.String("jwt", "", "Admin's JWT. (Required)")
	id := cmd.String("id", "", "Clinic's Id. (Required)")

	cmd.Parse(os.Args[2:])

	if !cmd.Parsed() {
		return
	}

	// Required Flags
	if *adminJWT == "" || *id == "" {
		cmd.PrintDefaults()
		os.Exit(1)
	}

	if *configPath == "" {
		*configPath = config.GetPathOfConfig()
	}

	configuration := config.ReadConf(*configPath)

	store, err := bolt.Open(configuration.DbPath)
	if err != nil {
		log.Fatal(err)
	}

	// Create admin service
	adminService := &bolt.AdminService{Store: store}

	_, err = adminService.Authenticate(*adminJWT)
	if err != nil {
		log.Fatal(err)
	}

	err = adminService.DeleteClinic(*id)

	if err != nil {
		log.Fatal(err)
	}

	// TODO: is there a way to tell if it was deleted?
}
