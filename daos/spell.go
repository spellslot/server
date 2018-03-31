package daos

import (
	"database/sql"
	"errors"
	"log"
	"net/url"
	"os"
	"server/models"

	// need unnamed to use the postgres driver
	_ "github.com/lib/pq"
)

type (
	// SpellDAOVersion is an enum used when creating a new dao through the factory
	SpellDAOVersion int

	// SpellDAO describes a spell dao
	SpellDAO interface {
		Get(queryString url.Values) (*models.Spells, error)
	}

	// DBSpellDAO persists spell data in database
	DBSpellDAO struct{}

	// MockSpellDAO mocks spell data
	MockSpellDAO struct{}
)

const (
	// RealDB describes a connection to a real DB
	RealDB SpellDAOVersion = iota
	// MockDB describes a connection to a mock DB
	MockDB
)

// NewSpellDAO creates a new SpellDAO
func NewSpellDAO(version SpellDAOVersion) (SpellDAO, error) {
	switch version {
	case RealDB:
		return &DBSpellDAO{}, nil
	case MockDB:
		return &MockSpellDAO{}, nil
	}
	return nil, errors.New("DAO version not implemented")
}

// Get for DBSpellDAO
func (dao *DBSpellDAO) Get(queryString url.Values) (*models.Spells, error) {
	spells := models.Spells{}
	db, err := sql.Open("postgres", os.Getenv("DATABASE_CONNECTION_STRING"))

	if err != nil {
		log.Fatal(err)
	}

	// TODO have to deal with different types even though queryString only holds slices of strings
	// where := queryStringToWhereClause(queryString)
	// rows, err := db.Query("SELECT * FROM spells " + where + ";")

	rows, err := db.Query("SELECT * FROM spells;")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var spell models.Spell
		if err := rows.Scan(&spell.ID, &spell.Name, &spell.School, &spell.Level, &spell.Ritual, &spell.CastingTime, &spell.Source, &spell.Range, &spell.Classes, &spell.Components, &spell.Duration, &spell.AtHigherLevel, &spell.Concentration, &spell.Slug, &spell.Page, &spell.Description); err != nil {
			log.Fatal(err)
		}
		spells = append(spells, spell)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return &spells, nil
}

// Get for MockSpellDAO
func (dao *MockSpellDAO) Get(queryString url.Values) (*models.Spells, error) {
	return &models.Spells{}, nil
}
