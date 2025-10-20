package models

type Contact struct {
	Id    uint   `gorm:"primaryKey"`               // Usa uint para IDs, y marca como clave primaria
	Name  string `gorm:"not null"`                 // Indica que el campo no puede ser nulo
	Email string `gorm:"type:varchar(100);unique"` // Tipo específico y único
	Phone string `gorm:"type:varchar(15)"`         // Tipo específico para el teléfono
}

// GORM por defecto sigue una convención de nombres en plural para las tablas.
// Si la estructura se llama "Contact" (en singular), GORM la mapeará a una tabla llamada "contacts" (en plural).
// Esto es parte de la funcionalidad de GORM y está diseñado para seguir una convención de nombres común en ORM.
// Si se decide usar "contact" como nombre de la tabla, se debe agregar la etiqueta correspondiente
// func (Contact) TableName() string {
// 	return "contact" // Nombra la tabla como "contact"
// }
