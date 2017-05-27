package main

// func GetClinic(cmd *flag.FlagSet) {
// 	jwt := cmd.String("jwt", "", "Admin's JWT. (Required)")
// 	id := cmd.String("id", "", "Clinic's Id. (Required)")

// 	cmd.Parse(os.Args[2:])

// 	if cmd.Parsed() {
// 		// Required Flags
// 		if *jwt == "" || *id == "" {
// 			cmd.PrintDefaults()
// 			os.Exit(1)
// 		}

// 		Admin, err := admin.New()
// 		if err != nil {
// 			panic(err)
// 		}

// 		clinic, err = Admin.GetClinic(*jwt, *id)
// 		admin.CheckErr(err)
// 		printClinic(Clinic)
// 	}
// }

// func printClinic(clinic admin.Clinic) {
// 	fmt.Println("\tName: ", clinic.Name)
// 	fmt.Println("\tAddress1: ", clinic.Address1)
// 	fmt.Println("\tCreated By: ", clinic.CreatedBy)
// }
