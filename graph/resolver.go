//go:generate go run github.com/99designs/gqlgen
package graph

import (
	"github.com/carlosvallim/gologin/auth"
	"github.com/carlosvallim/gologin/models"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

//Resolver -
type Resolver struct {
	auth auth.Authentication
	user models.Usuario
}
