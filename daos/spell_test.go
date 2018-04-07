package daos_test

import (
	"net/url"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spellslot/server/daos"
)

var _ = Describe("Spells DAO", func() {
	Describe("Mock DB", func() {
		It("Basic Get returns 0 rows", func() {
			dao, _ := daos.NewSpellDAO(daos.MockDB)
			res, _ := dao.Get(url.Values{})
			Expect(len(*res)).To(Equal(0))
		})
	})

	Describe("Real DB", func() {
		It("Basic Get returns 408 rows", func() {
			dao, _ := daos.NewSpellDAO(daos.RealDB)
			res, _ := dao.Get(url.Values{})
			Expect(len(*res)).To(Equal(408))
		})
	})
})
