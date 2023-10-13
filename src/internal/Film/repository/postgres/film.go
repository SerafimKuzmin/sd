package postgres

import (
	"github.com/SerafimKuzmin/sd/src/internal/Film/repository"
	"github.com/SerafimKuzmin/sd/src/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

type Film struct {
	ID          uint64    `gorm:"column:modie_id"`
	Name        string    `gorm:"column:name"`
	Description string    `gorm:"column:description"`
	Rate        float64   `gorm:"column:rate"`
	Genre       string    `gorm:"column:genre"`
	ReleaseDT   time.Time `gorm:"column:release_dt"`
	Duration    uint      `gorm:"column:duration"`
	CountryID   uint64    `gorm:"column:country_id"`
}

func (Film) TableName() string {
	return "movie"
}

type FilmPerson struct {
	FilmID   uint64 `gorm:"column:movie_id"`
	PersonID uint64 `gorm:"column:member_id"`
}

func (FilmPerson) TableName() string {
	return "moviemember"
}

func toPostgresFilm(g *models.Film) *Film {
	return &Film{
		ID:          g.ID,
		Name:        g.Name,
		Description: g.Description,
		Rate:        g.Rate,
		Genre:       g.Genre,
		ReleaseDT:   g.ReleaseDT,
		Duration:    g.Duration,
		CountryID:   g.CountryID,
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
		CountryID:   g.CountryID,
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
	db *gorm.DB
}

func (gr FilmRepository) CreateFilm(g *models.Film) error {
	postgresFilm := toPostgresFilm(g)

	tx := gr.db.Create(postgresFilm)

	if tx.Error != nil {
		return errors.Wrap(tx.Error, "database error (table Film)")
	}

	g.ID = postgresFilm.ID
	return nil
}

func (gr FilmRepository) UpdateFilm(g *models.Film) error {
	postgresFilm := toPostgresFilm(g)

	tx := gr.db.Omit("id").Updates(postgresFilm)

	if tx.Error != nil {
		return errors.Wrap(tx.Error, "database error (table Film)")
	}

	return nil
}

func (gr FilmRepository) GetFilm(id uint64) (*models.Film, error) {
	var Film Film

	tx := gr.db.Where("id = ?", id).Take(&Film)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, models.ErrNotFound
	} else if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "database error (table Film)")
	}

	return toModelFilm(&Film), nil
}

func (gr FilmRepository) DeleteFilm(id uint64) error {
	tx := gr.db.Delete(&Film{}, id)

	if tx.Error != nil {
		return errors.Wrap(tx.Error, "database error (table Film)")
	}

	return nil
}

func (gr FilmRepository) GetFilmByPerson(id uint64) ([]*models.Film, error) {
	listFilmRels := make([]*FilmPerson, 0, 10)
	tx := gr.db.Where(&FilmPerson{PersonID: id}).Find(&listFilmRels)

	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "database error (table List)")
	}

	films := make([]*Film, 0, 10)

	for idx := range listFilmRels {
		var List Film
		tx := gr.db.Where(&Film{ID: listFilmRels[idx].FilmID}).Take(&List)

		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, models.ErrNotFound
		} else if tx.Error != nil {
			return nil, errors.Wrap(tx.Error, "database error (table List)")
		}

		films = append(films, &List)
	}

	return toModelFilms(films), nil
}

func (gr FilmRepository) GetFilmByCountry(id uint64) ([]*models.Film, error) {
	films := make([]*Film, 0, 10)

	tx := gr.db.Where(&Film{CountryID: id}).Find(&films)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, models.ErrNotFound
	} else if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "database error (table tag)")
	}

	return toModelFilms(films), nil
}

func NewFilmRepository(db *gorm.DB) repository.RepositoryI {
	return &FilmRepository{
		db: db,
	}
}
