package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/oren/doc-api"
)

func ListAdmins(cmd *flag.FlagSet) {
	cmd.Parse(os.Args[2:])

	if cmd.Parsed() {
		Admin, err := admin.New()

		if err != nil {
			panic(err)
		}

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
