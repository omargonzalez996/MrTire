package types

type RegisterUserPayload struct {
	nombre      string `json:"nombre"`
	nick        string `json:"nick"`
	password    string `json:"password"`
	tipoUsuario string `json:"tipoUsuario"`
}
