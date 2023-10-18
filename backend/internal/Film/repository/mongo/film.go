package mongo

import (
	"context"
	"github.com/SerafimKuzmin/sd/backend/internal/Film/repository"
	"github.com/SerafimKuzmin/sd/backend/models"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"time"
)

type Film struct {
	ID          uint64    `bson:"modie_id"`
	Name        string    `bson:"name"`
	Description string    `bson:"description"`
	Rate        float64   `bson:"rate"`
	Genre       string    `bson:"genre"`
	ReleaseDT   time.Time `bson:"release_dt"`
	Duration    uint      `bson:"duration"`
	FilmID      uint64    `bson:"Film_id"`
}

func (Film) TableName() string {
	return "movie"
}

type FilmPerson struct {
	FilmID   uint64 `bson:"movie_id"`
	PersonID uint64 `bson:"member_id"`
}

func (FilmPerson) TableName() string {
	return "moviemember"
}

func toMongoFilm(g *models.Film) *Film {
	return &Film{
		ID:          g.ID,
		Name:        g.Name,
		Description: g.Description,
		Rate:        g.Rate,
		Genre:       g.Genre,
		ReleaseDT:   g.ReleaseDT,
		Duration:    g.Duration,
	}
}

func toModelFilm(g *Film) *models.Film {
	return &models.Film{
		ID:          g.ID,
		Name:        g.Name,
		Description: g.Description,
		Rate:        g.Rate,
		Genre:       g.Genre,
		ReleaseDT:   g.ReleaseDT,
		Duration:    g.Duration,
	}
}

func toModelFilms(Films []*Film) []*models.Film {
	out := make([]*models.Film, len(Films))

	for i, b := range Films {
		out[i] = toModelFilm(b)
	}

	return out
}

type FilmRepository struct {
	db *mongo.Database
}

func NewFilmRepository(db *mongo.Database) repository.RepositoryI {
	return &FilmRepository{
		db: db,
	}
}

func (gr FilmRepository) CreateFilm(g *models.Film) error {
	MongoFilm := interface{}(bson.Marshal(g))
	_, err := gr.db.Collection("film").InsertOne(context.TODO(), MongoFilm)

	if err != nil {
		return errors.Wrap(err, "database error (table film)")
	}
	return nil
}

func (gr FilmRepository) UpdateFilm(g *models.Film) error {
	MongoFilm := interface{}(bson.Marshal(g))
	_, err := gr.db.Collection("film").UpdateByID(context.TODO(), g.ID, MongoFilm)

	if err != nil {
		return errors.Wrap(err, "database error (table film)")
	}

	return nil
}

func (gr FilmRepository) GetFilm(id uint64) (*models.Film, error) {
	filter := bson.D{{"id", id}}

	var result models.Film
	err := gr.db.Collection("film").FindOne(context.TODO(), filter).Decode(result)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, models.ErrNotFound
	} else if err != nil {
		return nil, errors.Wrap(err, "database error (table film)")
	}

	return &result, nil
}

func (gr FilmRepository) DeleteFilm(id uint64) error {
	filter := bson.D{{"id", id}}
	_, err := gr.db.Collection("film").DeleteOne(context.TODO(), filter)

	if err != nil {
		return errors.Wrap(err, "database error (table film)")
	}

	return nil
}

func (gr FilmRepository) GetFilmByPerson(id uint64) ([]*models.Film, error) {
	filter := bson.D{{"film_person", id}}

	cursor, err := gr.db.Collection("film_person").Find(context.TODO(), filter)
	var results []Film

	if err != nil {
		return nil, errors.Wrap(err, "database error (table film)")
	}
	if err == cursor.All(context.TODO(), &results) {
		return nil, models.ErrNotFound
	} else if err != nil {
		return nil, errors.Wrap(err, "database error (table film)")
	}
	films := make([]*models.Film, 0, 10)

	for idx := range results {
		filter := bson.D{{"id", idx}}

		var result models.Film
		err := gr.db.Collection("film").FindOne(context.TODO(), filter).Decode(result)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrNotFound
		} else if err != nil {
			return nil, errors.Wrap(err, "database error (table film)")
		}

		films = append(films, &result)
	}

	return films, nil
}

func (gr FilmRepository) GetFilmByCountry(id uint64) ([]*models.Film, error) {
	filter := bson.D{{"country_id", id}}

	cursor, err := gr.db.Collection("film").Find(context.TODO(), filter)
	var results []*models.Film

	if err == cursor.All(context.TODO(), &results) {
		return nil, models.ErrNotFound
	} else if err != nil {
		return nil, errors.Wrap(err, "database error (table film)")
	}

	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	return results, nil
}
