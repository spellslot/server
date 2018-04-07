package apis_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/spellslot/server/mocks"

	"github.com/alexsasharegan/dotenv"
	"github.com/julienschmidt/httprouter"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spellslot/server/apis"
	"github.com/spellslot/server/services"
)

func TestApis(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "APIs Suite")
}

var _ = BeforeSuite(func() {
	err := dotenv.Load("../.env")
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}
})

func executeRequest(req *http.Request, spellDao *mocks.MockSpellDAO) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router := httprouter.New()
	apis.ServeSpellResource(router, services.NewSpellService(spellDao))
	router.ServeHTTP(rr, req)

	return rr
}
