// A stand-alone HTTP server
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math"
	"net/http"

	"github.com/cayleygraph/cayley"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/oren/doc-api/bolt"
	"github.com/oren/doc-api/config"
)

// Message struct represents the JSON document which the API sends when something wrong will happen.
// Type: If this is an error message, "Error" or "Message" can be the two values.
// UserMessage: This is the string which the developer can show to the user.
// DeveloperMessage: This is the technical message.
// DocumentationLink: If the message type is error, this will be the link to corresponding documentation.
type Message struct {
	Type              string `json:"messageType"`
	UserMessage       string `json:"userMessage"`
	DeveloperMessage  string `json:"developerMessage"`
	DocumentationLink string `json:"documentationLink"`
}

type MyCustomClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// TODO: store should not be global
var store *cayley.Handle
var adminService *bolt.AdminService

// var adminService admin.AdminService

func init() {
	// initialize the db
	configPath := flag.String("config", "", "Path to config.json")

	flag.Parse()

	if *configPath == "" {
		*configPath = config.GetPathOfConfig()
	}

	configuration := config.ReadConf(*configPath)

	var err error
	store, err = bolt.Open(configuration.DbPath)
	if err != nil {
		log.Fatal(err)
	}

	// Create admin service
	adminService = &bolt.AdminService{Store: store}

	// Create admin service
	// adminService = &bolt.AdminService{Store: store}

	// TODO: When do i close the db?
	// defer db.Close()

	// TODO: What about interface?
	// https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1

	// POST /signup - create jwt
	http.HandleFunc("/adminlogin", adminLogin)
	// GET /clinics - return all clinics
	// GET /clinics/1 - return a clinic
	// POST /clinics - create new clinic
	// PUT /clinics/1 - update a clinic
	// DELETE /clinics/1 - delete a clinic
	// http.HandleFunc("/clinics", clinicsHandler)
	// http.HandleFunc("/doctors", DoctorsHandler)
	// http.HandleFunc("/login", memberLogin)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func main() {
}

type User struct {
	Email string `json:"email"`
	JWT   string `json:"jwt"`
}

func ServerError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

// ReturnMessageJSON is a wrapper which will send JSON document of type Message, it takes the following arguments
// messageType: Error or Information, denotes if the message is just FYI or an error
// userMessage: Message in terms of user
// devMessage: Message in terms of Developer
func ReturnMessageJSON(w http.ResponseWriter, messageType, userMessage, devMessage string) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers:", "Origin, Content-Type, X-Auth-Token, Authorization")
	w.Header().Set("Content-Type", "application/json")

	message := Message{Type: messageType, UserMessage: userMessage, DeveloperMessage: devMessage}

	if messageType == "Information" {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}

	err := json.NewEncoder(w).Encode(message)
	fmt.Println("after Encoding")

	if err != nil {
		fmt.Println("error in Encoding", err)
		// panic(err)
	}

	//TODO: why is the message is not returned?
	return
}

// haversin(θ) function
func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

// Distance function returns the distance (in meters) between two points of
//     a given longitude and latitude relatively accurately (using a spherical
//     approximation of the Earth) through the Haversin Distance Formula for
//     great arc distance on a sphere with accuracy for small distances
//
// point coordinates are supplied in degrees and converted into rad. in the func
//
// distance returned is METERS!!!!!!
// http://en.wikipedia.org/wiki/Haversine_formula
func Distance(lat1, lon1, lat2, lon2 float64) float64 {
	// convert to radians
	// must cast radius as float to multiply later
	var la1, lo1, la2, lo2, r float64
	la1 = lat1 * math.Pi / 180
	lo1 = lon1 * math.Pi / 180
	la2 = lat2 * math.Pi / 180
	lo2 = lon2 * math.Pi / 180

	r = 6378100 // Earth radius in METERS

	// calculate
	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)

	return 2 * r * math.Asin(math.Sqrt(h))
}

// our 'database' - for now it will be in memory but later on we'll save it in postgres
// slice that each of it's elements is the Doctor struct
// var doctors = []Doctor{
// 	{1, "Novena Medical Center Singapore", "dan@gmail.com", "123", "cats", "10 Sinaran Dr, Singapore 307506", "1.328829", "103.844859", "43.00"},
// 	{2, "Central 24-HR Clinic (Clementi)", "laura@gmail.com", "143", "dogs", "Blk 450 Clementi Ave 3 , #01-291, Singapore 120450", "1.320201", "103.764449", "44.00"},
// 	{4, "Central 24-HR Clinic (Hougang)", "laura@gmail.com", "143", "dogs", "681 Hougang Avenue 8 #01-831, Singapore 530681", "1.380640", "103.884989", "44.00"},
// 	{5, "Central 24-HR Clinic (Bedok)", "laura@gmail.com", "143", "dogs", "Blk 219 Bedok Central #01-124, Singapore 460219", "1.332237", "103.932853", "45.50"},
// 	{6, "SingHealth Polyclinics", "laura@gmail.com", "143", "dogs", " 580 Stirling Road, Singapore 148958", "1.305447", "103.802419", "49.00"},
// 	{7, "International Medical Clinic", "laura@gmail.com", "143", "dogs", "1 Orchard Boulevard, #14-06 Camden Medical Centre, Singapore 248649", "1.310574", "103.824854", "50.00"},
// 	{8, "Central 24-HR Clinic (Jurong West)", "laura@gmail.com", "143", "dogs", "492 Jurong West Street 41, #01-54, Singapore 640492", "1.356569", "103.724055", "50.00"},
// 	{9, "Central 24-HR Clinic (Pasir Ris)", "laura@gmail.com", "143", "dogs", "446 Pasir Ris Drive 6 #01-122, Singapore 510446", "1.376433", "103.957140", "50.00"},
// 	{10, "Toh Guan Family Clinic", "laura@gmail.com", "143", "dogs", "267A Toh Guan Rd, Singapore 601267", "1.349597", "103.746308", "50.00"},
// 	{11, "Healthway Medical Clinic", "laura@gmail.com", "143", "dogs", "267 Compassvale Link #01-04, Singapore 544267", "1.391390", "103.898317", "50.00"},
// 	{12, "Thomson 24-Hour Family Clinic", "laura@gmail.com", "143", "dogs", "339 Thomson Rd, Singapore 307677", "1.332861", "103.841544", "50.00"},
// 	{13, "Shenton Family Medical Clinic", "laura@gmail.com", "143", "dogs", "201D Tampines Street 21, #01-1137, Singapore 524201", "1.358216", "103.952443", "50.00"},
// 	{14, "The Travel Clinic", "laura@gmail.com", "143", "dogs", "Level 4,17, Third Hospital Drive, Diabetes & Metabolism Centre, Singapore 168752", "1.287128", "103.836738", "50.00"},
// 	{15, "Dayspring Medical Clinic - Pasir Ris", "laura@gmail.com", "143", "dogs", "1 Pasir Ris Central Street 3, #05-09 White Sands, White Sands, Singapore 518457", "1.372844", "103.949525", "50.00"},
// 	{16, "Parkway Shenton Pte Ltd", "laura@gmail.com", "143", "dogs", "20 Bendemeer Rd, Singapore 33991", "1.314670", "103.862142", "50.00"},
// }
