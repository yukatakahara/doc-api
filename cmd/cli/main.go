package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/oren/doc-api"
)

func main() {
	addCommand := flag.NewFlagSet("add-admin", flag.ExitOnError)
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

	deleteClinicCmd := flag.NewFlagSet("delete-clinic", flag.ExitOnError)

	// os.Arg[1] will be the subcommand
	if len(os.Args) < 2 {
		fmt.Println("A subcommand is required")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add-admin":
		AddAdmin(addCommand)
	case "list-admins":
		ListAdmins(listCommand)
	case "login-admin":
		loginAdminCommand.Parse(os.Args[2:])
	case "add-clinic":
		addClinic.Parse(os.Args[2:])
	case "list-clinics":
		ListClinics(listClinics)
	case "list-quads":
		ListQuads(listQuads)
	case "delete-clinic":
		DeleteClinic(deleteClinicCmd)
	default:
		flag.PrintDefaults()
		fmt.Println("Command not found")
		os.Exit(1)
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
