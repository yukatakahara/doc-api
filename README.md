# DOC API

## Setup

    go run server.go

## Endpoints

### Register

    curl -v -H "Accept: application/json" -H "Content-type: application/json" -X POST -d '{"email":"foo@gmail.com","password":"password123"}' http://localhost:3000/register

### Get all projects

    curl 0.0.0.0:3000/projects/ | jq '.'

=>

    {
      "Errors": [
        ""
      ],
      "Projects": [
        {
          "idea_description": "instagram but for cats",
          "team_employees": "josh, dan, lea",
          "team_email": "cats@gmail.com",
          "team_name": "cats",
          "employee_id": "123",
          "email": "dan@gmail.com",
          "name": "dan",
          "id": 1
        },
        {
          "idea_description": "social network for dogs",
          "team_employees": "laura, josh",
          "team_email": "dogs@gmail.com",
          "team_name": "dogs",
          "employee_id": "143",
          "email": "laura@gmail.com",
          "name": "laura",
          "id": 2
        }
      ]
    }


### Add project

    curl -X POST -i localhost:3000/projects/ -d '{"name":"david", "email":"foo@gmail.com", "employee_id":"13", "team_name":"The Fools", "team_email":"awesometeam@lists.yp.com", "team_employees":"Manny(manny@yp.com), Mo(mo@yp.com), Jack(jack@yp.com)", "idea_description":"cool project! wow"}'

=>

    HTTP/1.1 201 Create

