package daos

import (
	"log"
	"testing"

	"github.com/alexsasharegan/dotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDaos(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DAO Suite")
}

var _ = BeforeSuite(func() {
	err := dotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
})
