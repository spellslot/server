package apis_test

import (
	"net/http"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spellslot/server/mocks"
	"github.com/spellslot/server/models"
)

var _ = Describe("Spells APIs", func() {
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

	Describe("Get", func() {
		Context("Successful", func() {
			It("Simple mock request should return an empty array", func() {
				mockSpellDao.EXPECT().Get().Return(&models.Spells{}, nil)
				req, _ := http.NewRequest("GET", "/api/v1/spells", nil)
				response := executeRequest(req, mockSpellDao)
				Expect(response.Code).To(Equal(http.StatusOK))
				Expect(response.Body.String()).To(Equal("[]"))
			})
		})
	})

	Describe("Post", func() {
		Context("Successful", func() {
			It("Not yet implemented", func() {
				req, _ := http.NewRequest("POST", "/api/v1/spells", nil)
				response := executeRequest(req, mockSpellDao)
				Expect(response.Code).To(Equal(http.StatusOK))
			})
		})
	})

	Describe("Put", func() {
		Context("Successful", func() {
			It("Not yet implemented", func() {
				req, _ := http.NewRequest("PUT", "/api/v1/spells", nil)
				response := executeRequest(req, mockSpellDao)
				Expect(response.Code).To(Equal(http.StatusOK))
				Expect(response.Body.String()).To(Equal("Not yet implemented"))
			})
		})
	})

	Describe("Delete", func() {
		Context("Successful", func() {
			It("Not yet implemented", func() {
				req, _ := http.NewRequest("DELETE", "/api/v1/spells", nil)
				response := executeRequest(req, mockSpellDao)
				Expect(response.Code).To(Equal(http.StatusOK))
				Expect(response.Body.String()).To(Equal("Not yet implemented"))
			})
		})
	})
})
