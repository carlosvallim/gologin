package data

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" //pacote de conexão com o postgresql
	"github.com/pkg/errors"
)

var Db *sqlx.DB

//ConnectPostgres - parametros de conexão com o banco de dados
func ConnectPostgres() (*sqlx.DB, error) {
	if err := godotenv.Load(); err != nil {
		return nil, errors.Wrap(err, "não pode carregar as váriaveis de conexão")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	psqlInfo := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName)
	db, err := sqlx.Connect("postgres", psqlInfo)

	if err != nil {
		return nil, errors.Wrap(err, "não foi possível criar conexão com o banco de dados")
	}

	err = db.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "falha ao testar a comunicação com o banco de dados")
	}

	return db, nil
}
