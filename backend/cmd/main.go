package main

import (
	"app/internal/bootstrap"
	"app/internal/database/connections"
	"app/internal/database/models"
	"app/internal/server"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main(){
	var cfg bootstrap.Config

	//Env Flags
	flag.IntVar(&cfg.Port, "port", 4000, "API server port")
	flag.StringVar(&cfg.Env, "env", "development", "Enviornment (development|staging|production)")
	flag.StringVar(&cfg.DB.DSN, "db-dsn", os.Getenv("DATABASE_URL"), "PostgreSQL DSN")
	//Db Flags
	flag.IntVar(&cfg.DB.MaxOpenConns, "db-max-open-conns", 25, "Postgres max open connections")
	flag.IntVar(&cfg.DB.MaxIdleConns, "db-max-idle-conns", 25, "Postgres max idle connections")
	flag.StringVar(&cfg.DB.MaxIdleTime, "db-max-idle-time", "15m", "Postgres max idle connection time")
	flag.Parse()

	//Intialize stdout logger prefixed with date and time
	logger := log.New(os.Stdout, "", log.Ldate | log.Ltime)

	//Initialize Database Connection
	db := databaseConnection(cfg, logger)
	defer db.Close()
	logger.Printf("database connection pool established")

	///Initialize Database Driver
	driver := migrationDriver(db, logger)
	logger.Printf("database driver established ")

	//Initalize migrator
	migrator := migratorInstance(driver, logger)

	//Migrate Up
	err := migrator.Up()
	if err != nil && err != migrate.ErrNoChange{
		logger.Fatal(err)
	}

	logger.Printf("database migrations applied")


	//Intitalize application containing config and logger
	app := &bootstrap.Application{
		Config: cfg,
		Logger: logger,
		Models: models.NewModels(db),
	}

	//Create Server
	httpServer := &http.Server{
		Addr: fmt.Sprintf(":%d", cfg.Port),
		Handler: server.NewServer(app),
		IdleTimeout: time.Minute,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("Starting %s server on %s", cfg.Env, httpServer.Addr)
	err  = httpServer.ListenAndServe()
	logger.Fatal(err)
}


func databaseConnection(cfg bootstrap.Config, logger *log.Logger) *sql.DB{
	//Create connection pool
	db, err := connections.Connect(cfg)

	if err != nil{
		//Log and Exit
		logger.Fatal(err)
	}

	return db
}

func migrationDriver(db *sql.DB, logger *log.Logger) database.Driver{
	driver, err := postgres.WithInstance(db, &postgres.Config{})

	if err != nil {
		logger.Fatal(err)
	}

	return driver
}


func migratorInstance(driver database.Driver, logger *log.Logger) *migrate.Migrate {
	migrator, err := migrate.NewWithDatabaseInstance("file://internal/migrations", "postgres", driver)

	if err != nil {
		logger.Fatal(err)
	}

	return migrator
}