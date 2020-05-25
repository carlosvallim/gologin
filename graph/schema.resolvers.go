package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/carlosvallim/gologin/dao"
	"github.com/carlosvallim/gologin/graph/generated"
	"github.com/carlosvallim/gologin/graph/model"
	"github.com/carlosvallim/gologin/models"
)

//ConnectBD - efetua a conexão com o banco de dados
func (r *Resolver) ConnectBD() *dao.DAO {
	db, err := dao.New()

	if err != nil {
		log.Fatalf("erro ao conectar com o banco de dados: %v", err)
	}

	return db
}

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateUsuario(ctx context.Context, name string, email string, password string) (bool, error) {
	db := r.ConnectBD()
	defer db.Close()

	_, err := db.CreateUsuario(name, email, password)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *mutationResolver) UpdateUsuario(ctx context.Context, id int, name *string, email *string, password *string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteUsuario(ctx context.Context, id int) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Usuarios(ctx context.Context) ([]*models.Usuario, error) {
	db := r.ConnectBD()
	defer db.Close()

	usuarios, err := db.Usuarios()
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
