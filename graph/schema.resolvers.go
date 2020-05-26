package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/carlosvallim/gologin/auth"
	"github.com/carlosvallim/gologin/dao"
	"github.com/carlosvallim/gologin/graph/generated"
	"github.com/carlosvallim/gologin/graph/model"
	"github.com/carlosvallim/gologin/models"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (bool, error) {
	user := auth.UserFromContext(ctx)
	if user == nil {
		return false, fmt.Errorf("acesso negado")
	}

	todo := model.NewTodo{
		Text:   input.Text,
		UserID: input.UserID,
	}
	_, err := dao.CreateTodo(todo)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *mutationResolver) CreateUsuario(ctx context.Context, username string, email string, password string) (string, error) {
	//a := auth.AuthenticatorFromContext(ctx)
	_, err := dao.CreateUsuario(username, email, password)
	if err != nil {
		return "", err
	}

	/*usuario, err := dao.GetUserByUsername(email)
	if err != nil {
		return "", err
	}

	_, err = a.GenerateToken(usuario)
	if err != nil {
		return "", err
	}*/

	return "Usuário criado com sucesso", nil
}

func (r *mutationResolver) UpdateUsuario(ctx context.Context, id int, username *string, email *string, password *string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteUsuario(ctx context.Context, id int) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (*models.Usuario, error) {
	/*rw := auth.WriterFromContext(ctx)
	a := auth.AuthenticatorFromContext(ctx)
	if rw == nil || a == nil {
		return nil, fmt.Errorf("Necessário injetar ResponseWriter no Context")
	}*/

	usuario, err := dao.Authenticate(input.Username, input.Password)
	if err != nil || usuario == nil {
		return nil, fmt.Errorf("Erro ao efetuar login, usuário ou senha inválido")
	}
	/*token, err := a.GenerateToken(usuario)
	if err != nil {
		return nil, err
	}
	auth.SetAuthToken(rw, token)*/
	return usuario, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*models.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Usuarios(ctx context.Context) ([]*models.Usuario, error) {
	usuarios, err := dao.Usuarios()
	if err != nil {
		return nil, err
	}

	return usuarios, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
