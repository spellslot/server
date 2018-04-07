package services_test

import (
	"github.com/spellslot/server/daos"
	"github.com/spellslot/server/services"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Spell Service", func() {
	Describe("Successful", func() {
		It("Simple mock returns 0 results", func() {
			dao, _ := daos.NewSpellDAO(daos.MockDB)
			svc := services.NewSpellService(dao)
			res, _ := svc.Get()
			Expect(len(*res)).To(Equal(0))
		})
	})
})
