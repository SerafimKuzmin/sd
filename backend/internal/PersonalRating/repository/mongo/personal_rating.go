package mongo

import (
	"context"
	"github.com/SerafimKuzmin/sd/backend/internal/PersonalRating/repository"
	"github.com/SerafimKuzmin/sd/backend/models"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type PersonalRatingalRating struct {
	ID     uint64  `bson:":rating_id"`
	UserID uint64  `bson:":user_id"`
	FilmID uint64  `bson:":movie_id"`
	Rate   float64 `bson:":user_rating"`
}

func (PersonalRatingalRating) TableName() string {
	return "rating"
}

func toClickhousePersonalRatingalRating(t *models.PersonalRating) *PersonalRatingalRating {
	return &PersonalRatingalRating{
		ID:     t.ID,
		UserID: t.UserID,
		FilmID: t.FilmID,
		Rate:   t.Rate,
	}
}

type PersonalRatingRepository struct {
	db *mongo.Database
}

func NewPersonalRatingRepository(db *mongo.Database) repository.RepositoryI {
	return &PersonalRatingRepository{
		db: db,
	}
}

func (gr PersonalRatingRepository) CreatePersonalRating(g *models.PersonalRating) error {
	MongoPersonalRating := interface{}(bson.Marshal(g))
	_, err := gr.db.Collection("PersonalRating").InsertOne(context.TODO(), MongoPersonalRating)

	if err != nil {
		return errors.Wrap(err, "database error (table PersonalRating)")
	}
	return nil
}

func (gr PersonalRatingRepository) UpdatePersonalRating(g *models.PersonalRating) error {
	MongoPersonalRating := interface{}(bson.Marshal(g))
	_, err := gr.db.Collection("PersonalRating").UpdateByID(context.TODO(), g.ID, MongoPersonalRating)

	if err != nil {
		return errors.Wrap(err, "database error (table PersonalRating)")
	}

	return nil
}

func (gr PersonalRatingRepository) GetPersonalRating(id uint64) (*models.PersonalRating, error) {
	filter := bson.D{{"id", id}}

	var result models.PersonalRating
	err := gr.db.Collection("PersonalRating").FindOne(context.TODO(), filter).Decode(result)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, models.ErrNotFound
	} else if err != nil {
		return nil, errors.Wrap(err, "database error (table PersonalRating)")
	}

	return &result, nil
}

func (gr PersonalRatingRepository) DeletePersonalRating(id uint64) error {
	filter := bson.D{{"id", id}}
	_, err := gr.db.Collection("PersonalRating").DeleteOne(context.TODO(), filter)

	if err != nil {
		return errors.Wrap(err, "database error (table PersonalRating)")
	}

	return nil
}
