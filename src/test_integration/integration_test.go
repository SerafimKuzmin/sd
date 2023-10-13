package testintegration

import (
	entryRep "github.com/SerafimKuzmin/sd/src/internal/Entry/repository/postgres"
	entryUsecase "github.com/SerafimKuzmin/sd/src/internal/Entry/usecase"
	projectRep "github.com/SerafimKuzmin/sd/src/internal/Project/repository/postgres"
	projectUsecase "github.com/SerafimKuzmin/sd/src/internal/Project/usecase"
	tagRep "github.com/SerafimKuzmin/sd/src/internal/Tag/repository/postgres"
	"github.com/SerafimKuzmin/sd/src/models"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	testCfg = postgres.Config{DSN: "host=localhost user=test password=test database=postgres port=13081"}
)

// in db
//INSERT INTO users (name, email, about, avatar_id, password)
//VALUES ('test', 'test', 'test', 0, '');

type UsecaseRepositoryTestSuite struct {
	suite.Suite
	db *gorm.DB
}

func (suite *UsecaseRepositoryTestSuite) SetupSuite() {
	db, err := gorm.Open(postgres.New(testCfg), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	suite.db = db
}

func (suite *UsecaseRepositoryTestSuite) TestUsecaseCreateProject() {
	_user_id := uint64(1)
	newProject := &models.Project{
		UserID:    &_user_id,
		Name:      "asdasd",
		About:     "asdasd",
		Color:     "aa",
		IsPrivate: true,
	}

	projectRepo := projectRep.NewProjectRepository(suite.db)
	useCase := projectUsecase.New(projectRepo)

	suite.Assert().NoError(useCase.CreateProject(newProject))

	id := newProject.ID // new id is created

	var result models.Project
	suite.Assert().NoError(suite.db.Table("project").First(&result, models.Project{ID: id}).Error)
	suite.Assert().Equal(newProject, &result)
}

func (suite *UsecaseRepositoryTestSuite) TestUsecaseGetProject() {
	_user_id := uint64(1)
	newProject := &models.Project{
		UserID:    &_user_id,
		Name:      "asdasd",
		About:     "asdasd",
		Color:     "aa",
		IsPrivate: true,
	}

	projectRepo := projectRep.NewProjectRepository(suite.db)
	useCase := projectUsecase.New(projectRepo)

	suite.Assert().NoError(useCase.CreateProject(newProject))

	id := newProject.ID // new id is created

	result, err := useCase.GetProject(id)
	suite.Assert().NoError(err)
	suite.Assert().Equal(newProject, result)
}

func (suite *UsecaseRepositoryTestSuite) TestUsecaseUpdateProject() {
	_user_id := uint64(1)
	newProject := &models.Project{
		UserID:    &_user_id,
		Name:      "asdasd",
		About:     "asdasd",
		Color:     "aa",
		IsPrivate: true,
	}

	projectRepo := projectRep.NewProjectRepository(suite.db)
	useCase := projectUsecase.New(projectRepo)

	suite.Assert().NoError(useCase.CreateProject(newProject))

	id := newProject.ID // new id is created

	newProjectUpdate := &models.Project{
		ID:        id,
		UserID:    &_user_id,
		Name:      "hello",
		About:     "asdasd",
		Color:     "aa",
		IsPrivate: true,
	}

	err := useCase.UpdateProject(newProjectUpdate)
	suite.Assert().NoError(err)

	result, err := useCase.GetProject(id)
	suite.Assert().NoError(err)
	suite.Assert().Equal(newProjectUpdate, result)
}

func (suite *UsecaseRepositoryTestSuite) TestUsecaseCreateEntry() {
	_user_id := uint64(1)
	newEntry := &models.Entry{
		UserID:      &_user_id,
		Description: "asdasd",
		TimeStart:   time.Now(),
		TimeEnd:     time.Now().Add(10),
	}

	entryRepo := entryRep.NewEntryRepository(suite.db)
	tagRepo := tagRep.NewTagRepository(suite.db)
	useCase := entryUsecase.New(entryRepo, tagRepo, nil)

	suite.Assert().NoError(useCase.CreateEntry(newEntry))

	id := newEntry.ID // new id is created

	var result entryRep.Entry
	suite.Assert().NoError(suite.db.Table("entry").First(&result, entryRep.Entry{ID: id}).Error)
	suite.Assert().Equal(newEntry.ID, result.ID)
	suite.Assert().Equal(newEntry.UserID, result.UserID)
	suite.Assert().Equal(newEntry.ProjectID, result.ProjectID)
	suite.Assert().Equal(newEntry.Description, result.Description)
}

func (suite *UsecaseRepositoryTestSuite) TestUsecaseGetEntry() {
	_user_id := uint64(1)
	newEntry := &models.Entry{
		UserID:      &_user_id,
		Description: "asdasd",
		TimeStart:   time.Now(),
		TimeEnd:     time.Now().Add(10),
		TagList:     nil,
	}

	entryRepo := entryRep.NewEntryRepository(suite.db)
	tagRepo := tagRep.NewTagRepository(suite.db)
	useCase := entryUsecase.New(entryRepo, tagRepo, nil)

	suite.Assert().NoError(useCase.CreateEntry(newEntry))

	id := newEntry.ID // new id is created

	result, err := useCase.GetEntry(id)
	suite.Assert().NoError(err)
	suite.Assert().Equal(newEntry.ID, result.ID)
	suite.Assert().Equal(newEntry.UserID, result.UserID)
	suite.Assert().Equal(newEntry.ProjectID, result.ProjectID)
	suite.Assert().Equal(newEntry.Description, result.Description)
}

func (suite *UsecaseRepositoryTestSuite) TestUsecaseUpdateEntry() {
	_user_id := uint64(1)
	newEntry := &models.Entry{
		UserID:      &_user_id,
		Description: "asdasd",
		TimeStart:   time.Now(),
		TimeEnd:     time.Now().Add(10),
		TagList:     nil,
	}

	entryRepo := entryRep.NewEntryRepository(suite.db)
	tagRepo := tagRep.NewTagRepository(suite.db)
	useCase := entryUsecase.New(entryRepo, tagRepo, nil)

	suite.Assert().NoError(useCase.CreateEntry(newEntry))

	id := newEntry.ID

	newEntryUpdated := &models.Entry{
		ID:          id,
		UserID:      &_user_id,
		Description: "a",
		TimeStart:   time.Now(),
		TimeEnd:     time.Now().Add(10),
		TagList:     nil,
	}

	err := useCase.UpdateEntry(newEntryUpdated)
	suite.Assert().NoError(err)

	result, err := useCase.GetEntry(id)
	suite.Assert().NoError(err)
	suite.Assert().Equal(newEntryUpdated.ID, result.ID)
	suite.Assert().Equal(newEntryUpdated.UserID, result.UserID)
	suite.Assert().Equal(newEntryUpdated.ProjectID, result.ProjectID)
	suite.Assert().Equal(newEntryUpdated.Description, result.Description)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(UsecaseRepositoryTestSuite))
}
