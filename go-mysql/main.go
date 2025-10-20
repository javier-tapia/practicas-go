package main

import (
	"bufio"
	"fmt"
	"go-mysql/database"
	"go-mysql/handlers"
	"go-mysql/models"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Establecer conexión a la BBDD
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	for {
		fmt.Println("\nMenú:")
		fmt.Println("1. Listar contactos")
		fmt.Println("2. Obtener contacto por ID")
		fmt.Println("3. Crear nuevo contacto")
		fmt.Println("4. Actualizar contacto")
		fmt.Println("5. Eliminar contacto")
		fmt.Println("6. Salir")
		fmt.Print("Seleccione una opción: ")

		// Leer la opción seleccionada por el usuario
		var option int
		fmt.Scanln(&option)

		switch option {
		case 1:
			handlers.ListContacts(db)
		case 2:
			fmt.Print("Ingrese el ID del contacto: ")
			var idContact int
			fmt.Scanln(&idContact)
			handlers.GetContactById(db, idContact)
		case 3:
			newContact := inputContactDetails(option)
			handlers.CreateContact(db, newContact)
			handlers.ListContacts(db)
		case 4:
			updateContact := inputContactDetails(option)
			handlers.UpdateContact(db, updateContact)
			handlers.ListContacts(db)
		case 5:
			fmt.Print("Ingrese el ID del contacto que quiere eliminar: ")
			var idContact int
			fmt.Scanln(&idContact)
			handlers.DeleteContact(db, idContact)
		case 6:
			fmt.Println("Saliendo del programa...")
			return
		default:
			fmt.Println("Opción incorrecta. Seleccione una opción válida.")
		}
	}
}

// Ingresa los detalles del contacto desde la entrada estandar
func inputContactDetails(option int) models.Contact {
	// Leer la entrada del usuario utilizando bufio
	reader := bufio.NewReader(os.Stdin)

	var contact models.Contact

	if option == 4 {
		fmt.Print("Ingrese el ID del contacto que quiere editar: ")
		var idContact int
		fmt.Scanln(&idContact)
		contact.Id = idContact
	}

	fmt.Print("Ingrese el nombre del contacto: ")
	name, _ := reader.ReadString('\n')
	contact.Name = strings.TrimSpace(name)

	fmt.Print("Ingrese el email del contacto: ")
	email, _ := reader.ReadString('\n')
	contact.Email = strings.TrimSpace(email)

	fmt.Print("Ingrese el teléfono del contacto: ")
	phone, _ := reader.ReadString('\n')
	contact.Phone = strings.TrimSpace(phone)

	return contact
}
