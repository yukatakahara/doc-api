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

func ListAdmins(cmd *flag.FlagSet) {
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
	var results []admin.Admin
	results, err = adminService.All()

	if err != nil {
		log.Fatal(err)
	}

	printAdmins(results)
}

func printAdmins(as []admin.Admin) {
	fmt.Println("\n==== All admins ====")

	for _, a := range as {
		fmt.Println("\tEmail: ", a.Email)
		fmt.Println("\tHashedPassword: ", a.HashedPassword)
	}
}
