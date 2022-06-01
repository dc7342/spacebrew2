package main

import (
	"fmt"
	"github.com/je09/spacebrew2/internal/repository"
	"github.com/je09/spacebrew2/internal/service"
	"github.com/je09/spacebrew2/pkg/config"
)

func main() {
	c, err := config.Read()
	if err != nil {
		panic(fmt.Errorf("fatal error during reading config file: %w", err))
	}

	db, err := repository.NewGORM(c.DB)
	if err != nil {
		panic(fmt.Errorf("fatal error during initialization of DB: %w", err))
	}
	if err := repository.Migrate(db); err != nil {
		panic(fmt.Errorf("can't automigrate: %w", err))
	}

	repos := repository.NewRepository(db)
	serv := service.NewService(repos)
}
