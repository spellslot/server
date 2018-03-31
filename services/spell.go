package services

import (
	"net/url"

	"github.com/spellslot/server/daos"
	"github.com/spellslot/server/models"
)

type (
	// SpellService specifies the interface for the spell service needed by spellResource.
	SpellService interface {
		Get(queryString url.Values) (*models.Spells, error)
	}

	// SpellServiceImpl provides services related with spells.
	SpellServiceImpl struct{ dao daos.SpellDAO }
)

// NewSpellService creates a new SpellService with the given spell DAO.
func NewSpellService(dao daos.SpellDAO) SpellService {
	return &SpellServiceImpl{dao}
}

// Get for SpellService
func (s SpellServiceImpl) Get(queryString url.Values) (*models.Spells, error) {
	return s.dao.Get(queryString)
}
