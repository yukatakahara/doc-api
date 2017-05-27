package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/oren/doc-api"
)

func AddClinic(cmd *flag.FlagSet) {
	adminJWT := cmd.String("jwt", "", "Admin's JWT. (Required)")
	clinicName := cmd.String("name", "", "Clinic's name. (Required)")
	clinicAddress1 := cmd.String("address1", "", "Clinic's address. (Required)")

	cmd.Parse(os.Args[2:])

	// Required Flags
	if *adminJWT == "" || *clinicName == "" || *clinicAddress1 == "" {
		cmd.PrintDefaults()
		os.Exit(1)
	}

	Admin, err := admin.New()
	if err != nil {
		panic(err)
	}

	var claims *admin.MyCustomClaims
	claims, err = Admin.Authenticate(*adminJWT)

	if err != nil {
		log.Fatal(err)
	}

	clinic := &admin.Clinic{
		Name:     *clinicName,
		Address1: *clinicAddress1,
	}

	err = Admin.AddClinic(clinic, claims.Email)
	admin.CheckErr(err)
	fmt.Println("Clinic was added")
}
