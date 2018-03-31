package apis

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/spellslot/server/daos"
	"github.com/spellslot/server/services"
	"github.com/spellslot/server/util"

	"github.com/gorilla/mux"
)

func executeRequest(req *http.Request, spellDaoVersion daos.SpellDAOVersion) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router := mux.NewRouter().StrictSlash(true)
	spellDao, _ := daos.NewSpellDAO(spellDaoVersion)
	ServeSpellResource(router, services.NewSpellService(spellDao))
	router.ServeHTTP(rr, req)

	return rr
}

func TestSimpleMockGetSpells(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/v1/spells", nil)
	response := executeRequest(req, daos.MockDB)

	util.CheckResponseCode(t, http.StatusOK, response.Code)

	util.AssertEqual(t, response.Body.String(), "[]")
}
