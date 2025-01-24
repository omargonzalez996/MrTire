package main

import (
	"database/sql"
	"log"
	"minnell/cmd/api"
	"minnell/config"
	"minnell/db"

	"github.com/go-sql-driver/mysql"
)

func main() {
	// configuracion de la conexion a la base de datos
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}

	// inicializacion de la base de datos
	initStorage(db)

	// inicializacion del servidor
	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB: Connected")
}
