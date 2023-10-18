package usecase_test

import (
	goalMocks "github.com/SerafimKuzmin/sd/backend/internal/Country/repository/mocks"
	"github.com/SerafimKuzmin/sd/backend/internal/Country/usecase"
	"github.com/SerafimKuzmin/sd/backend/models"
	"github.com/bxcodec/faker"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

type TestCaseGetCountry struct {
	ArgData     uint64
	ExpectedRes *models.Country
	Error       error
}

type TestCaseDeleteCountry struct {
	ArgData []*uint64
	Error   error
}

type TestCaseCreateUpdateCountry struct {
	ArgData *models.Country
	Error   error
}

type TestCaseGetUserCountrys struct {
	ArgData     *uint64
	ExpectedRes []*models.Country
	Error       error
}

func TestUsecaseGetCountry(t *testing.T) {
	var mockCountryRes models.Country
	err := faker.FakeData(&mockCountryRes)
	assert.NoError(t, err)

	mockExpectedCountry := mockCountryRes

	mockCountryRepo := goalMocks.NewRepositoryI(t)

	mockCountryRepo.On("GetCountry", mockCountryRes.ID).Return(&mockCountryRes, nil)

	useCase := usecase.New(mockCountryRepo, nil)

	cases := map[string]TestCaseGetCountry{
		"success": {
			ArgData:     mockCountryRes.ID,
			ExpectedRes: &mockExpectedCountry,
			Error:       nil,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			user, err := useCase.GetCountry(test.ArgData)
			require.Equal(t, test.Error, err)

			if err == nil {
				assert.Equal(t, test.ExpectedRes, user)
			}
		})
	}
	mockCountryRepo.AssertExpectations(t)
}

func TestUsecaseUpdateCountry(t *testing.T) {
	var mockCountry, invalidMockCountry models.Country
	err := faker.FakeData(&mockCountry)
	assert.NoError(t, err)

	invalidMockCountry.ID += mockCountry.ID + 1

	mockCountryRepo := goalMocks.NewRepositoryI(t)

	mockCountryRepo.On("GetCountry", mockCountry.ID).Return(&mockCountry, nil)
	mockCountryRepo.On("UpdateCountry", &mockCountry).Return(nil)

	mockCountryRepo.On("GetCountry", invalidMockCountry.ID).Return(nil, models.ErrNotFound)

	useCase := usecase.New(mockCountryRepo, nil)

	cases := map[string]TestCaseCreateUpdateCountry{
		"success": {
			ArgData: &mockCountry,
			Error:   nil,
		},
		"Country not found": {
			ArgData: &invalidMockCountry,
			Error:   models.ErrNotFound,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			err := useCase.UpdateCountry(test.ArgData)
			require.Equal(t, test.Error, errors.Cause(err))
		})
	}
	mockCountryRepo.AssertExpectations(t)
}

func TestUsecaseCreateCountry(t *testing.T) {
	var mockCountry models.Country
	err := faker.FakeData(&mockCountry)
	assert.NoError(t, err)

	mockCountryRepo := goalMocks.NewRepositoryI(t)

	mockCountryRepo.On("CreateCountry", &mockCountry).Return(nil)

	useCase := usecase.New(mockCountryRepo, nil)

	cases := map[string]TestCaseCreateUpdateCountry{
		"success": {
			ArgData: &mockCountry,
			Error:   nil,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			err := useCase.CreateCountry(test.ArgData)
			require.Equal(t, test.Error, errors.Cause(err))
		})
	}
	mockCountryRepo.AssertExpectations(t)
}

//func TestUsecaseDeleteCountry(t *testing.T) {
//	var mockCountry, invalidMockCountry models.Country
//	err := faker.FakeData(&mockCountry)
//	assert.NoError(t, err)
//
//	invalidMockCountry.ID += mockCountry.ID + 1
//	*invalidMockCountry.UserID += *mockCountry.UserID + 1
//
//	mockCountryRepo := goalMocks.NewRepositoryI(t)
//
//	mockCountryRepo.On("GetCountry", mockCountry.ID).Return(&mockCountry, nil)
//	mockCountryRepo.On("DeleteCountry", mockCountry.ID).Return(nil)
//
//	mockCountryRepo.On("GetCountry", invalidMockCountry.ID).Return(nil, models.ErrNotFound)
//
//	useCase := usecase.New(mockCountryRepo, nil)
//
//	cases := map[string]TestCaseDeleteCountry{
//		"success": {
//			ArgData: []*uint64{&mockCountry.ID, mockCountry.UserID},
//			Error:   nil,
//		},
//		"Country not found": {
//			ArgData: []*uint64{&invalidMockCountry.ID, invalidMockCountry.UserID},
//			Error:   models.ErrNotFound,
//		},
//	}
//
//	for name, test := range cases {
//		t.Run(name, func(t *testing.T) {
//			err := useCase.DeleteCountry(*test.ArgData[0], *test.ArgData[1])
//			require.Equal(t, test.Error, errors.Cause(err))
//		})
//	}
//	mockCountryRepo.AssertExpectations(t)
//}

//func TestUsecaseGetUserCountrys(t *testing.T) {
//	mockCountryRes := make([]*models.Country, 0, 10)
//	err := faker.FakeData(&mockCountryRes)
//
//	for idx := range mockCountryRes {
//		mockCountryRes[idx].UserID = mockCountryRes[0].UserID
//	}
//	assert.NoError(t, err)
//
//	mockExpectedCountry := mockCountryRes
//
//	mockCountryRepo := goalMocks.NewRepositoryI(t)
//
//	mockCountryRepo.On("GetUserCountrys", mockCountryRes[0].UserID).Return(mockCountryRes, nil)
//
//	useCase := usecase.New(mockCountryRepo, nil)
//
//	cases := map[string]TestCaseGetUserCountrys{
//		"success": {
//			ArgData:     mockCountryRes[0].UserID,
//			ExpectedRes: mockExpectedCountry,
//			Error:       nil,
//		},
//	}
//
//	for name, test := range cases {
//		t.Run(name, func(t *testing.T) {
//			user, err := useCase.GetUserCountrys(*test.ArgData)
//			require.Equal(t, test.Error, err)
//
//			if err == nil {
//				assert.Equal(t, test.ExpectedRes, user)
//			}
//		})
//	}
//	mockCountryRepo.AssertExpectations(t)
//}
