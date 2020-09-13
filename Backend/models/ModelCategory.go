package Model

import "github.com/gofrs/uuid"

// Category - структура
type Category struct {
	Catid       uuid.UUID `json:"category_id"`
	Name        string    `json:"category_name"`
	Description string    `json:"category_description"`
	Logo        string    `json:"category_logo"`
	Modid       uuid.UUID `json:"module_id"`
}
