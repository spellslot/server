package apis

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/spellslot/server/app"
	"github.com/spellslot/server/models"
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
	router.GET("/api/v1/spells", app.Middleware(r.get))
	router.POST("/api/v1/spells", app.Middleware(r.post))
	router.PUT("/api/v1/spells/:id", app.Middleware(r.put))
	router.DELETE("/api/v1/spells/:id", app.Middleware(r.delete))
}

func (s *spellResource) get(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	spells, err := s.service.Get()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	payload, err := json.Marshal(spells)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	sendResponse(w, "application/json", payload)
}

func (s *spellResource) post(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var spell models.Spell
	err = json.Unmarshal(body, &spell)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	createdSpell, err := s.service.Create(&spell)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	payload, err := json.Marshal(createdSpell)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	sendResponse(w, "application/json", payload)
}

func (s *spellResource) put(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sendResponse(w, "application/json", []byte("Not yet implemented"))
}

func (s *spellResource) delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sendResponse(w, "application/json", []byte("Not yet implemented"))
}
