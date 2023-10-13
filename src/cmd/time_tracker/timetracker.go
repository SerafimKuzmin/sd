package time_tracker

import (
	"fmt"
	"github.com/SerafimKuzmin/sd/src/cmd/time_tracker/flags"
	_authDelivery "github.com/SerafimKuzmin/sd/src/internal/Auth/delivery"
	authRepPostgres "github.com/SerafimKuzmin/sd/src/internal/Auth/repository/postgres"
	authRep "github.com/SerafimKuzmin/sd/src/internal/Auth/repository/redis"
	authUsecase "github.com/SerafimKuzmin/sd/src/internal/Auth/usecase"
	_entryDelivery "github.com/SerafimKuzmin/sd/src/internal/Entry/delivery"
	entryRep "github.com/SerafimKuzmin/sd/src/internal/Entry/repository/postgres"
	entryUsecase "github.com/SerafimKuzmin/sd/src/internal/Entry/usecase"
	_friendDelivery "github.com/SerafimKuzmin/sd/src/internal/Friends/delivery"
	friendRep "github.com/SerafimKuzmin/sd/src/internal/Friends/repository/postgres"
	friendUsecase "github.com/SerafimKuzmin/sd/src/internal/Friends/usecase"
	_goalDelivery "github.com/SerafimKuzmin/sd/src/internal/Goal/delivery"
	goalRep "github.com/SerafimKuzmin/sd/src/internal/Goal/repository/postgres"
	goalUsecase "github.com/SerafimKuzmin/sd/src/internal/Goal/usecase"
	_projectDelivery "github.com/SerafimKuzmin/sd/src/internal/Project/delivery"
	projectRep "github.com/SerafimKuzmin/sd/src/internal/Project/repository/postgres"
	projectUsecase "github.com/SerafimKuzmin/sd/src/internal/Project/usecase"
	_tagDelivery "github.com/SerafimKuzmin/sd/src/internal/Tag/delivery"
	tagRep "github.com/SerafimKuzmin/sd/src/internal/Tag/repository/postgres"
	tagUsecase "github.com/SerafimKuzmin/sd/src/internal/Tag/usecase"
	_userDelivery "github.com/SerafimKuzmin/sd/src/internal/User/delivery"
	userRep "github.com/SerafimKuzmin/sd/src/internal/User/repository/postgres"
	userUsecase "github.com/SerafimKuzmin/sd/src/internal/User/usecase"
	"github.com/SerafimKuzmin/sd/src/internal/cache"
	"github.com/SerafimKuzmin/sd/src/internal/middleware"
	echo "github.com/labstack/echo/v4"

	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

type TimeTracker struct {
	base
	PostgresClient            flags.PostgresFlags `toml:"postgres-client"`
	RedisSessionClient        flags.RedisFlags    `toml:"redis-client"`
	RedisProjectStorageClient flags.RedisFlags    `toml:"redis-project-storage-client"`
	Server                    flags.ServerFlags   `toml:"server"`
}

func (tt TimeTracker) Run(sessionDB string) error {
	e := echo.New()
	services, err := tt.Init(e)

	logger := services.Logger

	if err != nil {
		return fmt.Errorf("can not init services: %w", err)
	}

	postgresClient, err := tt.PostgresClient.Init()

	if err != nil {
		logger.Error("can not connect to Postgres client: %w", err)
		return err
	} else {
		logger.Info("Success conect to postgres")
	}

	redisSessionClient, err := tt.RedisSessionClient.Init()

	if err != nil {
		logger.Error("can not connect to Redis session client: %w", err)
		return err
	} else {
		logger.Info("Success conect to redis")
	}

	redisCacheClient, err := tt.RedisProjectStorageClient.Init()

	if err != nil {
		logger.Error("can not connect to Redis cache client: %w", err)
		return err
	} else {
		logger.Info("Success conect to redis")
	}

	entryRepo := entryRep.NewEntryRepository(postgresClient)
	userRepo := userRep.NewUserRepository(postgresClient)
	tagRepo := tagRep.NewTagRepository(postgresClient)
	goalRepo := goalRep.NewGoalRepository(postgresClient)
	projectRepo := projectRep.NewProjectRepository(postgresClient)
	authRepo := authRep.NewAuthRepository(redisSessionClient)
	authPostgresRepo := authRepPostgres.NewAuthRepositoryPostgres(postgresClient)
	friendRepo := friendRep.NewFriendRepository(postgresClient)
	cacheStorage := cache.NewStorageRedis(redisCacheClient)

	entryUC := entryUsecase.New(entryRepo, tagRepo, userRepo)
	goalUC := goalUsecase.New(goalRepo)
	projectUC := projectUsecase.New(projectRepo, cacheStorage)
	tagUC := tagUsecase.New(tagRepo)

	authUC := authUsecase.New(userRepo, authRepo)
	if sessionDB == "postgres" {
		authUC = authUsecase.New(userRepo, authPostgresRepo)
	}

	userUC := userUsecase.New(userRepo)
	friendUC := friendUsecase.New(friendRepo, userRepo)

	aclMiddleware := middleware.NewAclMiddleware(friendUC)

	_entryDelivery.NewDelivery(e, entryUC, aclMiddleware)
	_goalDelivery.NewDelivery(e, goalUC, aclMiddleware)
	_projectDelivery.NewDelivery(e, projectUC, aclMiddleware)
	_tagDelivery.NewDelivery(e, tagUC, aclMiddleware)
	_authDelivery.NewDelivery(e, authUC)
	_userDelivery.NewDelivery(e, userUC, aclMiddleware)
	_friendDelivery.NewDelivery(e, friendUC, aclMiddleware)

	e.Use(echoMiddleware.LoggerWithConfig(echoMiddleware.LoggerConfig{
		Format: tt.Logger.LogHttpFormat,
		Output: logger.Output(),
	}))

	e.Use(echoMiddleware.Recover())
	authMiddleware := middleware.NewMiddleware(authUC)
	e.Use(authMiddleware.Auth)

	httpServer := tt.Server.Init(e)
	server := Server{*httpServer}
	if err := server.Start(); err != nil {
		logger.Fatal(err)
	}
	return nil
}
