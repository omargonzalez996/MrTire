package user

import (
	"minnell/types"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	//obtener los datos del body
	//buscar el usuario en la base de datos
	//si no existe devolver un error
	//si existe comparar la contraseña
	//si la contraseña no coincide devolver un error
	//si la contraseña coincide devolver un token
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	//obtener los datos del body
	var payload types.RegisterUserPayload
	// verificar que el usuario no exista
	//si existe devolver un error
	//si no existe crearlo
	//encriptar la contraseña
}
