package apis_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alexsasharegan/dotenv"
	"github.com/gorilla/mux"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spellslot/server/apis"
	"github.com/spellslot/server/daos"
	"github.com/spellslot/server/services"
)

func TestApis(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "APIs Suite")
}

var _ = BeforeSuite(func() {
	err := dotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
})

func executeRequest(req *http.Request, spellDaoVersion daos.SpellDAOVersion) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router := mux.NewRouter().StrictSlash(true)
	spellDao, _ := daos.NewSpellDAO(spellDaoVersion)
	apis.ServeSpellResource(router, services.NewSpellService(spellDao))
	router.ServeHTTP(rr, req)

	return rr
}
