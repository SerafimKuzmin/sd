package mongo

import (
	"context"
	"github.com/SerafimKuzmin/sd/backend/internal/Country/usecase"
	"github.com/SerafimKuzmin/sd/backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Country struct {
	ID   uint64 `bson:"member_id"`
	Name string `bson:"name"`
}

func (Country) TableName() string {
	return "country"
}

func toClickhouseCountry(p *models.Country) *Country {
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
	db *mongo.Database
}

func NewCountryRepository(db *mongo.Database) usecase.RepositoryI {
	return &CountryRepository{
		db: db,
	}
}

func (pr CountryRepository) CreateCountry(e *models.Country) error {
	ClickhouseCountry := interface{}(bson.Marshal(e))
	_, err := pr.db.Collection("country").InsertOne(context.TODO(), ClickhouseCountry)

	if err != nil {
		return errors.Wrap(err, "database error (table Country)")
	}

	return nil
}

func (pr CountryRepository) UpdateCountry(e *models.Country) error {
	ClickhouseCountry := interface{}(bson.Marshal(e))
	_, err := pr.db.Collection("country").UpdateByID(context.TODO(), e.ID, ClickhouseCountry)

	if err != nil {
		return errors.Wrap(err, "database error (table Country)")
	}

	return nil
}

func (pr CountryRepository) GetCountry(id uint64) (*models.Country, error) {
	filter := bson.D{{"id", id}}

	var result models.Country
	err := pr.db.Collection("country").FindOne(context.TODO(), filter).Decode(result)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, models.ErrNotFound
	} else if err != nil {
		return nil, errors.Wrap(err, "database error (table Country)")
	}

	return &result, nil
}

func (pr CountryRepository) DeleteCountry(id uint64) error {
	filter := bson.D{{"id", id}}
	_, err := pr.db.Collection("country").DeleteOne(context.TODO(), filter)

	if err != nil {
		return errors.Wrap(err, "database error (table Country)")
	}

	return nil
}
