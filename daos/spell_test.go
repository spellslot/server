package daos_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spellslot/server/daos"
	"github.com/spellslot/server/mocks"
	"github.com/spellslot/server/models"
)

var _ = Describe("Spells DAO", func() {
	Describe("Mock DB", func() {
		var (
			mockCtrl     *gomock.Controller
			mockSpellDao *mocks.MockSpellDAO
		)

		BeforeEach(func() {
			mockCtrl = gomock.NewController(GinkgoT())
			mockSpellDao = mocks.NewMockSpellDAO(mockCtrl)
		})

		AfterEach(func() {
			mockCtrl.Finish()
		})

		It("Basic Get returns 0 rows", func() {
			mockSpellDao.EXPECT().Get().Return(&models.Spells{}, nil)
			res, _ := mockSpellDao.Get()
			Expect(len(*res)).To(Equal(0))
		})
	})

	Describe("Real DB", func() {
		var (
			spellDao daos.SpellDAO
		)

		BeforeEach(func() {
			spellDao, _ = daos.NewSpellDAO()
		})

		It("Basic Get returns 408 rows", func() {
			res, _ := spellDao.Get()
			Expect(len(*res)).To(Equal(408))
		})
	})
})
