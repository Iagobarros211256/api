package autenticacao

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//criar token retorna um token assinado com as permissoes do usuario
func CriarToken(usuarioID uint64) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["usuarioId"] = usuarioID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	return token.SignedString([]byte("config.SecretKey"))

}

//verifica se o token passado na requisicao e valido
func ValidarToken (r *http.Request) error {
	retrn nil
}