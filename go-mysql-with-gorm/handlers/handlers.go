package handlers

import (
	"fmt"
	"go-mysql/models"
	"log"

	"gorm.io/gorm"
)

// Lista todos los contactos desde la BBDD
func ListContacts(db *gorm.DB) {
	var contacts []models.Contact
	if err := db.Find(&contacts).Error; err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nLISTA DE CONTACTOS:")
	fmt.Println("----------------------------------------------------------------------------------------")
	for _, contact := range contacts {
		if contact.Email == "" {
			contact.Email = "No posee"
		}
		fmt.Printf("ID: %d, Nombre: %s, Email: %s, Teléfono: %s\n",
			contact.Id, contact.Name, contact.Email, contact.Phone,
		)
		fmt.Println("----------------------------------------------------------------------------------------")
	}
}

// Obtiene un contacto de la BBDD mediante su ID
func GetContactById(db *gorm.DB, contactId uint) {
	var contact models.Contact
	if err := db.First(&contact, contactId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Fatalf("No se encontró ningún contacto con el ID %d", contactId)
		}
		log.Fatal(err)
	}

	if contact.Email == "" {
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
func CreateContact(db *gorm.DB, contact models.Contact) {
	if err := db.Create(&contact).Error; err != nil {
		log.Fatal(err)
	}
	log.Println("Nuevo contacto registrado con éxito")
}

// Actualiza un contacto existente en la BBDD
func UpdateContact(db *gorm.DB, contact models.Contact) {
	if err := db.Save(&contact).Error; err != nil {
		log.Fatal(err)
	}
	log.Println("El contacto se ha actualizado con éxito")
}

// Elimina un contacto de la BBDD
func DeleteContact(db *gorm.DB, contactId uint) {
	if err := db.Delete(&models.Contact{}, contactId).Error; err != nil {
		log.Fatal(err)
	}
	log.Println("El contacto se ha eliminado con éxito")
}
