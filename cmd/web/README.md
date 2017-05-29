## Endpoints

### Admin Login

    curl -v -H "Accept: application/json" -H "Content-type: application/json" POST -d '{"email":"vera@gmail.com","password":"112233"}' "http://localhost:3000/adminlogin"

### Admin create clinic

    curl -v -H "Authorization: Bearer verylongtokenstring" -H "Accept: application/json" -H "Content-type: application/json" POST -d '{"name":"a nice place","address1":"2323 hesting road, singapore"}' "http://localhost:3000/clinics"

### Admin get all clinics

    curl -v -H "Authorization: Bearer verylongtokenstring" -H "Accept: application/json" -H "Content-type: application/json" "http://localhost:3000/clinics?lat=1.298653&lon=103.848456"

### Patient Signup

    curl -v -H "Accept: application/json" -H "Content-type: application/json" POST -d '{"email":"foo@gmail.com","password":"password123"}' "http://localhost:3000/signup"

### Patient Login

    curl -v -H "Accept: application/json" -H "Content-type: application/json" POST -d '{"email":"foo@gmail.com","password":"password123"}' "http://localhost:3000/login"


### Get User Settings

    curl -v -H "Accept: application/json" -H "Content-type: application/json" -d '{"email":"foo@gmail.com","password":"password123"}' "http://localhost:3000/login"
