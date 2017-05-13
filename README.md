# DOC API

## Setup

    cd cmd/web
    go run server.go

## Endpoints

### Signup

    curl -v -H "Accept: application/json" -H "Content-type: application/json" -X POST -d '{"email":"foo@gmail.com","password":"password123"}' "http://localhost:3000/signup"

### Login

    curl -v -H "Accept: application/json" -H "Content-type: application/json" -X POST -d '{"email":"foo@gmail.com","password":"password123"}' "http://localhost:3000/login"

### Get Doctors

    curl -v -H "Accept: application/json" -H "Content-type: application/json" "http://localhost:3000/doctors?lat=1.298653&lon=103.848456"

### Get User Settings

    curl -v -H "Accept: application/json" -H "Content-type: application/json" -d '{"email":"foo@gmail.com","password":"password123"}' "http://localhost:3000/login"
