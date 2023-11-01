package postgres

import (
	"github.com/SerafimKuzmin/sd/backend/internal/List/usecase"
	"github.com/SerafimKuzmin/sd/backend/models"
	"time"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Film struct {
	ID          uint64    `gorm:"column:id"`
	Name        string    `gorm:"column:title"`
	Description string    `gorm:"column:overview"`
	Rate        float64   `gorm:"column:rating"`
	ReleaseDT   time.Time `gorm:"column:release_date"`
	Duration    uint      `gorm:"column:runtime"`
}

type List struct {
	ID       uint64    `gorm:"column:id"`
	Name     string    `gorm:"column:name"`
	CreateDT time.Time `gorm:"column:create_dt"`
}

type UserListRelation struct {
	UserID uint64 `gorm:"column:user_id"`
	ListID uint64 `gorm:"column:list_id"`
}

type ListFilmRelation struct {
	ID     uint64 `gorm:"column:id"`
	FilmID uint64 `gorm:"column:movie_id"`
	ListID uint64 `gorm:"column:list_id"`
}

func (List) TableName() string {
	return "list"
}

func (Film) TableName() string {
	return "movie"
}

func (UserListRelation) TableName() string {
	return "user_list"
}

func (ListFilmRelation) TableName() string {
	return "list_movie"
}

type listRepository struct {
	db *gorm.DB
}

func toPostgresList(t *models.List) *List {
	return &List{
		ID:       t.ID,
		Name:     t.Name,
		CreateDT: t.CreateDT,
	}
}

func toModelFilm(t *Film) *models.Film {
	return &models.Film{
		ID:          t.ID,
		Name:        t.Name,
		Description: t.Description,
		Rate:        t.Rate,
		ReleaseDT:   t.ReleaseDT,
		Duration:    t.Duration,
	}
}

func toModelList(t *List) *models.List {
	return &models.List{
		ID:       t.ID,
		Name:     t.Name,
		CreateDT: t.CreateDT,
	}
}

func toModelFilms(films []*Film) []*models.Film {
	out := make([]*models.Film, len(films))

	for i, b := range films {
		out[i] = toModelFilm(b)
	}

	return out
}

func toModelLists(lists []*List) []*models.List {
	out := make([]*models.List, len(lists))

	for i, b := range lists {
		out[i] = toModelList(b)
	}

	return out
}

func (fr listRepository) CreateList(t *models.List) error {
	postgreslist := toPostgresList(t)

	tx := fr.db.Create(postgreslist)

	if tx.Error != nil {
		return errors.Wrap(tx.Error, "database error (table list_relation)")
	}

	return nil
}

func (tr listRepository) GetList(id uint64) (*models.List, error) {
	var List List

	tx := tr.db.Where("id = ?", id).Take(&List)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, models.ErrNotFound
	} else if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "database error (table List)")
	}

	return toModelList(&List), nil
}

func (tr listRepository) UpdateList(t *models.List) error {
	postgresList := toPostgresList(t)

	tx := tr.db.Omit("id").Updates(postgresList)

	if tx.Error != nil {
		return errors.Wrap(tx.Error, "database error (table List)")
	}

	return nil
}

func (fr listRepository) DeleteList(listId uint64) error {
	tx := fr.db.Delete(&List{}, listId)

	if tx.Error != nil {
		return errors.Wrap(tx.Error, "database error (table tag)")
	}

	return nil
}

func (fr listRepository) AddFilm(listID uint64, filmID uint64) error {
	postgresFilm := &ListFilmRelation{
		ListID: listID,
		FilmID: filmID,
	}

	tx := fr.db.Create(postgresFilm)

	if tx.Error != nil {
		return errors.Wrap(tx.Error, "database error (table list_relation)")
	}
	return nil
}

func (fr listRepository) GetUserLists(userID uint64) ([]*models.List, error) {
	userListsRels := make([]*UserListRelation, 0, 10)
	tx := fr.db.Where(&UserListRelation{UserID: userID}).Find(&userListsRels)

	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "database error (table List)")
	}

	lists := make([]*List, 0, 10)

	for idx := range userListsRels {
		var list List
		tx := fr.db.Where(&List{ID: userListsRels[idx].ListID}).Take(&list)

		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, models.ErrNotFound
		} else if tx.Error != nil {
			return nil, errors.Wrap(tx.Error, "database error (table List)")
		}

		lists = append(lists, &list)
	}

	return toModelLists(lists), nil
}

func (fr listRepository) GetFilmsByList(listID uint64) ([]*models.Film, error) {
	listFilmRels := make([]*ListFilmRelation, 0, 10)
	tx := fr.db.Where(&ListFilmRelation{ListID: listID}).Find(&listFilmRels)

	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "database error (table List)")
	}

	films := make([]*Film, 0, 10)

	for idx := range listFilmRels {
		var film Film
		tx := fr.db.Where("id = ?", listFilmRels[idx].FilmID).Find(&film)

		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, models.ErrNotFound
		} else if tx.Error != nil {
			return nil, errors.Wrap(tx.Error, "database error (table film)")
		}

		films = append(films, &film)
	}

	return toModelFilms(films), nil
}

func NewlistRepository(db *gorm.DB) usecase.RepositoryI {
	return &listRepository{
		db: db,
	}
}
