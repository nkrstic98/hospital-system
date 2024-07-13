//go:build wireinject
// +build wireinject

package app

import (
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/registry"
	kratosgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	consulapi "github.com/hashicorp/consul/api"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"hospital-system/proto_gen/authorization/v1"
	api_v1 "hospital-system/server/api/v1"
	admin_handler "hospital-system/server/api/v1/admin"
	session_handler "hospital-system/server/api/v1/session"
	admission_repo "hospital-system/server/app/repositories/admission"
	patient_repo "hospital-system/server/app/repositories/patient"
	user_repo "hospital-system/server/app/repositories/user"
	department_service "hospital-system/server/app/services/department"
	patient_service "hospital-system/server/app/services/patient"
	session_service "hospital-system/server/app/services/session"
	user_service "hospital-system/server/app/services/user"
	"hospital-system/server/config"
	"hospital-system/server/db"
	"time"
)

func Build(cfg config.Config) (*gin.Engine, func(), error) {
	panic(wire.Build(
		db.OpenConnection,
		provideRedisClient,
		provideRegistry,
		provideUserClient,
		buildAPI,
		initializeApp,
	))
}

func initializeApp(api *api_v1.API, cfg config.Config) *gin.Engine {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{cfg.Web.ClientAppUrl}
	config.AllowCredentials = true
	config.AllowHeaders = append(config.AllowHeaders, "Authorization")
	router.Use(cors.New(config))

	api.RegisterRoutes(router)

	return router
}

func buildAPI(
	db *gorm.DB,
	cfg config.Config,
	redisClient *redis.Client,
	userClient authorization.AuthorizationServiceClient,
) *api_v1.API {
	panic(wire.Build(
		wire.Bind(new(user_repo.Repository), new(*user_repo.RepositoryImpl)), user_repo.NewRepository,
		wire.Bind(new(patient_repo.Repository), new(*patient_repo.RepositoryImpl)), patient_repo.NewRepository,
		wire.Bind(new(admission_repo.Repository), new(*admission_repo.RepositoryImpl)), admission_repo.NewRepository,
		wire.Bind(new(user_service.Service), new(*user_service.ServiceImpl)), user_service.NewService,
		wire.Bind(new(session_service.Service), new(*session_service.ServiceImpl)), session_service.NewService,
		wire.Bind(new(patient_service.Service), new(*patient_service.ServiceImpl)), patient_service.NewService,
		wire.Bind(new(department_service.Service), new(*department_service.ServiceImpl)), department_service.NewService,
		wire.Bind(new(admin_handler.Handler), new(*admin_handler.HandlerImpl)), admin_handler.NewHandler,
		wire.Bind(new(session_handler.Handler), new(*session_handler.HandlerImpl)), session_handler.NewHandler,
		api_v1.NewAPI,
	))
}

func provideRedisClient(cfg config.Config) (*redis.Client, func()) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DatabaseIndex,
	})

	return client, func() {
		if err := client.Close(); err != nil {
			slog.Error(fmt.Sprintf("Failed to close Redis client. Error:\n%v", err))
			panic(err)
		}
	}
}

func provideUserClient(config config.Config, discovery registry.Discovery) (authorization.AuthorizationServiceClient, func()) {
	//conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	//if err != nil {
	//	panic(fmt.Sprintf("Failed to dial authorization service. Error:\n%v", err))
	//}

	conn := provideConnection("authorization", discovery)
	return authorization.NewAuthorizationServiceClient(conn), func() {
		if err := conn.Close(); err != nil {
			panic(fmt.Sprintf("Failed to close connection. Error:\n%v", err))
		}
	}
}

func provideConnection(serviceName string, discovery registry.Discovery) *grpc.ClientConn {
	conn, err := kratosgrpc.DialInsecure(
		context.Background(),
		kratosgrpc.WithEndpoint(fmt.Sprintf("discovery:///%s", serviceName)),
		kratosgrpc.WithDiscovery(discovery),
		kratosgrpc.WithTimeout(time.Second*0),
	)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to dial %s service. Error:\n%v", serviceName, err))
		return nil
	}

	return conn
}

func provideRegistry(cfg config.Config) registry.Discovery {
	consulConfig := consulapi.DefaultConfig()
	consulConfig.Address = fmt.Sprintf("%s:%v", cfg.Registry.RegistrarHost, cfg.Registry.RegistrarPort)
	client, err := consulapi.NewClient(consulConfig)
	if err != nil {
		panic(fmt.Sprintf("Failed to init Consul client. Error:\n%v", err))
	}

	return consul.New(client)
}
