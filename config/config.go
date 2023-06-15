package config

import (
	"os"
)

var (
	HOST     = os.Getenv("PG_HOST")
	PORT     = os.Getenv("PG_PORT")
	USERNAME = os.Getenv("PG_USER")
	PASSWORD = os.Getenv("PG_PASSWORD")
	DB_NAME  = os.Getenv("PG_DB_NAME")
	dialect = "postgres"
	SECRET_KEY string = "_OkTaa"
)