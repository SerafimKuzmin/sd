package usecase_test

import (
	FilmMocks "github.com/SerafimKuzmin/sd/backend/internal/Film/repository/mocks"
	"github.com/SerafimKuzmin/sd/backend/internal/Film/usecase"
	"github.com/SerafimKuzmin/sd/backend/models"
	"github.com/bxcodec/faker"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

type TestCaseGetFilm struct {
	ArgData     uint64
	ExpectedRes *models.Film
	Error       error
}

type TestCaseDeleteFilm struct {
	ArgData []uint64
	Error   error
}

type TestCaseCreateUpdateFilm struct {
	ArgData *models.Film
	Error   error
}

type TestCaseGetUserFilms struct {
	ArgData     uint64
	ExpectedRes []*models.Film
	Error       error
}

func TestUsecaseGetFilm(t *testing.T) {
	var mockFilmRes models.Film
	err := faker.FakeData(&mockFilmRes)
	assert.NoError(t, err)

	mockExpectedFilm := mockFilmRes

	mockFilmRepo := FilmMocks.NewRepositoryI(t)

	mockFilmRepo.On("GetFilm", mockFilmRes.ID).Return(&mockFilmRes, nil)

	useCase := usecase.New(mockFilmRepo)

	cases := map[string]TestCaseGetFilm{
		"success": {
			ArgData:     mockFilmRes.ID,
			ExpectedRes: &mockExpectedFilm,
			Error:       nil,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			user, err := useCase.GetFilm(test.ArgData)
			require.Equal(t, test.Error, err)

			if err == nil {
				assert.Equal(t, test.ExpectedRes, user)
			}
		})
	}
	mockFilmRepo.AssertExpectations(t)
}

func TestUsecaseUpdateFilm(t *testing.T) {
	var mockFilm, invalidMockFilm models.Film
	err := faker.FakeData(&mockFilm)
	assert.NoError(t, err)

	invalidMockFilm.ID += mockFilm.ID + 1

	mockFilmRepo := FilmMocks.NewRepositoryI(t)

	mockFilmRepo.On("GetFilm", mockFilm.ID).Return(&mockFilm, nil)
	mockFilmRepo.On("UpdateFilm", &mockFilm).Return(nil)

	mockFilmRepo.On("GetFilm", invalidMockFilm.ID).Return(nil, models.ErrNotFound)

	useCase := usecase.New(mockFilmRepo)

	cases := map[string]TestCaseCreateUpdateFilm{
		"success": {
			ArgData: &mockFilm,
			Error:   nil,
		},
		"Film not found": {
			ArgData: &invalidMockFilm,
			Error:   models.ErrNotFound,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			err := useCase.UpdateFilm(test.ArgData)
			require.Equal(t, test.Error, errors.Cause(err))
		})
	}
	mockFilmRepo.AssertExpectations(t)
}

func TestUsecaseCreateFilm(t *testing.T) {
	var mockFilm models.Film
	err := faker.FakeData(&mockFilm)
	assert.NoError(t, err)

	mockFilmRepo := FilmMocks.NewRepositoryI(t)

	mockFilmRepo.On("CreateFilm", &mockFilm).Return(nil)

	useCase := usecase.New(mockFilmRepo)

	cases := map[string]TestCaseCreateUpdateFilm{
		"success": {
			ArgData: &mockFilm,
			Error:   nil,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			err := useCase.CreateFilm(test.ArgData)
			require.Equal(t, test.Error, errors.Cause(err))
		})
	}
	mockFilmRepo.AssertExpectations(t)
}
