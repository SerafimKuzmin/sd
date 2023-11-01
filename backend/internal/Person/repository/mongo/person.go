package mongo

import (
	"context"
	"github.com/SerafimKuzmin/sd/backend/internal/Person/usecase"
	"github.com/SerafimKuzmin/sd/backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Person struct {
	ID   uint64 `gorm:"member_id"`
	Name string `gorm:"name"`
}

func (Person) TableName() string {
	return "person"
}

func toClickhousePerson(p *models.Person) *Person {
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
	db *mongo.Database
}

func NewPersonRepository(db *mongo.Database) usecase.RepositoryI {
	return &PersonRepository{
		db: db,
	}
}

func (gr PersonRepository) CreatePerson(g *models.Person) error {
	MongoPerson := interface{}(bson.Marshal(g))
	_, err := gr.db.Collection("Person").InsertOne(context.TODO(), MongoPerson)

	if err != nil {
		return errors.Wrap(err, "database error (table Person)")
	}
	return nil
}

func (gr PersonRepository) UpdatePerson(g *models.Person) error {
	MongoPerson := interface{}(bson.Marshal(g))
	_, err := gr.db.Collection("Person").UpdateByID(context.TODO(), g.ID, MongoPerson)

	if err != nil {
		return errors.Wrap(err, "database error (table Person)")
	}

	return nil
}

func (gr PersonRepository) GetPerson(id uint64) (*models.Person, error) {
	filter := bson.D{{"id", id}}

	var result models.Person
	err := gr.db.Collection("Person").FindOne(context.TODO(), filter).Decode(result)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, models.ErrNotFound
	} else if err != nil {
		return nil, errors.Wrap(err, "database error (table Person)")
	}

	return &result, nil
}

func (gr PersonRepository) DeletePerson(id uint64) error {
	filter := bson.D{{"id", id}}
	_, err := gr.db.Collection("Person").DeleteOne(context.TODO(), filter)

	if err != nil {
		return errors.Wrap(err, "database error (table Person)")
	}

	return nil
}
