package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/carlosvallim/gologin/dao"
	"github.com/carlosvallim/gologin/models"
	"github.com/dgrijalva/jwt-go"
)

const (
	idKey    = "id"
	nameKey  = "username"
	emailKey = "email"
)

var (
	SecretKey = []byte("29607b9e17f4c5266a2d33aca075ab62")
)

// Credentials contem os dados necessarios para efetuar login
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//GenerateToken generates a jwt token and assign a username to it's claims and return it
func (a *Authentication) GenerateToken(user *models.Usuario) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		idKey:    user.ID,
		nameKey:  user.Username,
		emailKey: user.Email,
	})
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		log.Fatal("Error in generating key")
		return "", err
	}
	return tokenString, nil
}

//ValidateToken parses a jwt token and returns the username it it's claims
func (a *Authentication) ValidateToken(tokenStr string) (*models.Usuario, error) {
	if tokenStr == "" {
		return nil, fmt.Errorf("Token não informado")
	}

	token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		id := int(claims[idKey].(float64))

		user, err := dao.GetUserByID(id)

		if err != nil {
			return nil, fmt.Errorf("não pode verificar usuário: %v", err)
		}
		if user == nil {
			return nil, fmt.Errorf("usuário [%d] não é mais válido", id)
		}
		return user, nil
	}

	return nil, fmt.Errorf("token inválido")
}

func catchAuthError(w http.ResponseWriter) {
	if r := recover(); w != nil {
		w.WriteHeader(http.StatusForbidden)
		_, _ = fmt.Fprintf(w, "erro ao autenticar: %v", r)
	}
}

// CreateTokenEndpoint autentica o usuario e devolve o token
func (a *Authentication) CreateTokenEndpoint(response http.ResponseWriter, request *http.Request) {

	defer catchAuthError(response)

	var creds Credentials
	if err := json.NewDecoder(request.Body).Decode(&creds); err != nil {
		panic(fmt.Errorf("payload inválido: %v", err))
	}

	user, err := dao.Authenticate(creds.Email, creds.Password)
	if err != nil {
		panic(err)
	}

	tokenString, err := a.GenerateToken(user)
	if err != nil {
		panic(fmt.Errorf("erro ao criar token JWT usuario %s: %v", user.Username, err))
	}

	SetAuthToken(response, tokenString)

	response.Header().Set("content-type", "application/json")
	_, _ = response.Write([]byte(`{ "token": "` + tokenString + `" }`))
}

// HandleLogout limpa o cookie
func (a *Authentication) HandleLogout(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(rw, "aceita somente POST")
		return
	}
	ClearAuthToken(rw)
	fmt.Fprintln(rw, "goodbye")
}
