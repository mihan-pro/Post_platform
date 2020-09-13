package Model

import "github.com/gofrs/uuid"

// Category - структура
type Comment struct {
	Comid   uuid.UUID `json:"comment_id"`
	Autor   uuid.UUID `json:"autor_id"`
	Pid     uuid.UUID `json:"product_id"`
	Comment string    `json:"comment_text"`
}
