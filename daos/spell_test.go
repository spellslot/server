package daos

import (
	"net/url"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Spells DAO", func() {
	Describe("Mock DB", func() {
		It("Basic Get returns 0 rows", func() {
			dao, _ := NewSpellDAO(MockDB)
			res, _ := dao.Get(url.Values{})
			Expect(len(*res)).To(Equal(0))
		})
	})

	Describe("Real DB", func() {
		It("Basic Get returns 408 rows", func() {
			dao, _ := NewSpellDAO(RealDB)
			res, _ := dao.Get(url.Values{})
			Expect(len(*res)).To(Equal(408))
		})
	})
})
