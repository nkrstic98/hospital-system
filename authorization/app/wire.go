//go:build wireinject
// +build wireinject

package app

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	consul_api "github.com/hashicorp/consul/api"
	"gorm.io/gorm"
	actor_repo "hospital-system/authorization/app/repositories/actor"
	resource_repo "hospital-system/authorization/app/repositories/resource"
	role_repo "hospital-system/authorization/app/repositories/role"
	team_repo "hospital-system/authorization/app/repositories/team"
	"hospital-system/authorization/app/rpc"
	actor_service "hospital-system/authorization/app/services/actor"
	resource_service "hospital-system/authorization/app/services/resource"
	team_service "hospital-system/authorization/app/services/team"
	"hospital-system/authorization/config"
	"hospital-system/authorization/db"
	"hospital-system/proto_gen/authorization/v1"
)

//func Build(cfg config.Config) (*gin.Engine, func(), error) {
//	panic(wire.Build(
//		wire.Bind(new(registry.Registrar), new(*consul.Registry)),
//		provideRegistry,
//
//		db.OpenConnection,
//		buildAPI,
//		initializeApp,
//	))
//}

func Build(cfg config.Config) (*kratos.App, func(), error) {
	panic(wire.Build(
		wire.Bind(new(registry.Registrar), new(*consul.Registry)),
		provideRegistry,

		db.OpenConnection,
		buildService,
		initializeApp,
	))
}

//func buildAPI(db *gorm.DB) *api_v1.API {
//	panic(wire.Build(
//		wire.Bind(new(actor_repo.Repository), new(*actor_repo.RepositoryImpl)), actor_repo.NewRepository,
//		wire.Bind(new(role_repo.Repository), new(*role_repo.RepositoryImpl)), role_repo.NewRepository,
//		wire.Bind(new(team_repo.Repository), new(*team_repo.RepositoryImpl)), team_repo.NewRepository,
//		wire.Bind(new(actor_service.Service), new(*actor_service.ServiceImpl)), actor_service.NewService,
//		wire.Bind(new(actors_handlers.Handler), new(*actors_handlers.HandlerImpl)), actors_handlers.NewHandler,
//		api_v1.NewAPI,
//	))
//}

func buildService(db *gorm.DB) *rpc.Service {
	panic(wire.Build(
		wire.Bind(new(actor_repo.Repository), new(*actor_repo.RepositoryImpl)), actor_repo.NewRepository,
		wire.Bind(new(resource_repo.Repository), new(*resource_repo.RepositoryImpl)), resource_repo.NewRepository,
		wire.Bind(new(role_repo.Repository), new(*role_repo.RepositoryImpl)), role_repo.NewRepository,
		wire.Bind(new(team_repo.Repository), new(*team_repo.RepositoryImpl)), team_repo.NewRepository,
		wire.Bind(new(actor_service.Service), new(*actor_service.ServiceImpl)), actor_service.NewService,
		wire.Bind(new(resource_service.Service), new(*resource_service.ServiceImpl)), resource_service.NewService,
		wire.Bind(new(team_service.Service), new(*team_service.ServiceImpl)), team_service.NewService,
		rpc.NewService,
	))
}

//func initializeApp(api *api_v1.API) *gin.Engine {
//	router := gin.Default()
//	api.RegisterRoutes(router)
//	return router
//}

func initializeApp(cfg config.Config, service *rpc.Service, registrar registry.Registrar) (*kratos.App, error) {
	server := provideServer(cfg)
	authorization.RegisterAuthorizationServiceServer(server, service)

	app := kratos.New(
		kratos.Name(cfg.Registry.ServiceName),
		kratos.Server(
			server,
		),
		kratos.Registrar(registrar),
		kratos.AfterStop(func(ctx context.Context) error {
			return db.CloseConnection()
		}),
	)

	return app, nil
}

func provideServer(cfg config.Config) *grpc.Server {
	return grpc.NewServer(
		grpc.Address(fmt.Sprintf(":%s", cfg.Web.Port)),
		grpc.Timeout(cfg.Web.Timeout),
		grpc.Middleware(
			recovery.Recovery(),
			validate.Validator(),
		),
	)
}

func provideRegistry(cfg config.Config) *consul.Registry {
	consulConfig := consul_api.DefaultConfig()
	consulConfig.Address = fmt.Sprintf("%s:%v", cfg.Registry.RegistrarHost, cfg.Registry.RegistrarPort)
	c, err := consul_api.NewClient(consulConfig)
	if err != nil {
		panic(fmt.Sprintf("Failed to init Consul client. Error:\n%v", err))
	}
	return consul.New(c)
}
