# D&D 5e Spells Server

[![Build Status](https://travis-ci.org/spellslot/server.svg?branch=master&service=github)](https://travis-ci.org/spellslot/server)
[![Coverage Status](https://coveralls.io/repos/github/spellslot/server/badge.svg?branch=master&service=github)](https://coveralls.io/github/spellslot/server?branch=master)

This is the spells server written in Go.

## Project Structure

The main entrypoint for the server is `main.go`.

* `apis` is the router layer that sets up all the HTTP routes with their corresponding services
* `app` holds all app-wide code such as middleware, logging, and versioning
* `daos` contains the data access objects for interactions with persistent storage
* `db` has all of the code to initialize our postgres database
* `mocks` hold all of the mocked code generated by `gomock` for testing purposes
* `models` contains all the data structures needed to communicate between layers
* `services` contains the business logic behind each route

## Set up a development environment

1. Install Go 1.11.x+, and Postgres
2. Set up the Postgres database and seed the table
    * `createuser <username>`
    * `createdb <database name> -U <username>`
    * `sh db/init.sh <database name> <username>`
3. Grab this code
    * `cd $GOHOME/src`
    * `go get github.com/spellslot/server`
4. Create a `.env` file in `$GOHOME/src/github.com/spellslot/server` and populate it like this:

```
PORT=<port>
APP_ENV=development
DATABASE_CONNECTION_STRING="host=localhost port=5432 user=<username> dbname=<database name> sslmode=disable"
```

You're ready to start hacking!

## Testing

1. `cd $GOHOME/src/github.com/spellslot/server`
2. `go test ./... -v`

## Running a local web server

1. `cd $GOHOME/src/github.com/spellslot/server`
2. `go build -o main`
3. `./main`
