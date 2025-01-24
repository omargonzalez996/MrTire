package types

import "time"

type UserStore interface {
	ObtenerUsuarioPorNick(nick string) (*User, error)
	ObtenerUsuarioPorParametros(nombre string, email string) (*User, error)
	ObtenerUsuarioPorID(id int) (*User, error)
	CrearUsuario(RegisterUserPayload) (*UserCredentials, error)
}

type User struct {
	ID             int       `json:"id"`
	Nombre         string    `json:"nombre"`
	Nick           string    `json:"nick"`
	Email          string    `json:"email"`
	Password       string    `json:"password"`
	TipoUsuario    string    `json:"tipoUsuario_id"`
	Fecha_registro time.Time `json:"fecha_registro"`
	Activo         int       `json:"activo"`
}

type RegisterUserPayload struct {
	Nombre       string `json:"nombre"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	PasswordHash string `json:"PasswordHash"`
	TipoUsuario  int    `json:"tipoUsuario_id"`
}

type UserCredentials struct {
	ID       int64  `json:"id"`
	Nick     string `json:"nick"`
	Password string `json:"password"`
}

type LoginPayload struct {
	Nick     string `json:"nick"`
	Password string `json:"password"`
}
