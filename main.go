package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/spellslot/server/apis"
	"github.com/spellslot/server/daos"
	"github.com/spellslot/server/services"

	"github.com/alexsasharegan/dotenv"
	"github.com/julienschmidt/httprouter"
)

func main() {
	err := dotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	// router := mux.NewRouter().StrictSlash(true)
	router := httprouter.New()

	if os.Getenv("APP_ENV") == "production" {
		fs := http.FileServer(http.Dir("client/build"))
		router.Handler("GET", "/", fs)
	}

	spellDao, _ := daos.NewSpellDAO()

	apis.ServeSpellResource(router, services.NewSpellService(spellDao))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", os.Getenv("PORT")), router))
}
