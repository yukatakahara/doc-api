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
	name := addCommand.String("name", "", "Admin's name. (Required)")

	listCommand := flag.NewFlagSet("list-admins", flag.ExitOnError)

	loginAdminCommand := flag.NewFlagSet("login-admin", flag.ExitOnError)
	email2 := loginAdminCommand.String("email", "", "Admin's email. (Required)")
	password2 := loginAdminCommand.String("password", "", "Admin's password. (Required)")

	// os.Arg[1] will be the subcommand
	if len(os.Args) < 2 {
		fmt.Println("add-admin or list-admins or login-admin subcommand is required")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add-admin":
		addCommand.Parse(os.Args[2:])
	case "list-admins":
		listCommand.Parse(os.Args[2:])
	case "login-admin":
		loginAdminCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if addCommand.Parsed() {
		// Required Flags
		if *email == "" || *password == "" || *name == "" {
			addCommand.PrintDefaults()
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

		results, err := Admin.All()
		admin.CheckErr(err)

		// data, err := json.Marshal(results)

		// if err != nil {
		// 	fmt.Errorf("encode response: %v", err)
		// }
		// os.Stdout.Write(data)

		PrintAdmins(results)
	}

	if loginAdminCommand.Parsed() {
		// Required Flags
		if *email2 == "" || *password2 == "" {
			addCommand.PrintDefaults()
			os.Exit(1)
		}

		Admin, err := admin.New()
		admin.CheckErr(err)

		Admin.Email = *email2

		err = Admin.Login(*password2)
		admin.CheckErr(err)

		fmt.Println("Admin exist")
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
