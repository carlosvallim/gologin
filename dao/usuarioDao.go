package dao

import (
	"database/sql"
	"fmt"

	"github.com/carlosvallim/gologin/models"
)

//Usuarios - busca os usuários
func (d *DAO) Usuarios() ([]*models.Usuario, error) {
	usuarios := []*models.Usuario{}

	err := d.db.Select(&usuarios, "select id, name, email, password from usuario ")

	if err != sql.ErrNoRows && err != nil {
		return nil, fmt.Errorf("erro ao buscar usuários: %w", err)
	}

	return usuarios, nil
}

//CreateUsuario - insere usuário no banco de dados
func (d *DAO) CreateUsuario(name string, email string, password string) (bool, error) {
	res, err := d.db.Exec(`
		insert into usuario(name, email, password) values ($1, $2, $3)
	`, name, email, password)

	if err != nil {
		return false, fmt.Errorf("erro ao criar usuário: %w", err)
	} else if rows, _ := res.RowsAffected(); rows == 0 {
		return false, fmt.Errorf("Nenhuma linha afetada ao inserir o usuário")
	}

	return true, nil
}

//UpdateUsuario - altera usuário no banco de dados
func (d *DAO) UpdateUsuario(id int, name string, email string, password string) (bool, error) {
	res, err := d.db.Exec(`
		update usuario set name = $1, email = $2, password = $3 where id = $4
	`, name, email, password, id)

	if err != nil {
		return false, fmt.Errorf("erro ao alterar usuário: %w", err)
	} else if rows, _ := res.RowsAffected(); rows == 0 {
		return false, fmt.Errorf("Nenhuma linha afetada ao alterar usuário")
	}

	return true, nil
}

//DeleteUsuario - deleta o usuário do banco de dados
func (d *DAO) DeleteUsuario(id int) (bool, error) {
	res, err := d.db.Exec(`
		delete from usuario where id = $1
	`, id)

	if err != nil {
		return false, fmt.Errorf("erro ao deletar usuário: %w", err)
	} else if rows, _ := res.RowsAffected(); rows == 0 {
		return false, fmt.Errorf("Nenhuma linha afetada ao deletar usuário")
	}

	return true, nil
}
