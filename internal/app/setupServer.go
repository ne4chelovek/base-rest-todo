package app

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	authRepository "github.com/ne4chelovek/base-rest-todo/internal/repository/auth"
	itemRepository "github.com/ne4chelovek/base-rest-todo/internal/repository/item"
	listRepository "github.com/ne4chelovek/base-rest-todo/internal/repository/list"
	"github.com/ne4chelovek/base-rest-todo/internal/service"
	authService "github.com/ne4chelovek/base-rest-todo/internal/service/auth"
	itemService "github.com/ne4chelovek/base-rest-todo/internal/service/item"
	listService "github.com/ne4chelovek/base-rest-todo/internal/service/list"
	tokenService "github.com/ne4chelovek/base-rest-todo/internal/service/token"
	"github.com/ne4chelovek/base-rest-todo/pkg/handler"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"time"
)

type Server struct {
	HTTPServer *http.Server
	DB         *pgxpool.Pool
}

func SetupServer(ctx context.Context) (*Server, error) {
	pool, err := initDB(ctx)
	if err != nil {
		return nil, err
	}

	authSrv := createAuthService(pool)
	listSrv := createListService(pool)
	itemSrv := createItemService(pool)
	tokenSrv := createTokenService(pool)

	initHandlers := handler.NewHandler(authSrv, listSrv, itemSrv, tokenSrv)
	handlers := initHandlers.InitRouts()

	if err := initConfig(); err != nil {
		return nil, err
	}

	return &Server{
		HTTPServer: &http.Server{
			Addr:           ":" + viper.GetString("port"),
			MaxHeaderBytes: 1 << 20, // 1 MB
			Handler:        handlers,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
		},
		DB: pool,
	}, nil

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func initDB(ctx context.Context) (*pgxpool.Pool, error) {
	if err := godotenv.Load(); err != nil {
		logrus.Fatal("Error loading .env file")
	}

	conf := os.Getenv("DB_DSN")

	pool, err := pgxpool.New(ctx, conf)
	if err != nil {
		logrus.Errorf("failed to connect to database: %v", err)
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		logrus.Errorf("failed to ping database: %v", err)
		return nil, err
	}
	return pool, nil
}

func createAuthService(pool *pgxpool.Pool) service.Authorization {
	return authService.NewService(
		authRepository.NewAuthRepository(pool),
	)
}

func createListService(pool *pgxpool.Pool) service.TodoList {
	return listService.NewService(
		listRepository.NewListRepository(pool),
		pool,
	)
}

func createItemService(pool *pgxpool.Pool) service.TodoItem {
	return itemService.NewService(
		itemRepository.NewItemRepository(pool),
		listRepository.NewListRepository(pool),
		pool,
	)
}

func createTokenService(pool *pgxpool.Pool) service.Token {
	return tokenService.NewService(
		authRepository.NewAuthRepository(pool),
	)

}
