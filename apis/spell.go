package apis

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/spellslot/server/app"
	"github.com/spellslot/server/services"
)

type (
	// spellResource defines the handlers for the CRUD APIs.
	spellResource struct {
		service services.SpellService
	}
)

// ServeSpellResource sets up routes for the spell service
func ServeSpellResource(router *httprouter.Router, service services.SpellService) {
	r := &spellResource{service}
	router.Handler("GET", "/api/v1/spells", app.Middleware(http.HandlerFunc(r.get)))
	router.Handler("POST", "/api/v1/spells", app.Middleware(http.HandlerFunc(r.post)))
	router.Handler("PUT", "/api/v1/spells", app.Middleware(http.HandlerFunc(r.put)))
	router.Handler("DELETE", "/api/v1/spells", app.Middleware(http.HandlerFunc(r.delete)))
}

func (s *spellResource) get(w http.ResponseWriter, r *http.Request) {
	spells, _ := s.service.Get()
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
