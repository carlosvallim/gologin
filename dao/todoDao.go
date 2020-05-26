package dao

import (
	"fmt"

	"github.com/carlosvallim/gologin/data"
	"github.com/carlosvallim/gologin/graph/model"
)

//CreateTodo - cria todo
func CreateTodo(todo model.NewTodo) (bool, error) {
	res, err := data.Db.Exec(`
		insert into todo(text, done, user_id) values ($1, $2, $3)
	`, todo.Text, false, todo.UserID)

	if err != nil {
		return false, fmt.Errorf("erro ao criar todo: %w", err)
	} else if rows, _ := res.RowsAffected(); rows == 0 {
		return false, fmt.Errorf("Nenhuma linha afetada ao inserir todo")
	}

	return true, nil
}
