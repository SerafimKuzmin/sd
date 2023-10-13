package usecase_test

import (
	PersonalRatingMocks "github.com/SerafimKuzmin/sd/src/internal/PersonalRating/repository/mocks"
	"github.com/SerafimKuzmin/sd/src/internal/PersonalRating/usecase"
	"github.com/SerafimKuzmin/sd/src/models"
	"github.com/bxcodec/faker"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

type TestCaseGetPersonalRating struct {
	ArgData     uint64
	ExpectedRes *models.PersonalRating
	Error       error
}

type TestCaseDeletePersonalRating struct {
	ArgData []uint64
	Error   error
}

type TestCaseCreateUpdatePersonalRating struct {
	ArgData *models.PersonalRating
	Error   error
}

type TestCaseGetUserPersonalRatings struct {
	ArgData     uint64
	ExpectedRes []*models.PersonalRating
	Error       error
}

func TestUsecaseGetPersonalRating(t *testing.T) {
	var mockPersonalRatingRes models.PersonalRating
	err := faker.FakeData(&mockPersonalRatingRes)
	assert.NoError(t, err)

	mockExpectedPersonalRating := mockPersonalRatingRes

	mockPersonalRatingRepo := PersonalRatingMocks.NewRepositoryI(t)

	mockPersonalRatingRepo.On("GetPersonalRating", mockPersonalRatingRes.ID).Return(&mockPersonalRatingRes, nil)

	useCase := usecase.New(mockPersonalRatingRepo)

	cases := map[string]TestCaseGetPersonalRating{
		"success": {
			ArgData:     mockPersonalRatingRes.ID,
			ExpectedRes: &mockExpectedPersonalRating,
			Error:       nil,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			user, err := useCase.GetPersonalRating(test.ArgData)
			require.Equal(t, test.Error, err)

			if err == nil {
				assert.Equal(t, test.ExpectedRes, user)
			}
		})
	}
	mockPersonalRatingRepo.AssertExpectations(t)
}

func TestUsecaseUpdatePersonalRating(t *testing.T) {
	var mockPersonalRating, invalidMockPersonalRating models.PersonalRating
	err := faker.FakeData(&mockPersonalRating)
	assert.NoError(t, err)

	invalidMockPersonalRating.ID += mockPersonalRating.ID + 1

	mockPersonalRatingRepo := PersonalRatingMocks.NewRepositoryI(t)

	mockPersonalRatingRepo.On("GetPersonalRating", mockPersonalRating.ID).Return(&mockPersonalRating, nil)
	mockPersonalRatingRepo.On("UpdatePersonalRating", &mockPersonalRating).Return(nil)

	mockPersonalRatingRepo.On("GetPersonalRating", invalidMockPersonalRating.ID).Return(nil, models.ErrNotFound)

	useCase := usecase.New(mockPersonalRatingRepo)

	cases := map[string]TestCaseCreateUpdatePersonalRating{
		"success": {
			ArgData: &mockPersonalRating,
			Error:   nil,
		},
		"PersonalRating not found": {
			ArgData: &invalidMockPersonalRating,
			Error:   models.ErrNotFound,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			err := useCase.UpdatePersonalRating(test.ArgData)
			require.Equal(t, test.Error, errors.Cause(err))
		})
	}
	mockPersonalRatingRepo.AssertExpectations(t)
}

func TestUsecaseCreatePersonalRating(t *testing.T) {
	var mockPersonalRating models.PersonalRating
	err := faker.FakeData(&mockPersonalRating)
	assert.NoError(t, err)

	mockPersonalRatingRepo := PersonalRatingMocks.NewRepositoryI(t)

	mockPersonalRatingRepo.On("CreatePersonalRating", &mockPersonalRating).Return(nil)

	useCase := usecase.New(mockPersonalRatingRepo)

	cases := map[string]TestCaseCreateUpdatePersonalRating{
		"success": {
			ArgData: &mockPersonalRating,
			Error:   nil,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			err := useCase.CreatePersonalRating(test.ArgData)
			require.Equal(t, test.Error, errors.Cause(err))
		})
	}
	mockPersonalRatingRepo.AssertExpectations(t)
}

func TestUsecaseDeletePersonalRating(t *testing.T) {
	var mockPersonalRating, invalidMockPersonalRating models.PersonalRating
	err := faker.FakeData(&mockPersonalRating)
	assert.NoError(t, err)

	invalidMockPersonalRating.ID += mockPersonalRating.ID + 1
	//invalidMockPersonalRating.UserID += mockPersonalRating.UserID + 1

	mockPersonalRatingRepo := PersonalRatingMocks.NewRepositoryI(t)

	mockPersonalRatingRepo.On("GetPersonalRating", mockPersonalRating.ID).Return(&mockPersonalRating, nil)
	mockPersonalRatingRepo.On("DeletePersonalRating", mockPersonalRating.ID).Return(nil)

	mockPersonalRatingRepo.On("GetPersonalRating", invalidMockPersonalRating.ID).Return(nil, models.ErrNotFound)

	//useCase := usecase.New(mockPersonalRatingRepo)

	//cases := map[string]TestCaseDeletePersonalRating{
	//	"success": {
	//		ArgData: []uint64{mockPersonalRating.ID, mockPersonalRating.UserID},
	//		Error:   nil,
	//	},
	//	"PersonalRating not found": {
	//		ArgData: []uint64{invalidMockPersonalRating.ID, invalidMockPersonalRating.UserID},
	//		Error:   models.ErrNotFound,
	//	},
	//}
	//
	//for name, test := range cases {
	//	t.Run(name, func(t *testing.T) {
	//		err := useCase.DeletePersonalRating(test.ArgData[0], test.ArgData[1])
	//		require.Equal(t, test.Error, errors.Cause(err))
	//	})
	//}
	//mockPersonalRatingRepo.AssertExpectations(t)
}

func TestUsecaseGetUserPersonalRatings(t *testing.T) {
	mockPersonalRatingRes := make([]*models.PersonalRating, 0, 10)
	err := faker.FakeData(&mockPersonalRatingRes)

	for idx := range mockPersonalRatingRes {
		mockPersonalRatingRes[idx].UserID = mockPersonalRatingRes[0].UserID
	}
	assert.NoError(t, err)

	//mockExpectedPersonalRating := mockPersonalRatingRes

	mockPersonalRatingRepo := PersonalRatingMocks.NewRepositoryI(t)

	mockPersonalRatingRepo.On("GetUserPersonalRatings", mockPersonalRatingRes[0].UserID).Return(mockPersonalRatingRes, nil)

	//useCase := usecase.New(mockPersonalRatingRepo)
	//
	//cases := map[string]TestCaseGetUserPersonalRatings{
	//	"success": {
	//		ArgData:     mockPersonalRatingRes[0].UserID,
	//		ExpectedRes: mockExpectedPersonalRating,
	//		Error:       nil,
	//	},
	//}
	//
	//for name, test := range cases {
	//	t.Run(name, func(t *testing.T) {
	//		user, err := useCase.GetUserPersonalRatings(test.ArgData)
	//		require.Equal(t, test.Error, err)
	//
	//		if err == nil {
	//			assert.Equal(t, test.ExpectedRes, user)
	//		}
	//	})
	//}
	//mockPersonalRatingRepo.AssertExpectations(t)
}
