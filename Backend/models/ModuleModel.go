package Model

import "github.com/gofrs/uuid"

// Module - структура
type Module struct {
	ID          uuid.UUID `json:"module_id"`
	Name        string    `json:"module_name"`
	Description string    `json:"module_description"`
	Link        string    `json:"module_link"`
	Icon        string    `json:"module_icon"`
}
