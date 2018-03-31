package apis

import (
	"encoding/json"
	"net/http"
	"server/services"

	"github.com/gorilla/mux"
)

type (
	// spellResource defines the handlers for the CRUD APIs.
	spellResource struct {
		service services.SpellService
	}
)

// ServeSpellResource sets up routes for the spell service
func ServeSpellResource(router *mux.Router, service services.SpellService) {
	r := &spellResource{service}
	router.Methods("GET").Path("/api/v1/spells").Name("Get Spell").HandlerFunc(r.get)
	router.Methods("POST").Path("/api/v1/spells").Name("Post Spell").HandlerFunc(r.post)
	router.Methods("PUT").Path("/api/v1/spells").Name("Put Spell").HandlerFunc(r.put)
	router.Methods("DELETE").Path("/api/v1/spells").Name("Delete Spell").HandlerFunc(r.delete)
}

func (s *spellResource) get(w http.ResponseWriter, r *http.Request) {
	spells, _ := s.service.Get(r.URL.Query())
	payload, _ := json.Marshal(spells)
	sendResponse(w, "application/json", payload)
}

func (s *spellResource) post(w http.ResponseWriter, r *http.Request) {
	sendResponse(w, "application/json", []byte("Not yet implemented"))
}

func (s *spellResource) put(w http.ResponseWriter, r *http.Request) {
	sendResponse(w, "application/json", []byte("Not yet implemented"))
}

func (s *spellResource) delete(w http.ResponseWriter, r *http.Request) {
	sendResponse(w, "application/json", []byte("Not yet implemented"))
}
