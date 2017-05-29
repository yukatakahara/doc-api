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

func ListClinics(cmd *flag.FlagSet) {
	configPath := cmd.String("config", "", "Config file (Optional)")
	cmd.Parse(os.Args[2:])

	if !cmd.Parsed() {
		return
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
	var results []admin.Clinic
	results, err = adminService.AllClinics()

	if err != nil {
		log.Fatal(err)
	}

	admin.CheckErr(err)
	printClinics(results)
}

func printClinics(as []admin.Clinic) {
	fmt.Println("\n==== All clinics ====")

	for _, a := range as {
		fmt.Println("\tName: ", a.Name)
		fmt.Println("\tAddress1: ", a.Address1)
		fmt.Println("\tCreated By: ", a.CreatedBy)

	}
}
