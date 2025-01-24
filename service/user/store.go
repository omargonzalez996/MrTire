package user

import (
	"database/sql"
	"fmt"
	"minnell/types"
	"strings"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) ObtenerUsuarioPorNick(nick string) (*types.User, error) {
	query := `SELECT *
	FROM usuario
	WHERE nick = ?`

	rows, err := s.db.Query(query, nick)
	if err != nil {
		return nil, err
	}

	u := new(types.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}
	if u.ID == 0 {
		return nil, fmt.Errorf("usuario no encontrado")
	}
	return u, nil
}

func (s *Store) ObtenerUsuarioPorParametros(nombre string, email string) (*types.User, error) {
	query := `SELECT * FROM usuario
	WHERE nombre = ? OR email = ?`

	rows, err := s.db.Query(query, nombre, email)
	if err != nil {
		return nil, err
	}

	u := new(types.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}
	if u.ID == 0 {
		return nil, fmt.Errorf("usuario no encontrado")
	}
	return u, nil
}

func (s *Store) ObtenerUsuarioPorID(id int) (*types.User, error) {
	query := `SELECT * FROM usuario
	WHERE id = ?`

	rows, err := s.db.Query(query, id)
	if err != nil {
		return nil, err
	}

	u := new(types.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}
	if u.ID == 0 {
		return nil, fmt.Errorf("usuario no encontrado")
	}
	return u, nil
}

func (s *Store) CrearUsuario(user types.RegisterUserPayload) (*types.UserCredentials, error) {

	fmt.Println(user.TipoUsuario)

	//Generar nick a partir del email
	nick, err := GenerarNick(user.Email)
	if err != nil {
		return nil, err
	}

	query := `INSERT INTO usuario (nombre, nick, email, password, tipousuario_id)
		VALUES (?, ?, ?, ?, ?)`

	result, err := s.db.Exec(query, user.Nombre, nick, user.Email, user.Password, user.TipoUsuario)
	if err != nil {
		return nil, err
	}

	// Obtén el ID del último registro insertado
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	response := &types.UserCredentials{
		ID:       id,
		Nick:     nick,
		Password: user.Password,
	}

	return response, nil
}

func GenerarNick(email string) (string, error) {
	fmt.Println("email:", email)
	// Verifica que el email contenga un "@"
	if !strings.Contains(email, "@") {
		return "", fmt.Errorf("el email no es valido")
	}

	// Divide el email en dos partes usando "@"
	nick := strings.SplitN(email, "@", 2)

	// Retorna la parte antes del "@" como nick
	return string(nick[0]), nil
}

func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)

	err := rows.Scan(
		&user.ID,
		&user.Nombre,
		&user.Nick,
		&user.Password,
		&user.TipoUsuario,
		&user.Fecha_registro,
		&user.Activo,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}
