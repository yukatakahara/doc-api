package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/oren/doc-api"
	"github.com/oren/doc-api/config"
)

func ListAdmins(cmd *flag.FlagSet) {
	configPath := cmd.String("config", "", "Config file (Optional)")
	cmd.Parse(os.Args[2:])

	if cmd.Parsed() {
		Admin, err := admin.New()

		if err != nil {
			panic(err)
		}

		if *configPath == "" {
			*configPath = "/tmp/config.json"
		}

		configuration := config.ReadConf(*configPath)
		fmt.Println(configuration)

		results, err := Admin.All()
		admin.CheckErr(err)
		printAdmins(results)
	}
}

func printAdmins(as []admin.Admin) {
	fmt.Println("\n==== All admins ====")

	for _, a := range as {
		fmt.Println("\tEmail: ", a.Email)
		fmt.Println("\tHashedPassword: ", a.HashedPassword)
	}
}
