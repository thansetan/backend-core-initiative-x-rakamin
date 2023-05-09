package main

import (
	"fmt"
	"log"
	"net/http"
	"self-payroll/config"
	"self-payroll/delivery"
	"self-payroll/repository"
	"self-payroll/usecase"
	"self-payroll/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	server struct {
		httpServer *echo.Echo
		cfg        config.Config
	}

	Server interface {
		Run()
	}
)

func initServer(cfg config.Config) Server {
	e := echo.New()
	e.HideBanner = true

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return &server{
		httpServer: e,
		cfg:        cfg,
	}
}

func (s *server) Run() {

	s.httpServer.GET("", func(e echo.Context) error {

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status":  "success",
			"message": "Hello, World!" + s.cfg.ServiceName() + " " + s.cfg.ServiceEnvironment(),
		})
	})

	positionRepo := repository.NewPositionRepository(s.cfg)
	positionUsecase := usecase.NewPositionUsecase(positionRepo)
	positionDelivery := delivery.NewPositionDelivery(positionUsecase)
	positionGroup := s.httpServer.Group("/positions", utils.AuthMiddleware)
	positionDelivery.Mount(positionGroup)

	companyRepo := repository.NewCompanyRepository(s.cfg)
	companyUsecase := usecase.NewCompanyUsecase(companyRepo)
	companyDelivery := delivery.NewCompanyDelivery(companyUsecase)
	companyGroup := s.httpServer.Group("/company", utils.AuthMiddleware, utils.CheckIsAdmin)
	companyDelivery.Mount(companyGroup)

	transactionRepo := repository.NewTransactionRepository(s.cfg)
	transactionUsecase := usecase.NewTransactionUsecase(transactionRepo)
	transactionDelivery := delivery.NewTransactionDelivery(transactionUsecase)
	transactionGroup := s.httpServer.Group("/transactions", utils.AuthMiddleware)
	transactionDelivery.Mount(transactionGroup)

	userRepo := repository.NewUserRepository(s.cfg)
	userUseCase := usecase.NewUserUsecase(userRepo, positionRepo, transactionRepo, companyRepo)
	userDelivery := delivery.NewUserDelivery(userUseCase)
	userGroup := s.httpServer.Group("/employee")
	userDelivery.Mount(userGroup)

	if err := s.httpServer.Start(fmt.Sprintf(":%d", s.cfg.ServicePort())); err != nil {
		log.Panic(err)
	}
}
