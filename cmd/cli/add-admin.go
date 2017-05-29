package main

import (
	"flag"
	"log"
	"os"

	"github.com/oren/doc-api"
	"github.com/oren/doc-api/bolt"
	"github.com/oren/doc-api/config"
)

func AddAdmin(cmd *flag.FlagSet) {
	configPath := cmd.String("config", "", "Config file (Optional)")
	email := cmd.String("email", "", "Admin's email. (Required)")
	password := cmd.String("password", "", "Admin's password. (Required)")
	name := cmd.String("name", "", "Admin's name. (Required)")

	cmd.Parse(os.Args[2:])

	if !cmd.Parsed() {
		return
	}

	// Required Flags
	if *email == "" || *password == "" || *name == "" {
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
	newAdmin := &admin.Admin{
		Name:  *name,
		Email: *email,
	}

	err = adminService.CreateAdmin(newAdmin, *password)

	if err != nil {
		log.Fatal(err)
	}
}
