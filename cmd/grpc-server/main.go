package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/halink0803/zerolog-graylog-hook/graylog"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
	"github.com/ozonmp/trv-airport-api/internal/config"
	"github.com/ozonmp/trv-airport-api/internal/database"
	"github.com/ozonmp/trv-airport-api/internal/server"
	"github.com/ozonmp/trv-airport-api/internal/tracer"
)

var (
	batchSize uint = 2
)

func main() {
	if err := config.ReadConfigYML("config.yml"); err != nil {
		log.Fatal().Err(err).Msg("Failed init configuration")
	}
	cfg := config.GetConfigInstance()

	log.Logger = zerolog.New(os.Stdout).With().Timestamp().Caller().Logger().Output(zerolog.ConsoleWriter{Out: os.Stderr})
	hook, err := graylog.NewGraylogHook(cfg.Gelf.URL)
	if err != nil {
		panic(err)
	}
	log.Logger = log.Hook(hook)

	migration := flag.Bool("migration", true, "Defines the migration start option")
	flag.Parse()

	log.Info().
		Str("version", cfg.Project.Version).
		Str("commitHash", cfg.Project.CommitHash).
		Bool("debug", cfg.Project.Debug).
		Str("environment", cfg.Project.Environment).
		Msgf("Starting service: %s", cfg.Project.Name)

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	// default: zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if cfg.Project.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.SslMode,
	)

	db, err := database.NewPostgres(dsn, cfg.Database.Driver)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed init postgres")
	}
	defer db.Close()

	*migration = false // todo: need to delete this line for homework-4
	if *migration {
		if err = goose.Up(db.DB, cfg.Database.Migrations); err != nil {
			log.Error().Err(err).Msg("Migration failed")

			return
		}
	}

	tracing, err := tracer.NewTracer(&cfg)
	if err != nil {
		log.Error().Err(err).Msg("Failed init tracing")

		return
	}
	defer tracing.Close()

	if err := server.NewGrpcServer(db, batchSize).Start(&cfg); err != nil {
		log.Error().Err(err).Msg("Failed creating gRPC server")

		return
	}
}
