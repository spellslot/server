package apis_test

import (
	"net/http"

	"github.com/spellslot/server/daos"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Spells APIs", func() {
	Describe("Successful", func() {
		It("Simple mock request should return an empty array", func() {
			req, _ := http.NewRequest("GET", "/api/v1/spells", nil)
			response := executeRequest(req, daos.MockDB)
			Expect(response.Code).To(Equal(http.StatusOK))
			Expect(response.Body.String()).To(Equal("[]"))
		})
	})
})
