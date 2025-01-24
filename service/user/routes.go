package user

import (
	"database/sql"
	"fmt"
	"minnell/service/auth"
	"minnell/types"
	"minnell/utils"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	Store types.UserStore
}

func NewHandler(Store types.UserStore) *Handler {
	return &Handler{Store: Store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	//obtener los datos del body
	var payload types.LoginPayload

	if err := utils.ParsearJSON(r, &payload); err != nil {
		utils.CrearError(w, http.StatusBadRequest, err)
		return
	}

	//buscar el usuario en la base de datos
	usuario, err := h.Store.ObtenerUsuarioPorNick(payload.Nick)
	if err != nil {
		if err == sql.ErrNoRows {
			// No se encontró el usuario
			utils.CrearJSON(w, http.StatusNotFound, fmt.Errorf("credenciales incorrectas"))
			return
		}
	}

	//validar Password
	passwordCorrecto := auth.VerificarPassword(usuario.Password, payload.Password)
	if !passwordCorrecto {
		utils.CrearJSON(w, http.StatusBadRequest, "credenciales incorrectas")
		return
	}

	//generar Token
	token, err := auth.GenerarJWT(uint(usuario.ID), usuario.Email)
	if err != nil {
		utils.CrearError(w, http.StatusBadRequest, err)
	}

	utils.CrearJSON(w, http.StatusOK, token)
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	//obtener los datos del body
	var payload types.RegisterUserPayload

	if err := utils.ParsearJSON(r, &payload); err != nil {
		utils.CrearError(w, http.StatusBadRequest, err)
		return
	}

	// verificar que el usuario no exista
	_, err := h.Store.ObtenerUsuarioPorParametros(payload.Nombre, payload.Email)
	//si existe devolver un error
	if err == nil {
		utils.CrearError(w, http.StatusBadRequest, fmt.Errorf("ya existe un usuario con esos parametros"))
		return
	}

	//Encriptar password
	hash_pass, err := auth.EncriptarPassword(payload.Password)
	if err != nil {
		utils.CrearError(w, http.StatusInternalServerError, err)
		return
	}
	fmt.Println(payload.TipoUsuario)

	//si no existe crearlo
	creds, err := h.Store.CrearUsuario(types.RegisterUserPayload{
		Nombre:       payload.Nombre,
		Email:        payload.Email,
		PasswordHash: hash_pass,
		Password:     payload.Password,
		TipoUsuario:  payload.TipoUsuario,
	})

	if err != nil {
		utils.CrearError(w, http.StatusInternalServerError, err)
		return
	}

	//retornar credenciales
	utils.CrearJSON(w, http.StatusCreated, creds)
}
