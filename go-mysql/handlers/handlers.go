package handlers

import (
	"database/sql"
	"fmt"
	"go-mysql/models"
	"log"
)

// Lista todos los contactos desde la BBDD
func ListContacts(db *sql.DB) {
	// Consulta SQL para seleccionar todos los contactos
	query := "SELECT * FROM contact"

	// Ejecutar la consulta a la BBDD
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	// Iterar sobre los resultados y mostrarlos
	fmt.Println("\nLISTA DE CONTACTOS:")
	fmt.Println("----------------------------------------------------------------------------------------")
	for rows.Next() {
		// Instancia de modelo 'contact'
		contact := models.Contact{}

		var valueEmail sql.NullString

		err := rows.Scan(&contact.Id, &contact.Name, &valueEmail, &contact.Phone)
		if err != nil {
			log.Fatal(err)
		}

		if valueEmail.Valid {
			contact.Email = valueEmail.String
		} else {
			contact.Email = "No posee"
		}

		fmt.Printf("ID: %d, Nombre: %s, Email: %s, Teléfono: %s\n",
			contact.Id, contact.Name, contact.Email, contact.Phone,
		)
		fmt.Println("----------------------------------------------------------------------------------------")
	}
}

// Obtiene un contacto de la BBDD mediante su ID
func GetContactById(db *sql.DB, contactId int) {
	// Consulta SQL para seleccionar un contacto por su ID
	query := "SELECT * FROM contact WHERE id = ?"

	row := db.QueryRow(query, contactId)

	// Instancia de modelo 'contact'
	contact := models.Contact{}
	var valueEmail sql.NullString

	// Escanear el resultado en el modelo contact
	err := row.Scan(&contact.Id, &contact.Name, &valueEmail, &contact.Phone)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatalf("No se encontró ningún contacto con el ID %d", contactId)
		}
	}

	if valueEmail.Valid {
		contact.Email = valueEmail.String
	} else {
		contact.Email = "No posee"
	}

	fmt.Println("\nCONTACTO:")
	fmt.Println("----------------------------------------------------------------------------------------")
	fmt.Printf("ID: %d, Nombre: %s, Email: %s, Teléfono: %s\n",
		contact.Id, contact.Name, contact.Email, contact.Phone,
	)
	fmt.Println("----------------------------------------------------------------------------------------")
}

// Registra un nuevo contacto en la base de datos
func CreateContact(db *sql.DB, contact models.Contact) {
	// Consulta SQL para insertar un nuevo contacto
	query := "INSERT INTO contact (name, email, phone) VALUES (?, ?, ?)"

	// Ejecutar la sentencia SQL
	_, err := db.Exec(query, contact.Name, contact.Email, contact.Phone)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Nuevo contacto registrado con éxito")
}

// Actualiza un contacto existente en la BBDD
func UpdateContact(db *sql.DB, contact models.Contact) {
	// Consulta SQL para actualizar un contacto
	query := "UPDATE contact SET name = ?, email = ?, phone = ? WHERE id = ?"

	// Ejecutar la sentencia SQL
	_, err := db.Exec(query, contact.Name, contact.Email, contact.Phone, contact.Id)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("El contacto se ha actualizado con éxito")
}

// Elimina un contacto de la BBDD
func DeleteContact(db *sql.DB, contactId int) {
	// Sentencia SQL para eliminar un contacto por su ID
	query := "DELETE FROM contact WHERE id = ?"

	// Ejecutar la sentencia SQL
	_, err := db.Exec(query, contactId)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("El contacto se ha eliminado con éxito")
}
