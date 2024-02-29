package bootstrap

import (
	"app/internal/database/models"
	"log"
)

//Application
const Version = "1.0.0"

//Application config settings
type Config struct{
	Port int  	//server network port
	Env string	//Operating env (development, staging, production)
	DB struct { //DB struct that holds config for database connection
		DSN string
		MaxOpenConns int
		MaxIdleConns int
		MaxIdleTime string
	}
}

//Hold dependencies of HTTP handlers, helpers, and middleware
type Application struct{
	Config Config
	Logger *log.Logger
	Models models.Models
}