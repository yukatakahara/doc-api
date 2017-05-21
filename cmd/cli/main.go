package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/cayleygraph/cayley/quad"
	"github.com/oren/doc-api"
)

func main() {
	addCommand := flag.NewFlagSet("add-admin", flag.ExitOnError)
	email := addCommand.String("email", "", "Admin's email. (Required)")
	password := addCommand.String("password", "", "Admin's password. (Required)")
	name := addCommand.String("name", "", "Admin's name. (Required)")

	listCommand := flag.NewFlagSet("list-admins", flag.ExitOnError)
	listClinics := flag.NewFlagSet("list-clinics", flag.ExitOnError)
	listQuads := flag.NewFlagSet("list-quads", flag.ExitOnError)

	loginAdminCommand := flag.NewFlagSet("login-admin", flag.ExitOnError)
	email2 := loginAdminCommand.String("email", "", "Admin's email. (Required)")
	password2 := loginAdminCommand.String("password", "", "Admin's password. (Required)")

	addClinic := flag.NewFlagSet("add-clinic", flag.ExitOnError)
	adminJWT := addClinic.String("jwt", "", "Admin's JWT. (Required)")
	clinicName := addClinic.String("name", "", "Clinic's name. (Required)")
	clinicAddress1 := addClinic.String("address1", "", "Clinic's address. (Required)")

	// os.Arg[1] will be the subcommand
	if len(os.Args) < 2 {
		fmt.Println("A subcommand is required")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add-admin":
		addCommand.Parse(os.Args[2:])
	case "list-admins":
		listCommand.Parse(os.Args[2:])
	case "login-admin":
		loginAdminCommand.Parse(os.Args[2:])
	case "add-clinic":
		addClinic.Parse(os.Args[2:])
	case "list-clinics":
		listClinics.Parse(os.Args[2:])
	case "list-quads":
		listQuads.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		fmt.Println("Command not found")
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
	}

	if listCommand.Parsed() {
		Admin, err := admin.New()

		if err != nil {
			panic(err)
		}

		results, err := Admin.All()
		admin.CheckErr(err)
		PrintAdmins(results)
	}

	if listClinics.Parsed() {
		Admin, err := admin.New()

		if err != nil {
			panic(err)
		}

		results, err := Admin.AllClinics()
		admin.CheckErr(err)
		PrintClinics(results)
	}

	if listQuads.Parsed() {
		Admin, err := admin.New()

		if err != nil {
			panic(err)
		}

		var quads []quad.Quad
		quads, err = Admin.AllQuads()
		admin.CheckErr(err)
		printQuads(quads)
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

		var jwt string
		jwt, err = Admin.Login(*password2)
		admin.CheckErr(err)

		fmt.Println("Admin exist. jwt:", jwt)
	}

	if addClinic.Parsed() {
		// Required Flags
		if *adminJWT == "" || *clinicName == "" || *clinicAddress1 == "" {
			addClinic.PrintDefaults()
			os.Exit(1)
		}

		Admin, err := admin.New()
		if err != nil {
			panic(err)
		}

		clinic := &admin.Clinic{
			Name:     *clinicName,
			Address1: *clinicAddress1,
		}
		err = Admin.AddClinic(clinic, *adminJWT)
		admin.CheckErr(err)
	}
}

func printQuads(quads []quad.Quad) {
	fmt.Println("\n==== All quads ====")

	for _, q := range quads {
		fmt.Println("quad", q)
	}
}

func PrintAdmins(as []admin.Admin) {
	fmt.Println("\n==== All admins ====")

	for _, a := range as {
		fmt.Println("\tEmail: ", a.Email)
		fmt.Println("\tHashedPassword: ", a.HashedPassword)
	}
}

func PrintClinics(as []admin.Clinic) {
	fmt.Println("\n==== All clinics ====")

	for _, a := range as {
		fmt.Println("\tName: ", a.Name)
		fmt.Println("\tAddress1: ", a.Address1)
		fmt.Println("\tCreated By: ", a.CreatedBy)

	}
}
