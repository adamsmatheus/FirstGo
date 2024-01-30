package autenticacao

import (
	"GoTest/src/config"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
	"time"
)

func CriarToken(usuarioID uint64) (string, error) {

	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6)
	permissoes["usuarioId"] = usuarioID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	return token.SignedString(config.SecretKey)
}

func ValidarToken(r *http.Request) error {
	tokenString := ExtrairToken(r)
	token, erro := jwt.Parse(tokenString)
}

func ExtrairToken(r *http.Request) string {

	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return " "
}
