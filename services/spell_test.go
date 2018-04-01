package services

import (
	"net/url"
	"testing"

	"github.com/spellslot/server/util"

	"github.com/spellslot/server/daos"
)

func TestSimpleGet(t *testing.T) {
	dao, _ := daos.NewSpellDAO(daos.MockDB)
	svc := NewSpellService(dao)
	res, _ := svc.Get(url.Values{})
	util.AssertEqual(t, len(*res), 0)
}
