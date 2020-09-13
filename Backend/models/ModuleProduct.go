package Model

import "github.com/gofrs/uuid"

// Product - структура

type Product struct {
	Pid         uuid.UUID `json:"product_id"`
	Name        string    `json:"product_name"`
	Price       int16     `json:"product_price"`
	Dtime       int16     `json:"delivery_time"`
	Rating      int8      `json:"product_rating"`
	Description string    `json:"product_description"`
	Image       string    `json:"product_image"`
	Orgid       uuid.UUID `json:"organisation_id"`
	Orgname     string    `json:"organisation_name"`
	Catname     string    `json:"category_name"`
}
