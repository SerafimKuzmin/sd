package mongo

import (
	"context"
	"github.com/SerafimKuzmin/sd/backend/internal/Auth/repository"
	"github.com/SerafimKuzmin/sd/backend/models"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
	"time"
)

type authRepository struct {
	db *mongo.Database
}

type Cookie struct {
	UserID       *uint64   `bson:"user_id"`
	SessionToken string    `bson:"session_token"`
	ExpireTime   time.Time `bson:"expire_time"`
}

func NewAuthRepository(db *mongo.Database) repository.RepositoryI {
	return &authRepository{
		db: db,
	}
}

func (ar authRepository) CreateCookie(cookie *models.Cookie) error {

	mongoCookie := interface{}(bson.Marshal(cookie))
	_, err := ar.db.Collection("cookie").InsertOne(context.TODO(), mongoCookie)

	if err != nil {
		return errors.Wrap(err, "database error (table cookie)")
	}

	return nil
}

func (ar authRepository) GetUserByCookie(value string) (string, error) {
	filter := bson.D{{"session_token", value}}

	var result models.Cookie
	err := ar.db.Collection("cookie").FindOne(context.TODO(), filter).Decode(result)

	if err != nil {
		return "", errors.Wrap(err, "database error (table cookie)")
	}

	//if result.MaxAge.Before(time.Now()) {
	//	err := ar.DeleteCookie(value)
	//
	//	if err != nil {
	//		return "", errors.Wrap(err, "database error (table cookie)")
	//	}
	//
	//	return "", models.ErrNotFound
	//}

	return strconv.Itoa(int(result.UserID)), nil
}

func (ar authRepository) DeleteCookie(value string) error {
	filter := bson.D{{"session_token", value}}
	_, err := ar.db.Collection("cookie").DeleteOne(context.TODO(), filter)

	if err != nil {
		return errors.Wrap(err, "database error (table cookie)")
	}

	return nil
}
