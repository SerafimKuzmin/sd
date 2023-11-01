package postgres

import (
	"github.com/SerafimKuzmin/sd/backend/internal/Country/usecase"
	"github.com/SerafimKuzmin/sd/backend/models"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Country struct {
	ID   uint64 `gorm:"column:member_id"`
	Name string `gorm:"column:name"`
}

func (Country) TableName() string {
	return "country"
}

func toPostgresCountry(p *models.Country) *Country {
	return &Country{
		ID:   p.ID,
		Name: p.Name,
	}
}

func toModelCountry(p *Country) *models.Country {
	return &models.Country{
		ID:   p.ID,
		Name: p.Name,
	}
}

func toModelCountrys(Countrys []*Country) []*models.Country {
	out := make([]*models.Country, len(Countrys))

	for i, b := range Countrys {
		out[i] = toModelCountry(b)
	}

	return out
}

type CountryRepository struct {
	db *gorm.DB
}

func (pr CountryRepository) CreateCountry(e *models.Country) error {
	postgresCountry := toPostgresCountry(e)
	tx := pr.db.Create(postgresCountry)

	if tx.Error != nil {
		return errors.Wrap(tx.Error, "database error (table Country)")
	}

	e.ID = postgresCountry.ID
	return nil
}

func (pr CountryRepository) UpdateCountry(e *models.Country) error {
	postgresCountry := toPostgresCountry(e)

	tx := pr.db.Omit("id").Updates(postgresCountry)

	if tx.Error != nil {
		return errors.Wrap(tx.Error, "database error (table Country)")
	}

	return nil
}

func (pr CountryRepository) GetCountry(id uint64) (*models.Country, error) {
	var Country Country

	tx := pr.db.Where("id = ?", id).Take(&Country)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, models.ErrNotFound
	} else if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "database error (table Country)")
	}

	return toModelCountry(&Country), nil
}

func (pr CountryRepository) DeleteCountry(id uint64) error {
	tx := pr.db.Delete(&Country{}, id)

	if tx.Error != nil {
		return errors.Wrap(tx.Error, "database error (table Country)")
	}

	return nil
}

func NewCountryRepository(db *gorm.DB) usecase.RepositoryI {
	return &CountryRepository{
		db: db,
	}
}
