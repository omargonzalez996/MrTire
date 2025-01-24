package auth

import "golang.org/x/crypto/bcrypt"

func EncriptarPassword(password string) (string, error) {
	// Genera el hash usando bcrypt
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func VerificarPassword(hash, password string) bool {
	// Convertir el hash y la contraseña a []byte
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	// Si err es nil, la contraseña es correcta, sino es incorrecta
	return err == nil
}
