package dao

import (
	"github.com/carlosvallim/gologin/data"
	"github.com/jmoiron/sqlx"
)

//DAO - struct de conexão
type DAO struct {
	db *sqlx.DB
}

//New - funcão de conexão com o banco de dados
func New() (*DAO, error) {
	database, err := data.ConnectPostgres()

	if err != nil {
		return nil, err
	}

	return &DAO{database}, nil
}

//Close - função para fechar conexão com o banco de dados
func (d *DAO) Close() error {
	return d.db.Close()
}
