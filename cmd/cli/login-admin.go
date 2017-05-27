package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/oren/doc-api"
)

func LoginAdmin(cmd *flag.FlagSet) {
	email := cmd.String("email", "", "Admin's email. (Required)")
	password := cmd.String("password", "", "Admin's password. (Required)")

	cmd.Parse(os.Args[2:])

	if cmd.Parsed() {
		if *email == "" || *password == "" {
			cmd.PrintDefaults()
			os.Exit(1)
		}

		Admin, err := admin.New()
		admin.CheckErr(err)

		Admin.Email = *email

		var jwt string
		jwt, err = Admin.Login(*password)
		admin.CheckErr(err)

		fmt.Println("Admin exist. jwt:", jwt)
	}
}
