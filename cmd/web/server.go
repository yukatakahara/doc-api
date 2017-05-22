// A stand-alone HTTP server
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/oren/doc-api"
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

var mySigningKey = []byte("secret")

type Doctor struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	EmployeeId string `json:"employee_id"`
	TeamName   string `json:"team_name"`
	Address    string `json:"adress"`
	Lat        string `json:"lat"`
	Lon        string `json:"lon"`
	Price      string `json:"price"`
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

func init() {
	// POST /signup - create jwt
	http.HandleFunc("/adminlogin", adminLogin)

	// GET /clinics - return all clinics
	// GET /clinics/1 - return a clinic
	// POST /clinics - create new clinic
	// PUT /clinics/1 - update a clinic
	// DELETE /clinics/1 - delete a clinic
	http.HandleFunc("/clinics", clinicsHandler)

	http.HandleFunc("/signup", GetTokenHandler)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/settings", GetSettingsHandler)
	http.HandleFunc("/doctors", DoctorsHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func main() {
}

type User struct {
	Email string `json:"email"`
	JWT   string `json:"jwt"`
}

func DoctorsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getDoctors(w, r)
	case "OPTIONS":
		getDoctors(w, r)
	default:
		http.Error(w, r.Method+" not allowed", http.StatusMethodNotAllowed)
	}
}

func getDoctors(w http.ResponseWriter, r *http.Request) {
	fmt.Println("0000000000000")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers:", "Origin, Content-Type, X-Auth-Token, Authorization")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		ReturnMessageJSON(w, "Information", "", "")
		return
	}

	lat := r.URL.Query()["lat"]
	lon := r.URL.Query()["lon"]

	if lat != nil && lon != nil {
		// order the list of doctors based on proximity
		fmt.Println("URL", r.URL)
	}

	Admin, err := admin.New()
	if err != nil {
		ReturnMessageJSON(w, "Error", "Authentication Failed", fmt.Sprintf("Error in admin login: %s", err))
		return
	}

	tokenHeader, found := r.Header["Authorization"]

	if !found {
		fmt.Println("not found token header")
		ReturnMessageJSON(w, "Error", "Authentication Failed", "Create Clinic - no auth token")
		return
	}

	jwt := tokenHeader[0]
	var claims *admin.MyCustomClaims
	claims, err = Admin.Authenticate(jwt[7:])
	fmt.Println(claims)

	if err != nil {
		fmt.Println("Error in Auth", err)
		ReturnMessageJSON(w, "Error", "Authentication Failed", "Create Clinic - error with token")
		return
	}

	var clinics []admin.Clinic
	clinics, err = Admin.AllClinics()
	if err != nil {
		fmt.Println("1111")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// then we encode it as JSON on the response
	err = json.NewEncoder(w).Encode(clinics)

	// And if encoding fails we log the error
	if err != nil {
		fmt.Println("2222")
		fmt.Errorf("encode response: %v", err)
	}
}

func sortDoctors(doctors []Doctor, lat string, lon string) []Doctor {

	return doctors
}

func ServerError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func clinicsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers:", "Origin, Content-Type, X-Auth-Token, Authorization")

	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		ReturnMessageJSON(w, "Information", "", "")
		return
	}

	if r.Method != "POST" {
		ReturnMessageJSON(w, "Error", "Page not available", "GetTokenHandler only accepts a POST")
		return
	}

	// authenticate admin
	// validate clinic input
	// create clinic in bolt
	Admin, err := admin.New()
	if err != nil {
		ReturnMessageJSON(w, "Error", "Authentication Failed", fmt.Sprintf("Error in admin login: %s", err))
		return
	}

	tokenHeader, found := r.Header["Authorization"]

	if !found {
		fmt.Println("not found token header")
		ReturnMessageJSON(w, "Error", "Authentication Failed", "Create Clinic - no auth token")
		return
	}

	jwt := tokenHeader[0]
	var claims *admin.MyCustomClaims
	claims, err = Admin.Authenticate(jwt[7:])

	if err != nil {
		fmt.Println("Error in Auth", err)
		ReturnMessageJSON(w, "Error", "Authentication Failed", "Create Clinic - error with token")
		return
	}

	newClinic := &admin.NewClinic{}

	if err := json.NewDecoder(r.Body).Decode(newClinic); err != nil {
		fmt.Println("1", err)
		ReturnMessageJSON(w, "Error", "Error with decoding of clinic", "Error with decoding of clinic")
		return
	}

	fmt.Println("newClinic", newClinic)

	if newClinic.Name == "" || newClinic.Address1 == "" {
		fmt.Println("2", err)
		ReturnMessageJSON(w, "Error", "Invalid Clinic fields", "Invalid Clinic fields")
		return
	}

	c := &admin.Clinic{
		Name:     newClinic.Name,
		Address1: newClinic.Address1,
	}
	err = Admin.AddClinic(c, claims.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(claims)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func adminLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		ReturnMessageJSON(w, "Information", "", "")
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers:", "Origin, Content-Type, X-Auth-Token")
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "POST" {
		ReturnMessageJSON(w, "Error", "Page not available", "GetTokenHandler only accepts a POST")
		return
	}

	a := &admin.EmailAndPassword{}

	if err := json.NewDecoder(r.Body).Decode(a); err != nil {
		fmt.Println("here", err)
		ServerError(w, err)
		return
	}

	if a.Email == "" || a.Password == "" {
		ReturnMessageJSON(w, "Error", "Invalid Email/Password", "Invalid Email or password in GetTokenHandler")
		return
	}

	Admin, err := admin.New()
	if err != nil {
		ReturnMessageJSON(w, "Error", "Authentication Failed", fmt.Sprintf("Error in admin login: %s", err))
		return
	}

	Admin.Email = a.Email

	var jwt string

	jwt, err = Admin.Login(a.Password)
	user := User{a.Email, jwt}
	js, err := json.Marshal(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(js)
}

//GetTokenHandler will get a token for the username and password
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		ReturnMessageJSON(w, "Information", "", "")
		return
	}

	if r.Method != "POST" {
		ReturnMessageJSON(w, "Error", "Page not available", "GetTokenHandler only accepts a POST")
		return
	}

	// r.ParseForm()
	username := "josh"
	password := "password123"
	// log.Println(r.Form)

	if username == "" || password == "" {
		ReturnMessageJSON(w, "Error", "Invalid Username/Password", "Invalid Username or password in GetTokenHandler")
		return
	}

	// if db.ValidUser(username, password) {
	if true {
		// Create the Claim which expires after EXPIRATION_HOURS hrs, default is 5.
		claims := MyCustomClaims{
			username,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 5).Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		/* Sign the token with our secret */
		tokenString, err := token.SignedString(mySigningKey)
		if err != nil {
			log.Println("Something went wrong with signing token")
			ReturnMessageJSON(w, "Error", "Authentication Failed", "Authentication Failed")

			return
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers:", "Origin, Content-Type, X-Auth-Token")
		w.Header().Set("Content-Type", "application/json")

		user := User{"josh@gmail.com", tokenString}

		js, err := json.Marshal(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(js)

		// w.Write([]byte(tokenString))
	} else {
		ReturnMessageJSON(w, "Error", "Authentication Failed", "Authentication Failed")
	}
}

//GetTokenHandler will get a token for the username and password
func GetTokenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		ReturnMessageJSON(w, "Information", "", "")
		return
	}

	if r.Method != "POST" {
		ReturnMessageJSON(w, "Error", "Page not available", "GetTokenHandler only accepts a POST")
		return
	}

	// r.ParseForm()
	username := "josh"
	password := "password123"
	// log.Println(r.Form)

	if username == "" || password == "" {
		ReturnMessageJSON(w, "Error", "Invalid Username/Password", "Invalid Username or password in GetTokenHandler")
		return
	}

	// if db.ValidUser(username, password) {
	if true {
		// Create the Claim which expires after EXPIRATION_HOURS hrs, default is 5.
		claims := MyCustomClaims{
			username,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 5).Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		/* Sign the token with our secret */
		tokenString, err := token.SignedString(mySigningKey)
		if err != nil {
			log.Println("Something went wrong with signing token")
			ReturnMessageJSON(w, "Error", "Authentication Failed", "Authentication Failed")

			return
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers:", "Origin, Content-Type, X-Auth-Token")
		w.Header().Set("Content-Type", "application/json")

		user := User{"josh@gmail.com", tokenString}

		js, err := json.Marshal(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(js)

		// w.Write([]byte(tokenString))
	} else {
		ReturnMessageJSON(w, "Error", "Authentication Failed", "Authentication Failed")
	}
}

// func GetSettingsHandler(w http.ResponseWriter, r *http.Request) {
// look in the db for email/password
// and return user info
// if not found return message about it
// }

// ProductsHandler handles the products page, returns all the products.
func GetSettingsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	ok, username := CheckTokenGetUsername(r)
	if !ok {
		ReturnMessageJSON(w, "Error", "You are not authorized to do this, login & try again", "Invalid token")
		return
	}

	if r.Method == "GET" {
		var user User
		// user, err = db.GetUser(username)

		// if err != nil {
		// 	ReturnMessageJSON(w, "Error", "Could't get user", "error")
		// 	log.Println(err)
		// 	return
		// }

		user = User{"josh@gmail.com", username}
		js, err := json.Marshal(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(js)

	}
}

// CheckTokenGetUsername gets the HTTP request as the argument and returns if the token is valid
// which is passed as a header of the name Token and returns the username of the logged in user.
func CheckTokenGetUsername(r *http.Request) (bool, string) {
	token := r.Header["Token"][0]
	ok, username := ValidateToken(token)
	return ok, username
}

//ValidateToken will validate the token
func ValidateToken(myToken string) (bool, string) {
	token, err := jwt.ParseWithClaims(myToken, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(mySigningKey), nil
	})

	if err != nil {
		log.Println("Invalid token.", token)
		return false, ""
	}

	claims := token.Claims.(*MyCustomClaims)
	return token.Valid, claims.Email
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
		panic(err)
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
