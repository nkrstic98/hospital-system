package app

import (
	"context"
	"fmt"
	"hospital-system/proto_gen/authorization/v1"
	"hospital-system/server/app/handlers"
	"hospital-system/server/app/repositories"
	"hospital-system/server/app/services"
	"hospital-system/server/config"
	"hospital-system/server/db"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/registry"
	kratosgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-redis/redis/v8"
	consulapi "github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func Build(cfg config.Config) (*gin.Engine, func(), error) {
	dbConn, err := db.OpenConnection(cfg)
	if err != nil {
		return nil, nil, err
	}

	redisClient, redisClean := provideRedisClient(cfg)

	authServiceClient, authClientCleanup, err := provideAuthorizationClient(cfg)
	if err != nil {
		return nil, nil, err
	}

	logger, err := provideLogger()
	if err != nil {
		return nil, nil, err
	}

	repo := repositories.NewRepository(dbConn)

	userService := services.NewUserService(authServiceClient, repo)
	sessionService := services.NewSessionService(redisClient, cfg.AuthToken)
	patientService := services.NewPatientService(authServiceClient, repo, userService)

	handler := handlers.NewHandler(logger, userService, sessionService, patientService)

	router := gin.Default()
	routerConfig := cors.DefaultConfig()
	routerConfig.AllowOrigins = []string{cfg.Web.ClientAppUrl}
	routerConfig.AllowCredentials = true
	routerConfig.AllowHeaders = append(routerConfig.AllowHeaders, "Authorization")
	router.Use(cors.New(routerConfig))

	handler.RegisterRoutes(router)

	return router, func() {
		authClientCleanup()
		redisClean()
	}, nil
}

func provideLogger() (*zap.Logger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, fmt.Errorf("failed to create logger. %w", err)
	}

	logger.Sync()

	return logger, nil
}

func provideRedisClient(cfg config.Config) (*redis.Client, func()) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DatabaseIndex,
	})

	return client, func() {
		if err := client.Close(); err != nil {
			panic(fmt.Errorf("failed to close Redis client connection. %w", err))
		}
	}
}

func provideAuthorizationClient(cfg config.Config) (authorization.AuthorizationServiceClient, func(), error) {
	discovery, err := provideRegistry(cfg)
	if err != nil {
		return nil, nil, err
	}

	conn, err := provideConnection("authorization", discovery)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to provide connection. %w", err)
	}

	return authorization.NewAuthorizationServiceClient(conn), func() {
		if err := conn.Close(); err != nil {
			panic(fmt.Sprintf("Failed to close connection. Error:\n%v", err))
		}
	}, nil
}

func provideConnection(serviceName string, discovery registry.Discovery) (*grpc.ClientConn, error) {
	conn, err := kratosgrpc.DialInsecure(
		context.Background(),
		kratosgrpc.WithEndpoint(fmt.Sprintf("discovery:///%s", serviceName)),
		kratosgrpc.WithDiscovery(discovery),
		kratosgrpc.WithTimeout(time.Second*0),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to dial connection. %w", err)
	}

	return conn, nil
}

func provideRegistry(cfg config.Config) (registry.Discovery, error) {
	consulConfig := consulapi.DefaultConfig()
	consulConfig.Address = fmt.Sprintf("%s:%v", cfg.Registry.RegistrarHost, cfg.Registry.RegistrarPort)
	client, err := consulapi.NewClient(consulConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to init Consul client. Error:\n%v", err)
	}

	return consul.New(client), nil
}
