package connections

import (
	"app/internal/bootstrap"
	"context"
	"database/sql"
	"time"
)

func Connect (cfg bootstrap.Config) (*sql.DB, error){
	//Connect to database
	db, err := sql.Open("pgx", cfg.DB.DSN)
	if err != nil {
		return nil, err
	}

	//Passing a value less than or equal to 0 means no limit
	db.SetMaxOpenConns(cfg.DB.MaxOpenConns)
	db.SetMaxIdleConns(cfg.DB.MaxIdleConns)

	//Convert idle duration
	duration, err := time.ParseDuration(cfg.DB.MaxIdleTime)

	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(duration)
	
	//Create Contex with 5-second timeout deadline
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()


	//Return error if connection couldnt be established within context deadline
	err = db.PingContext(ctx)
	if err != nil{
		return nil, err
	}
	return db, nil
}