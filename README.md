# CHARACTERS CRUD

This is a simple project for learning about Gin framework. It has simple CRUD operations related to fictional characters, for now, it can just read all the characters, character by ID and can creates character. It has an authentication system included that returns a JWT for sending the requests to the protected endpoints.

Project structure:

```
├── go.mod
├── go.sum
├── internal
│   ├── api
│   │   ├── authentication
│   │   ├── characters
│   │   │   ├── data
│   │   │   │   ├── daoImplementation.go
│   │   │   │   └── daoInterface.go
│   │   │   ├── handlers
│   │   │   │   └── handlers.go
│   │   │   ├── models
│   │   │   │   └── character.go
│   │   │   ├── responses
│   │   │   │   └── response.go
│   │   │   └── routes
│   │   │       └── charRoutes.go
│   │   └── server
│   │       ├── ping
│   │       │   ├── pingHandler.go
│   │       │   ├── pingResponse.go
│   │       │   └── pingRoute.go
│   │       └── server.go
│   └── database
│       ├── dbInterface.go
│       ├── migrations.go
│       └── sqlite.go
├── main.go
├── pkg
├── README.md
```

And this is a prototype of the requests flow using the authentication service

![](https://github.com/Edmartt/characters-crud/blob/dev/assets/concept.png)

## Requirements

 - Go 1.19+
 - Docker
 - SQLite

## Running Locally

 ```
 git clone https://github.com/Edmartt/characters-crud.git
 ```

 or ssh instead:

 ```
git clone git@github.com:Edmartt/characters-crud.git
 ```

 browse into project directory:

 ```
 cd characters-crud/
 ```

 download dependencies

 ```
 go mod tidy
 ```

 set environment variables following the .env.example file and run

 ```
 go run main.go
 ```
