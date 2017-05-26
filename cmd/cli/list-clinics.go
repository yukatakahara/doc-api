package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/oren/doc-api"
)

func ListClinics(cmd *flag.FlagSet) {
	cmd.Parse(os.Args[2:])

	if cmd.Parsed() {
		Admin, err := admin.New()

		if err != nil {
			panic(err)
		}

		results, err := Admin.AllClinics()
		admin.CheckErr(err)
		printClinics(results)
	}
}

func printClinics(as []admin.Clinic) {
	fmt.Println("\n==== All clinics ====")

	for _, a := range as {
		fmt.Println("\tName: ", a.Name)
		fmt.Println("\tAddress1: ", a.Address1)
		fmt.Println("\tCreated By: ", a.CreatedBy)

	}
}
