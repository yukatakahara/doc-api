package main

import (
	"flag"
	"os"

	"github.com/oren/doc-api"
)

func deleteClinic2(cmd *flag.FlagSet) {
	jwt := cmd.String("jwt", "", "Admin's JWT. (Required)")
	id := cmd.String("id", "", "Clinic's Id. (Required)")

	cmd.Parse(os.Args[2:])

	if cmd.Parsed() {
		// Required Flags
		if *jwt == "" || *id == "" {
			cmd.PrintDefaults()
			os.Exit(1)
		}

		Admin, err := admin.New()
		if err != nil {
			panic(err)
		}

		err = Admin.DeleteClinic(*jwt, *id)
		admin.CheckErr(err)
	}
}
