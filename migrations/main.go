package main

import (
	"embed"
	"fmt"
	"github.com/brenik/product-service/internal/config"
	"github.com/brenik/product-service/internal/pkg/db"

	_ "github.com/jackc/pgx/v4/stdlib"
	goose "github.com/pressly/goose/v3"
	"github.com/rs/zerolog/log"
)

//"github.com/pressly/goose"
//go install github.com/pressly/goose/v3/cmd/goose@latest

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {
	fmt.Println("Work with DB. Migrations")
	if err := config.ReadConfigYML("config.yml"); err != nil {
		log.Fatal().Err(err).Msg("Failed init configuration")
	}
	cfg := config.GetConfigInstance()

	conn, err := db.ConnectDB(&cfg.DB)
	if err != nil {
		log.Fatal().Err(err).Msg("sql.Open() error")
	}
	defer conn.Close()

	goose.SetBaseFS(embedMigrations)

	const cmd = "up"

	err = goose.Run(cmd, conn.DB, "migrations")
	if err != nil {
		log.Fatal().Err(err).Msg("goose.Status() error")
	}

}
