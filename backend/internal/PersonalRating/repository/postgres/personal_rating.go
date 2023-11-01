package postgres

import (
	"github.com/SerafimKuzmin/sd/backend/internal/PersonalRating/usecase"
	"github.com/SerafimKuzmin/sd/backend/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type PersonalRating struct {
	ID     uint64  `gorm:"column:rating_id"`
	UserID uint64  `gorm:"column:user_id"`
	FilmID uint64  `gorm:"column:movie_id"`
	Rate   float64 `gorm:"column:user_rating"`
}

func (PersonalRating) TableName() string {
	return "rating"
}

func toPostgresPersonalRating(t *models.PersonalRating) *PersonalRating {
	return &PersonalRating{
		ID:     t.ID,
		UserID: t.UserID,
		FilmID: t.FilmID,
		Rate:   t.Rate,
	}
}

func toModelPersonalRating(t *PersonalRating) *models.PersonalRating {
	return &models.PersonalRating{
		ID:     t.ID,
		UserID: t.UserID,
		FilmID: t.FilmID,
		Rate:   t.Rate,
	}
}

func toModelPersonalRatings(PersonalRatings []*PersonalRating) []*models.PersonalRating {
	out := make([]*models.PersonalRating, len(PersonalRatings))

	for i, b := range PersonalRatings {
		out[i] = toModelPersonalRating(b)
	}

	return out
}

type PersonalRatingRepository struct {
	db *gorm.DB
}

func (tr PersonalRatingRepository) CreatePersonalRating(t *models.PersonalRating) error {
	postgresPersonalRating := toPostgresPersonalRating(t)

	tx := tr.db.Create(postgresPersonalRating)

	if tx.Error != nil {
		return errors.Wrap(tx.Error, "database error (table PersonalRating)")
	}

	t.ID = postgresPersonalRating.ID
	return nil
}

func (tr PersonalRatingRepository) UpdatePersonalRating(t *models.PersonalRating) error {
	postgresPersonalRating := toPostgresPersonalRating(t)

	tx := tr.db.Omit("id").Updates(postgresPersonalRating)

	if tx.Error != nil {
		return errors.Wrap(tx.Error, "database error (table PersonalRating)")
	}

	return nil
}

func (tr PersonalRatingRepository) GetPersonalRating(id uint64) (*models.PersonalRating, error) {
	var PersonalRating PersonalRating

	tx := tr.db.Where("id = ?", id).Take(&PersonalRating)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, models.ErrNotFound
	} else if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "database error (table PersonalRating)")
	}

	return toModelPersonalRating(&PersonalRating), nil
}

func (tr PersonalRatingRepository) DeletePersonalRating(id uint64) error {
	tx := tr.db.Delete(&PersonalRating{}, id)

	if tx.Error != nil {
		return errors.Wrap(tx.Error, "database error (table PersonalRating)")
	}

	return nil
}

func NewPersonalRatingRepository(db *gorm.DB) usecase.RepositoryI {
	return &PersonalRatingRepository{
		db: db,
	}
}
