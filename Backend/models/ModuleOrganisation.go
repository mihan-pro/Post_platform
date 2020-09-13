package Model

import "github.com/gofrs/uuid"

// Organisation - структура

type Organisation struct {
	Orgid       uuid.UUID `json:"organisation_id"`
	Name        string    `json:"organisation_name"`
	Description string    `json:"organisation_description"`
	Logo        string    `json:"organisation_logo"`
	Address     string    `json:"organisation_address"`
	Rating      int8      `json:"organisation_rating"`
}
