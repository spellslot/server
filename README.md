# D&D 5e Spells Server

This is the spells server written in Go.

## Project Structure

The main entrypoint for the server is `main.go`.

* `apis` is the router layer that sets up all the HTTP routes with their corresponding services
* `app` holds all app-wide code such as middleware, logging, and versioning
* `daos` contains the data access objects for interactions with persistent storage
* `models` contains all the data structures needed to communicate between layers
* `services` contains the business logic behind each route

## Hacking

1. Install Go, Go Dep, and Postgres
2. Create a database and use `db/create_spells_table.sql` to initialize the spells table in that database
    * `psql -d <database name> -U <username> -f db/create_spells_table.sql`
3. `go get github.com/spellslot/server`
4. Create a `.env` file in this folder to hold some important information
    * `PORT` is the port you want to run the server on
    * `APP_ENV` specifies development, testing, or production environment
    * `DATABASE_CONNECTION_STRING` is the connection string to your database

* To test: `go test ./...`
* To build: `go build -o main`
* To run: `./main`
