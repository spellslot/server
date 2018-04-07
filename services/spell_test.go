package services_test

import (
	"github.com/golang/mock/gomock"
	"github.com/spellslot/server/mocks"
	"github.com/spellslot/server/models"
	"github.com/spellslot/server/services"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Spell Service", func() {
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

	Describe("Successful", func() {
		It("Simple mock returns 0 results", func() {
			mockSpellDao.EXPECT().Get().Return(&models.Spells{}, nil)
			svc := services.NewSpellService(mockSpellDao)
			res, _ := svc.Get()
			Expect(len(*res)).To(Equal(0))
		})
	})
})
