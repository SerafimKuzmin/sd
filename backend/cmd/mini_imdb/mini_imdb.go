package mini_imdb

import (
	"fmt"
	"github.com/SerafimKuzmin/sd/backend/cmd/mini_imdb/flags"
	_authDelivery "github.com/SerafimKuzmin/sd/backend/internal/Auth/delivery"
	authRepPostgres "github.com/SerafimKuzmin/sd/backend/internal/Auth/repository/postgres"
	authRep "github.com/SerafimKuzmin/sd/backend/internal/Auth/repository/redis"
	authUsecase "github.com/SerafimKuzmin/sd/backend/internal/Auth/usecase"
	_countryDelivery "github.com/SerafimKuzmin/sd/backend/internal/Country/delivery"
	countryRep "github.com/SerafimKuzmin/sd/backend/internal/Country/repository/postgres"
	countryUsecase "github.com/SerafimKuzmin/sd/backend/internal/Country/usecase"

	_filmDelivery "github.com/SerafimKuzmin/sd/backend/internal/Film/delivery"
	filmRep "github.com/SerafimKuzmin/sd/backend/internal/Film/repository/postgres"
	filmUsecase "github.com/SerafimKuzmin/sd/backend/internal/Film/usecase"

	_listDelivery "github.com/SerafimKuzmin/sd/backend/internal/List/delivery"
	listRep "github.com/SerafimKuzmin/sd/backend/internal/List/repository/postgres"
	listUsecase "github.com/SerafimKuzmin/sd/backend/internal/List/usecase"

	_personDelivery "github.com/SerafimKuzmin/sd/backend/internal/Person/delivery"
	personRep "github.com/SerafimKuzmin/sd/backend/internal/Person/repository/postgres"
	personUsecase "github.com/SerafimKuzmin/sd/backend/internal/Person/usecase"

	_persononalRatingDelivery "github.com/SerafimKuzmin/sd/backend/internal/PersonalRating/delivery"
	persononalRatingRep "github.com/SerafimKuzmin/sd/backend/internal/PersonalRating/repository/postgres"
	persononalRatingUsecase "github.com/SerafimKuzmin/sd/backend/internal/PersonalRating/usecase"

	_userDelivery "github.com/SerafimKuzmin/sd/backend/internal/User/delivery"
	userRep "github.com/SerafimKuzmin/sd/backend/internal/User/repository/postgres"
	userUsecase "github.com/SerafimKuzmin/sd/backend/internal/User/usecase"
	"github.com/SerafimKuzmin/sd/backend/internal/cache"
	"github.com/SerafimKuzmin/sd/backend/internal/middleware"
	echo "github.com/labstack/echo/v4"

	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

type MiniIMDB struct {
	base
	PostgresClient           flags.PostgresFlags `toml:"postgres-client"`
	RedisSessionClient       flags.RedisFlags    `toml:"redis-client"`
	RedispersonStorageClient flags.RedisFlags    `toml:"redis-person-storage-client"`
	Server                   flags.ServerFlags   `toml:"server"`
}

func (mi MiniIMDB) Run(sessionDB string) error {
	e := echo.New()
	services, err := mi.Init(e)

	logger := services.Logger

	if err != nil {
		return fmt.Errorf("can not init services: %w", err)
	}

	postgresClient, err := mi.PostgresClient.Init()

	if err != nil {
		logger.Error("can not connect to Postgres client: %w", err)
		return err
	} else {
		logger.Info("Success conect to postgres")
	}

	redisSessionClient, err := mi.RedisSessionClient.Init()

	if err != nil {
		logger.Error("can not connect to Redis session client: %w", err)
		return err
	} else {
		logger.Info("Success conect to redis")
	}

	redisCacheClient, err := mi.RedispersonStorageClient.Init()

	if err != nil {
		logger.Error("can not connect to Redis cache client: %w", err)
		return err
	} else {
		logger.Info("Success conect to redis")
	}

	countryRepo := countryRep.NewCountryRepository(postgresClient)
	userRepo := userRep.NewUserRepository(postgresClient)
	persononalRatingRepo := persononalRatingRep.NewPersonalRatingRepository(postgresClient)
	filmRepo := filmRep.NewFilmRepository(postgresClient)
	personRepo := personRep.NewPersonRepository(postgresClient)
	authRepo := authRep.NewAuthRepository(redisSessionClient)
	authPostgresRepo := authRepPostgres.NewAuthRepositoryPostgres(postgresClient)
	listRepo := listRep.NewlistRepository(postgresClient)
	cacheStorage := cache.NewStorageRedis(redisCacheClient)

	countryUC := countryUsecase.New(countryRepo, cacheStorage)
	filmUC := filmUsecase.New(filmRepo)
	personUC := personUsecase.New(personRepo, cacheStorage)
	persononalRatingUC := persononalRatingUsecase.New(persononalRatingRepo)

	authUC := authUsecase.New(userRepo, authRepo)
	if sessionDB == "postgres" {
		authUC = authUsecase.New(userRepo, authPostgresRepo)
	}

	userUC := userUsecase.New(userRepo)
	listUC := listUsecase.New(listRepo)

	aclMiddleware := middleware.NewAclMiddleware(listUC)

	_countryDelivery.NewDelivery(e, countryUC, aclMiddleware)
	_filmDelivery.NewDelivery(e, filmUC, aclMiddleware)
	_personDelivery.NewDelivery(e, personUC, aclMiddleware)
	_persononalRatingDelivery.NewDelivery(e, persononalRatingUC, aclMiddleware)
	_authDelivery.NewDelivery(e, authUC)
	_userDelivery.NewDelivery(e, userUC, aclMiddleware)
	_listDelivery.NewDelivery(e, listUC, aclMiddleware)

	e.Use(echoMiddleware.LoggerWithConfig(echoMiddleware.LoggerConfig{
		Format: mi.Logger.LogHttpFormat,
		Output: logger.Output(),
	}))

	e.Use(echoMiddleware.Recover())
	authMiddleware := middleware.NewMiddleware(authUC)
	e.Use(authMiddleware.Auth)

	httpServer := mi.Server.Init(e)
	server := Server{*httpServer}
	if err := server.Start(); err != nil {
		logger.Fatal(err)
	}
	return nil
}
