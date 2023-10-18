package usecase_test

import (
	goalMocks "github.com/SerafimKuzmin/sd/backend/internal/Person/repository/mocks"
	"github.com/SerafimKuzmin/sd/backend/internal/Person/usecase"
	"github.com/SerafimKuzmin/sd/backend/models"
	"github.com/bxcodec/faker"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

type TestCaseGetPerson struct {
	ArgData     uint64
	ExpectedRes *models.Person
	Error       error
}

type TestCaseDeletePerson struct {
	ArgData []*uint64
	Error   error
}

type TestCaseCreateUpdatePerson struct {
	ArgData *models.Person
	Error   error
}

type TestCaseGetUserPersons struct {
	ArgData     *uint64
	ExpectedRes []*models.Person
	Error       error
}

func TestUsecaseGetPerson(t *testing.T) {
	var mockPersonRes models.Person
	err := faker.FakeData(&mockPersonRes)
	assert.NoError(t, err)

	mockExpectedPerson := mockPersonRes

	mockPersonRepo := goalMocks.NewRepositoryI(t)

	mockPersonRepo.On("GetPerson", mockPersonRes.ID).Return(&mockPersonRes, nil)

	useCase := usecase.New(mockPersonRepo, nil)

	cases := map[string]TestCaseGetPerson{
		"success": {
			ArgData:     mockPersonRes.ID,
			ExpectedRes: &mockExpectedPerson,
			Error:       nil,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			user, err := useCase.GetPerson(test.ArgData)
			require.Equal(t, test.Error, err)

			if err == nil {
				assert.Equal(t, test.ExpectedRes, user)
			}
		})
	}
	mockPersonRepo.AssertExpectations(t)
}

func TestUsecaseUpdatePerson(t *testing.T) {
	var mockPerson, invalidMockPerson models.Person
	err := faker.FakeData(&mockPerson)
	assert.NoError(t, err)

	invalidMockPerson.ID += mockPerson.ID + 1

	mockPersonRepo := goalMocks.NewRepositoryI(t)

	mockPersonRepo.On("GetPerson", mockPerson.ID).Return(&mockPerson, nil)
	mockPersonRepo.On("UpdatePerson", &mockPerson).Return(nil)

	mockPersonRepo.On("GetPerson", invalidMockPerson.ID).Return(nil, models.ErrNotFound)

	useCase := usecase.New(mockPersonRepo, nil)

	cases := map[string]TestCaseCreateUpdatePerson{
		"success": {
			ArgData: &mockPerson,
			Error:   nil,
		},
		"Person not found": {
			ArgData: &invalidMockPerson,
			Error:   models.ErrNotFound,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			err := useCase.UpdatePerson(test.ArgData)
			require.Equal(t, test.Error, errors.Cause(err))
		})
	}
	mockPersonRepo.AssertExpectations(t)
}

func TestUsecaseCreatePerson(t *testing.T) {
	var mockPerson models.Person
	err := faker.FakeData(&mockPerson)
	assert.NoError(t, err)

	mockPersonRepo := goalMocks.NewRepositoryI(t)

	mockPersonRepo.On("CreatePerson", &mockPerson).Return(nil)

	useCase := usecase.New(mockPersonRepo, nil)

	cases := map[string]TestCaseCreateUpdatePerson{
		"success": {
			ArgData: &mockPerson,
			Error:   nil,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			err := useCase.CreatePerson(test.ArgData)
			require.Equal(t, test.Error, errors.Cause(err))
		})
	}
	mockPersonRepo.AssertExpectations(t)
}

//func TestUsecaseDeletePerson(t *testing.T) {
//	var mockPerson, invalidMockPerson models.Person
//	err := faker.FakeData(&mockPerson)
//	assert.NoError(t, err)
//
//	invalidMockPerson.ID += mockPerson.ID + 1
//	*invalidMockPerson.UserID += *mockPerson.UserID + 1
//
//	mockPersonRepo := goalMocks.NewRepositoryI(t)
//
//	mockPersonRepo.On("GetPerson", mockPerson.ID).Return(&mockPerson, nil)
//	mockPersonRepo.On("DeletePerson", mockPerson.ID).Return(nil)
//
//	mockPersonRepo.On("GetPerson", invalidMockPerson.ID).Return(nil, models.ErrNotFound)
//
//	useCase := usecase.New(mockPersonRepo, nil)
//
//	cases := map[string]TestCaseDeletePerson{
//		"success": {
//			ArgData: []*uint64{&mockPerson.ID, mockPerson.UserID},
//			Error:   nil,
//		},
//		"Person not found": {
//			ArgData: []*uint64{&invalidMockPerson.ID, invalidMockPerson.UserID},
//			Error:   models.ErrNotFound,
//		},
//	}
//
//	for name, test := range cases {
//		t.Run(name, func(t *testing.T) {
//			err := useCase.DeletePerson(*test.ArgData[0], *test.ArgData[1])
//			require.Equal(t, test.Error, errors.Cause(err))
//		})
//	}
//	mockPersonRepo.AssertExpectations(t)
//}

//func TestUsecaseGetUserPersons(t *testing.T) {
//	mockPersonRes := make([]*models.Person, 0, 10)
//	err := faker.FakeData(&mockPersonRes)
//
//	for idx := range mockPersonRes {
//		mockPersonRes[idx].UserID = mockPersonRes[0].UserID
//	}
//	assert.NoError(t, err)
//
//	mockExpectedPerson := mockPersonRes
//
//	mockPersonRepo := goalMocks.NewRepositoryI(t)
//
//	mockPersonRepo.On("GetUserPersons", mockPersonRes[0].UserID).Return(mockPersonRes, nil)
//
//	useCase := usecase.New(mockPersonRepo, nil)
//
//	cases := map[string]TestCaseGetUserPersons{
//		"success": {
//			ArgData:     mockPersonRes[0].UserID,
//			ExpectedRes: mockExpectedPerson,
//			Error:       nil,
//		},
//	}
//
//	for name, test := range cases {
//		t.Run(name, func(t *testing.T) {
//			user, err := useCase.GetUserPersons(*test.ArgData)
//			require.Equal(t, test.Error, err)
//
//			if err == nil {
//				assert.Equal(t, test.ExpectedRes, user)
//			}
//		})
//	}
//	mockPersonRepo.AssertExpectations(t)
//}
