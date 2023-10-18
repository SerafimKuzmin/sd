package postgres

import (
	"github.com/SerafimKuzmin/sd/backend/internal/Person/repository"
	"github.com/SerafimKuzmin/sd/backend/models"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Person struct {
	ID   uint64 `gorm:"column:member_id"`
	Name string `gorm:"column:name"`
}

func (Person) TableName() string {
	return "person"
}

func toPostgresPerson(p *models.Person) *Person {
	return &Person{
		ID:   p.ID,
		Name: p.Name,
	}
}

func toModelPerson(p *Person) *models.Person {
	return &models.Person{
		ID:   p.ID,
		Name: p.Name,
	}
}

func toModelPersons(Persons []*Person) []*models.Person {
	out := make([]*models.Person, len(Persons))

	for i, b := range Persons {
		out[i] = toModelPerson(b)
	}

	return out
}

type PersonRepository struct {
	db *gorm.DB
}

func (pr PersonRepository) CreatePerson(e *models.Person) error {
	postgresPerson := toPostgresPerson(e)
	tx := pr.db.Create(postgresPerson)

	if tx.Error != nil {
		return errors.Wrap(tx.Error, "database error (table Person)")
	}

	e.ID = postgresPerson.ID
	return nil
}

func (pr PersonRepository) UpdatePerson(e *models.Person) error {
	postgresPerson := toPostgresPerson(e)

	tx := pr.db.Omit("id").Updates(postgresPerson)

	if tx.Error != nil {
		return errors.Wrap(tx.Error, "database error (table Person)")
	}

	return nil
}

func (pr PersonRepository) GetPerson(id uint64) (*models.Person, error) {
	var Person Person

	tx := pr.db.Where("id = ?", id).Take(&Person)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, models.ErrNotFound
	} else if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "database error (table Person)")
	}

	return toModelPerson(&Person), nil
}

func (pr PersonRepository) DeletePerson(id uint64) error {
	tx := pr.db.Delete(&Person{}, id)

	if tx.Error != nil {
		return errors.Wrap(tx.Error, "database error (table Person)")
	}

	return nil
}

func NewPersonRepository(db *gorm.DB) repository.RepositoryI {
	return &PersonRepository{
		db: db,
	}
}
