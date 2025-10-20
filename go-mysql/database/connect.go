package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Connect() (*sql.DB, error) {
	// Cargar variables de entorno desde el archivo .env
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	// Cadena de conexión a la base de datos MySQL
	dns := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	// Abrir conexión a la BBDD
	db, err := sql.Open("mysql", dns)
	if err != nil {
		return nil, err
	}

	// Verificar la conexión a la BBDD
	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Conexión a la base de datos MySQL exitosa")
	return db, nil
}
