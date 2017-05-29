package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/oren/doc-api"
	"github.com/oren/doc-api/bolt"
	"github.com/oren/doc-api/config"
)

func AddClinic(cmd *flag.FlagSet) {
	configPath := cmd.String("config", "", "Config file (Optional)")
	adminJWT := cmd.String("jwt", "", "Admin's JWT. (Required)")
	clinicName := cmd.String("name", "", "Clinic's name. (Required)")
	clinicAddress1 := cmd.String("address1", "", "Clinic's address. (Required)")

	cmd.Parse(os.Args[2:])

	if !cmd.Parsed() {
		return
	}

	// Required Flags
	if *adminJWT == "" || *clinicName == "" || *clinicAddress1 == "" {
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

	adminService := &bolt.AdminService{Store: store}

	var claims *admin.MyCustomClaims
	claims, err = adminService.Authenticate(*adminJWT)
	if err != nil {
		log.Fatal(err)
	}

	clinic := &admin.Clinic{
		Name:     *clinicName,
		Address1: *clinicAddress1,
	}

	err = adminService.AddClinic(clinic, claims.Email)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Clinic was added")
}
