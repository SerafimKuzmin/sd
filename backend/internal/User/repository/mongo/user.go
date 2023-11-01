package mongo

import (
	"context"
	"github.com/SerafimKuzmin/sd/backend/internal/User/usecase"
	"github.com/SerafimKuzmin/sd/backend/models"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint64    `bson:"id"`
	Email     string    `bson:"email"`
	Login     string    `bson:"username"`
	Password  string    `bson:"password"`
	CreateDT  time.Time `bson:"create_dt"`
	CountryID *uint64   `json:"country_id"`
	Role      int       `bson:"role_id"`
	FullName  string    `gorm:"full_name"`
}

func (User) TableName() string {
	return "users"
}

func toClickhouseUser(u *models.User) *User {
	return &User{
		ID:        u.ID,
		Email:     u.Email,
		Login:     u.Login,
		Password:  u.Password,
		CreateDT:  u.CreateDT,
		CountryID: &u.CountryID,
		Role:      u.Role,
		FullName:  u.FullName,
	}
}

func toModelUser(u *User) *models.User {
	return &models.User{
		ID:        u.ID,
		Email:     u.Email,
		Login:     u.Login,
		Password:  u.Password,
		CreateDT:  u.CreateDT,
		CountryID: *u.CountryID,
		Role:      u.Role,
		FullName:  u.FullName,
	}
}

func toModelUsers(entries []*User) []*models.User {
	out := make([]*models.User, len(entries))

	for i, b := range entries {
		out[i] = toModelUser(b)
	}

	return out
}

type UserRepository struct {
	db *mongo.Database
}

func NewUserRepository(db *mongo.Database) usecase.RepositoryI {
	return &UserRepository{
		db: db,
	}
}

func (gr UserRepository) CreateUser(g *models.User) error {
	MongoUser := interface{}(bson.Marshal(g))
	_, err := gr.db.Collection("User").InsertOne(context.TODO(), MongoUser)

	if err != nil {
		return errors.Wrap(err, "database error (table User)")
	}
	return nil
}

func (gr UserRepository) UpdateUser(g *models.User) error {
	MongoUser := interface{}(bson.Marshal(g))
	_, err := gr.db.Collection("User").UpdateByID(context.TODO(), g.ID, MongoUser)

	if err != nil {
		return errors.Wrap(err, "database error (table User)")
	}

	return nil
}

func (gr UserRepository) GetUser(id uint64) (*models.User, error) {
	filter := bson.D{{"id", id}}

	var result models.User
	err := gr.db.Collection("User").FindOne(context.TODO(), filter).Decode(result)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, models.ErrNotFound
	} else if err != nil {
		return nil, errors.Wrap(err, "database error (table User)")
	}

	return &result, nil
}

func (gr UserRepository) DeleteUser(id uint64) error {
	filter := bson.D{{"id", id}}
	_, err := gr.db.Collection("User").DeleteOne(context.TODO(), filter)

	if err != nil {
		return errors.Wrap(err, "database error (table User)")
	}

	return nil
}

func (ur UserRepository) GetUsers() ([]*models.User, error) {
	filter := bson.D{}

	cursor, err := ur.db.Collection("user").Find(context.TODO(), filter)
	var results []*models.User

	if err == cursor.All(context.TODO(), &results) {
		return nil, models.ErrNotFound
	} else if err != nil {
		return nil, errors.Wrap(err, "database error (table user)")
	}

	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	return results, nil
}

func (ur UserRepository) GetUsersByIDs(userIDs []uint64) ([]*models.User, error) {
	filter := bson.D{{"id", userIDs}}

	cursor, err := ur.db.Collection("user").Find(context.TODO(), filter)
	var results []*models.User

	if err == cursor.All(context.TODO(), &results) {
		return nil, models.ErrNotFound
	} else if err != nil {
		return nil, errors.Wrap(err, "database error (table user)")
	}

	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	return results, nil
}

func (ur UserRepository) GetUserByEmail(email string) (*models.User, error) {
	filter := bson.D{{"email", email}}

	var result models.User
	err := ur.db.Collection("User").FindOne(context.TODO(), filter).Decode(result)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, models.ErrNotFound
	} else if err != nil {
		return nil, errors.Wrap(err, "database error (table User)")
	}

	return &result, nil
}
