package auth

import (
	"context"
	"net/http"
	"time"

	"github.com/carlosvallim/gologin/dao"
	"github.com/carlosvallim/gologin/models"
	"github.com/jmoiron/sqlx"
)

const (
	authCookieKey = "auth-token"
)

var (
	userCtxKey = &contextKey{"user"}
	rwCtxKey   = &contextKey{"responseWriter"}
	authCtxKey = &contextKey{"authenticator"}
)

type contextKey struct {
	name string
}

type Authentication struct {
	db        *sqlx.DB
	jwtSecret string
}

// New cria uma nova instancia do autenticador
func New(db *sqlx.DB, jwtSecret string) *Authentication {
	return &Authentication{db, jwtSecret}
}

//Middleware - função middleware de autenticação
func (a *Authentication) HTTPMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c, err := r.Cookie(authCookieKey)

			//Salva o responsewriter no context
			ctx := context.WithValue(r.Context(), rwCtxKey, w)
			ctx = context.WithValue(ctx, authCtxKey, a)
			r = r.WithContext(ctx)

			// Allow unauthenticated users in
			if err != nil || c == nil {
				next.ServeHTTP(w, r)
				return
			}

			//validate jwt token
			tokenStr := c.Value
			username, err := a.ValidateToken(tokenStr)
			if err != nil {
				ClearAuthToken(w)
				next.ServeHTTP(w, r)
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			// create user and check if user exists in db
			user, err := dao.GetUserByUsername(username.Email)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			// put it in context
			ctx = context.WithValue(ctx, userCtxKey, user)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

//ClearAuthToken grava no response um cookie vencido, para remove-lo do browser
func ClearAuthToken(w http.ResponseWriter) {
	expiration := time.Now().Add(-1 * 24 * time.Hour)
	cookie := http.Cookie{Name: authCookieKey, Value: "", Expires: expiration, HttpOnly: true}
	http.SetCookie(w, &cookie)
}

//SetAuthToken grava no response o cookie com o token fornecido
func SetAuthToken(w http.ResponseWriter, tokenString string) {
	expiration := time.Now().Add(30 * 24 * time.Hour)
	cookie := http.Cookie{Name: authCookieKey, Value: tokenString, Expires: expiration, HttpOnly: true}
	http.SetCookie(w, &cookie)
}

// WriterFromContext retorna o writer do contexto
func WriterFromContext(ctx context.Context) http.ResponseWriter {
	raw, _ := ctx.Value(rwCtxKey).(http.ResponseWriter)
	return raw
}

// UserFromContext finds the user from the context. REQUIRES Middleware to have run.
func UserFromContext(ctx context.Context) *models.Usuario {
	raw, _ := ctx.Value(userCtxKey).(*models.Usuario)
	return raw
}

// AuthenticatorFromContext retorna a instancia do Authentication utilizada
func AuthenticatorFromContext(ctx context.Context) *Authentication {
	return ctx.Value(authCtxKey).(*Authentication)
}
