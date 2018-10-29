package services

import (
	"github.com/spellslot/server/daos"
	"github.com/spellslot/server/models"
)

type (
	// SpellService specifies the interface for the spell service needed by spellResource.
	SpellService interface {
		Get() (*models.Spells, error)
		Create(spell *models.Spell) (*models.Spell, error)
	}

	// SpellServiceImpl provides services related with spells.
	SpellServiceImpl struct{ dao daos.SpellDAO }
)

// NewSpellService creates a new SpellService with the given spell DAO.
func NewSpellService(dao daos.SpellDAO) SpellService {
	return &SpellServiceImpl{dao}
}

// Get for SpellServiceImpl
func (s SpellServiceImpl) Get() (*models.Spells, error) {
	return s.dao.Get()
}

// Create for SpellServiceImpl
func (s SpellServiceImpl) Create(spell *models.Spell) (*models.Spell, error) {
	return s.dao.Create(spell)
}
