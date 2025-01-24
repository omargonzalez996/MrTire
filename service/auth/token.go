package auth

import (
	"fmt"
	"minnell/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Clave secreta para firmar el token
var secretKey = []byte(config.Envs.SecretKey)

// Estructura del JWT Claims (información que estará dentro del token)
type Claims struct {
	UserID uint   `json:"usuario_id"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

// Función para generar un JWT
func GenerarJWT(userID uint, email string) (string, error) {
	// Define las reclamaciones (payload)
	claims := Claims{
		UserID: userID,
		Email:  email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(), // El token expira en 24 horas
			Issuer:    "MrTire",
		},
	}

	// Crea un nuevo token con las reclamaciones y la clave secreta
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("error al generar el token: %v", err)
	}

	return signedToken, nil
}
