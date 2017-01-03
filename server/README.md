# CRUD API

## Setup

    go run main.go

## Endpoints

### Get all projects

    curl 0.0.0.0:3001/doctors/ | jq '.'

=>

    {
      "Errors": [
        ""
      ],
      "Doctors": [
        {
          "employee_id": "123",
          "email": "dan@gmail.com",
          "name": "dan",
          "id": 1
        },
        {
          "employee_id": "143",
          "email": "laura@gmail.com",
          "name": "laura",
          "id": 2
        }
      ]
    }


### Add project

    curl -X POST -i localhost:3001/doctors/ -d '{"name":"david", "email":"foo@gmail.com", "employee_id":"13"}'

=>

    HTTP/1.1 201 Create

