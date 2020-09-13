package Model

import "github.com/gofrs/uuid"

// User - Модель для пользователя
type User struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Location string    `json:"location"`
	Age      int64     `json:"age"`
	Password string    `json:"password"`
}
