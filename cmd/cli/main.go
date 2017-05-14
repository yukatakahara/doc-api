package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/oren/doc-api"
)

func main() {
	addCommand := flag.NewFlagSet("add-admin", flag.ExitOnError)
	email := addCommand.String("email", "", "Admin's email. (Required)")
	password := addCommand.String("password", "", "Admin's password. (Required)")

	listCommand := flag.NewFlagSet("list-admins", flag.ExitOnError)

	// os.Arg[1] will be the subcommand
	if len(os.Args) < 2 {
		fmt.Println("add-admin or list-admins subcommand is required")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add-admin":
		addCommand.Parse(os.Args[2:])
	case "list-admins":
		listCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if addCommand.Parsed() {
		// Required Flags
		if *email == "" || *password == "" {
			addCommand.PrintDefaults()
			os.Exit(1)
		}

		Admin, err := admin.New()
		if err != nil {
			panic(err)
		}

		Admin.Email = *email
		Admin.Password = *password

		_ = Admin.Create()
		// data, err := json.Marshal(results)

		// if err != nil {
		// 	fmt.Errorf("encode response: %v", err)
		// }

		// os.Stdout.Write(data)
	}

	if listCommand.Parsed() {
		Admin, err := admin.New()

		if err != nil {
			panic(err)
		}

		results := Admin.All()

		// data, err := json.Marshal(results)

		// if err != nil {
		// 	fmt.Errorf("encode response: %v", err)
		// }
		// os.Stdout.Write(data)

		PrintAdmins(results)
	}
}

func PrintAdmins(as []admin.Admin) {
	fmt.Println("\n==== All admins ====")

	for _, a := range as {
		fmt.Println("ID: ", a.ID)
		fmt.Println("\tEmail: ", a.Email)
		fmt.Println("\tHashedPassword: ", a.HashedPassword)
	}
}
