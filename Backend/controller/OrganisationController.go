package Controller

import (

	// import userModel
	"database/sql"
	"errors"
	Database "go-postgres/config"
	ModuleOrganisation "go-postgres/models"

	// Postgres Driver
	_ "github.com/lib/pq"
	"github.com/prometheus/common/log"
)

//------------------------- Controlling Services Between Repository and Database  ----------------

//GetSingleOrganisation asdsadjk
func GetSingleOrganisation(id string) (ModuleOrganisation.Organisation, error) {

	db := Database.CreateConnection()
	defer db.Close()

	var organisation ModuleOrganisation.Organisation

	// create the select sql query
	sqlStatement := `select
	organisation_id,
	organisation_name,
	organisation_description,
	organisation_logo,
	organisation_address,
	organisation_rating
from
	app_organisation where organisation_id =$1 `

	// execute the sql statement
	row := db.QueryRow(sqlStatement, id)

	// unmarshal the row object to user
	err := row.Scan(&organisation.Orgid, &organisation.Name, &organisation.Description, &organisation.Logo, &organisation.Address, &organisation.Rating)

	switch err {
	case sql.ErrNoRows:
		err = errors.New("Нет данных")
	case nil:
		return organisation, nil
	default:
		err = errors.New("Ошибка запроса")
	}

	// return empty module on error
	return organisation, err

}

// GetOrganisations all
func GetOrganisations(id string) ([]ModuleOrganisation.Organisation, error) {

	db := Database.CreateConnection()
	defer db.Close()

	var organisations []ModuleOrganisation.Organisation

	sqlStatement := `select distinct
	app_organisation.organisation_id,
	app_organisation.organisation_name,
	app_organisation.organisation_description,
	app_organisation.organisation_logo,
	app_organisation.organisation_address,
	app_organisation.organisation_rating
from
	app_organisation
left join app_product on
	app_product.organisation_id = app_organisation.organisation_id
where
	app_product.module_id =$1`

	rows, err := db.Query(sqlStatement, id)

	if err != nil {
		log.Errorln("Unable to execute the query. %v", err)
	}

	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var organisation ModuleOrganisation.Organisation

		err = rows.Scan(&organisation.Orgid, &organisation.Name, &organisation.Description, &organisation.Logo, &organisation.Address, &organisation.Rating)

		if err != nil {
			log.Errorln("Unable to scan the row. %v", err)
		}

		organisations = append(organisations, organisation)

	}

	return organisations, err
}
