package dao

import (
	"database/sql"
	"fmt"

	"github.com/carlosvallim/gologin/data"
	"github.com/carlosvallim/gologin/models"
	"golang.org/x/crypto/bcrypt"
)

//HashPassword hashes given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckPasswordHash hash compares raw password with it's hashed values
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//GetUserByUsername check if a user exists in database by given username
func GetUserByUsername(email string) (*models.Usuario, error) {
	var usuario models.Usuario

	err := data.Db.Get(&usuario, `
		select * from usuario where email = $1
	`, email)

	if err != sql.ErrNoRows && err != nil {
		return nil, fmt.Errorf("erro ao buscar usuário pelo username: %w", err)
	}

	return &usuario, nil
}

//GetUserByID check if a user exists in database and return the user object.
func GetUserByID(userID int) (*models.Usuario, error) {
	var username models.Usuario
	err := data.Db.Get(&username, `
		select * from usuario where id = $1
	`, userID)

	if err != sql.ErrNoRows && err != nil {
		return nil, fmt.Errorf("erro ao buscar usuário pelo ID: %w", err)
	}

	return &username, nil
}

//Authenticate - verifica a senha do usuário
func Authenticate(username string, password string) (*models.Usuario, error) {
	var user models.Usuario
	err := data.Db.Get(&user, `
		select * from usuario where email = $1
	`, username)

	if err != sql.ErrNoRows && err != nil {
		return nil, err
	}

	if CheckPasswordHash(password, user.Password) {
		return &user, nil
	}

	return nil, err
}

//Usuarios - busca os usuários
func Usuarios() ([]*models.Usuario, error) {
	usuarios := []*models.Usuario{}

	err := data.Db.Select(&usuarios, "select id, username, email, password from usuario ")

	if err != sql.ErrNoRows && err != nil {
		return nil, fmt.Errorf("erro ao buscar usuários: %w", err)
	}

	return usuarios, nil
}

//CreateUsuario - insere usuário no banco de dados
func CreateUsuario(name string, email string, password string) (bool, error) {
	hashPassword, err := HashPassword(password)
	if err != nil {
		return false, fmt.Errorf("Token de autenticação inválido")
	}

	res, err := data.Db.Exec(`
		insert into usuario(username, email, password) values ($1, $2, $3)
	`, name, email, hashPassword)

	if err != nil {
		return false, fmt.Errorf("erro ao criar usuário: %w", err)
	} else if rows, _ := res.RowsAffected(); rows == 0 {
		return false, fmt.Errorf("Nenhuma linha afetada ao inserir o usuário")
	}

	return true, nil
}

//UpdateUsuario - altera usuário no banco de dados
func UpdateUsuario(id int, name string, email string, password string) (bool, error) {
	res, err := data.Db.Exec(`
		update usuario set username = $1, email = $2, password = $3 where id = $4
	`, name, email, password, id)

	if err != nil {
		return false, fmt.Errorf("erro ao alterar usuário: %w", err)
	} else if rows, _ := res.RowsAffected(); rows == 0 {
		return false, fmt.Errorf("Nenhuma linha afetada ao alterar usuário")
	}

	return true, nil
}

//DeleteUsuario - deleta o usuário do banco de dados
func DeleteUsuario(id int) (bool, error) {
	res, err := data.Db.Exec(`
		delete from usuario where id = $1
	`, id)

	if err != nil {
		return false, fmt.Errorf("erro ao deletar usuário: %w", err)
	} else if rows, _ := res.RowsAffected(); rows == 0 {
		return false, fmt.Errorf("Nenhuma linha afetada ao deletar usuário")
	}

	return true, nil
}
