package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/mondayy1/brainrot-battle/internal/config"
	"github.com/mondayy1/brainrot-battle/internal/db"
	"github.com/mondayy1/brainrot-battle/internal/handler"
	repository "github.com/mondayy1/brainrot-battle/internal/repository/postgres"
	"github.com/mondayy1/brainrot-battle/internal/usecase"
)

func main() {
	// YAML Load
	env := flag.String("env", "dev", "Environment (dev, prod)")
	flag.Parse()

	validEnvs := map[string]bool{"dev": true, "prod": true}
	if !validEnvs[*env] {
		log.Fatalf("Invalid environment: %s", *env)
	}

	cfg, err := config.LoadConfig(*env)
	if err != nil {
		log.Fatalf("Config load error: %v", err)
	}

	fmt.Printf("config: %+v\n", cfg)

	//DB
	dbConn, err := db.NewDBConnection(cfg)
	if err != nil {
		log.Fatalf("DB connection error: %v", err)
	}

	userRepo := repository.NewUserRepository(dbConn)
	userUseCase := usecase.NewUserUseCase(userRepo)
	userHandler := handler.NewUserHandler(userUseCase)

	//Server
	e := echo.New()
	//e.POST("/users")
	e.GET("/users", userHandler.GetAll)

	log.Printf("Server started at %s", fmt.Sprintf(":%d", cfg.App.Port))
	if err := e.Start(fmt.Sprintf(":%d", cfg.App.Port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
