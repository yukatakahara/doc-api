# DOC API

## Setup

    cd cmd/web
    go run server.go

## Endpoints

### Admin Login

    curl -v -H "Accept: application/json" -H "Content-type: application/json" POST -d '{"email":"foo@gmail.com","password":"password123"}' "http://localhost:3000/adminlogin"

### Patient Signup

    curl -v -H "Accept: application/json" -H "Content-type: application/json" POST -d '{"email":"foo@gmail.com","password":"password123"}' "http://localhost:3000/signup"

### Patient Login

    curl -v -H "Accept: application/json" -H "Content-type: application/json" POST -d '{"email":"foo@gmail.com","password":"password123"}' "http://localhost:3000/login"

### Get Clinics

    curl -v -H "Accept: application/json" -H "Content-type: application/json" "http://localhost:3000/doctors?lat=1.298653&lon=103.848456"

### Get User Settings

    curl -v -H "Accept: application/json" -H "Content-type: application/json" -d '{"email":"foo@gmail.com","password":"password123"}' "http://localhost:3000/login"
