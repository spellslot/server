package daos

import (
	"net/url"
	"testing"

	"github.com/spellslot/server/util"
)

// TODO need to load .env for every test, switch to ginkgo or something

func TestMockGet(t *testing.T) {
	dao, _ := NewSpellDAO(MockDB)
	res, _ := dao.Get(url.Values{})
	util.AssertEqual(t, len(*res), 0)
}

func TestRealGet(t *testing.T) {
	dao, _ := NewSpellDAO(RealDB)
	res, _ := dao.Get(url.Values{})
	util.AssertEqual(t, len(*res), 408)
}
