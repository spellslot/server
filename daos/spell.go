package daos

import (
	"database/sql"
	"os"

	"github.com/spellslot/server/models"

	// need unnamed to use the postgres driver
	_ "github.com/lib/pq"
)

type (
	// SpellDAO describes a spell dao
	SpellDAO interface {
		Get() (*models.Spells, error)
		Create(spell *models.Spell) (*models.Spell, error)
	}

	// DBSpellDAO persists spell data in database
	DBSpellDAO struct{}
)

// NewSpellDAO creates a new SpellDAO
func NewSpellDAO() (SpellDAO, error) {
	return &DBSpellDAO{}, nil
}

// Get for DBSpellDAO
func (dao *DBSpellDAO) Get() (*models.Spells, error) {
	spells := models.Spells{}
	db, err := sql.Open("postgres", os.Getenv("DATABASE_CONNECTION_STRING"))

	if err != nil {
		return &models.Spells{}, err
	}

	rows, err := db.Query("SELECT * FROM spells;")

	if err != nil {
		return &models.Spells{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var spell models.Spell
		if err := rows.Scan(&spell.ID, &spell.Name, &spell.School, &spell.Level, &spell.Ritual, &spell.CastingTime, &spell.Source, &spell.Range, &spell.Classes, &spell.Components, &spell.Duration, &spell.AtHigherLevel, &spell.Concentration, &spell.Slug, &spell.Page, &spell.Description); err != nil {
			return &models.Spells{}, err
		}
		spells = append(spells, spell)
	}

	if err := rows.Err(); err != nil {
		return &models.Spells{}, err
	}

	return &spells, nil
}

// Create for DBSpellDAO
func (dao *DBSpellDAO) Create(spell *models.Spell) (*models.Spell, error) {
	return &models.Spell{}, nil // TODO
}
