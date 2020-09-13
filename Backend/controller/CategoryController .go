package Controller

import (

	// import userModel
	"database/sql"
	"errors"
	Database "go-postgres/config"
	ModuleCategory "go-postgres/models"

	// Postgres Driver
	_ "github.com/lib/pq"
	"github.com/prometheus/common/log"
)

//------------------------- Controlling Services Between Repository and Database  ----------------

//GetSingleCategory asdsadjk
func GetSingleCategory(id string) (ModuleCategory.Category, error) {

	db := Database.CreateConnection()
	defer db.Close()

	var category ModuleCategory.Category

	// create the select sql query
	sqlStatement := `select
	category_id,
	category_name,
	category_description,
	category_logo,
	module_id
from
	app_category
where
	category_id = $1`

	// execute the sql statement
	row := db.QueryRow(sqlStatement, id)

	// unmarshal the row object to user
	err := row.Scan(&category.Catid, &category.Name, &category.Description, &category.Logo, &category.Modid)

	switch err {
	case sql.ErrNoRows:
		err = errors.New("Нет данных")
	case nil:
		return category, nil
	default:
		log.Error(err)
		err = errors.New("Ошибка запроса")
	}

	// return empty module on error
	return category, err

}

// GetOrganisations all
func GetCategorys(id string) ([]ModuleCategory.Category, error) {

	db := Database.CreateConnection()
	defer db.Close()

	var categorys []ModuleCategory.Category

	sqlStatement := `select
	category_id,
	category_name,
	category_description,
	category_logo,
	module_id
from
	app_category
where module_id =$1`

	rows, err := db.Query(sqlStatement, id)

	if err != nil {
		log.Errorln("Unable to execute the query. %v", err)
	}

	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var category ModuleCategory.Category

		err = rows.Scan(&category.Catid, &category.Name, &category.Description, &category.Logo, &category.Modid)

		if err != nil {
			log.Errorln("Unable to scan the row. %v", err)
		}

		categorys = append(categorys, category)

	}

	return categorys, err
}
