package main

import (
	"flag"
	"os"

	"github.com/oren/doc-api"
)

func AddAdmin(cmd *flag.FlagSet) {
	email := cmd.String("email", "", "Admin's email. (Required)")
	password := cmd.String("password", "", "Admin's password. (Required)")
	name := cmd.String("name", "", "Admin's name. (Required)")

	cmd.Parse(os.Args[2:])

	if cmd.Parsed() {
		// Required Flags
		if *email == "" || *password == "" || *name == "" {
			cmd.PrintDefaults()
			os.Exit(1)
		}

		Admin, err := admin.New()
		if err != nil {
			panic(err)
		}

		Admin.Email = *email
		Admin.Name = *name

		err = Admin.Create(*password)
		admin.CheckErr(err)
	}
}
