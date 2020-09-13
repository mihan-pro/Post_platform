package Controller

import (
	"database/sql"
	"errors"

	// import userModel
	Database "go-postgres/config"
	ModuleModel "go-postgres/models"

	// Postgres Driver
	_ "github.com/lib/pq"
	"github.com/prometheus/common/log"
)

//------------------------- Controlling Services Between Repository and Database  ----------------

// get one module from the DB by its module_id
func GetModule(id string) (ModuleModel.Module, error) {

	db := Database.CreateConnection()
	defer db.Close()

	// create a user of models.User type
	var module ModuleModel.Module

	// create the select sql query
	sqlStatement := `SELECT module_id, module_name, module_description,module_link, module_icon  FROM app_module WHERE module_id=$1`

	// execute the sql statement
	row := db.QueryRow(sqlStatement, id)

	// unmarshal the row object to user
	err := row.Scan(&module.ID, &module.Name, &module.Description, &module.Link, &module.Icon)

	switch err {
	case sql.ErrNoRows:
		err = errors.New("Нет данных")
	case nil:
		return module, nil
	default:
		err = errors.New("Ошибка запроса")
	}

	// return empty module on error
	return module, err
}

//RenderModule - gjkextybt cnhfyb
func RenderModule(id string) (string, error) {

	db := Database.CreateConnection()
	defer db.Close()
	apiKey := ""
	// create the select sql query
	sqlStatement := `SELECT api_key FROM app_module WHERE module_id=$1`

	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&apiKey)
	log.Infoln("get module: ", apiKey)
	return apiKey, err
}

// GetAllModules all
func GetAllModules() ([]ModuleModel.Module, error) {

	db := Database.CreateConnection()
	defer db.Close()

	var modules []ModuleModel.Module

	sqlStatement := `SELECT module_id, module_name, module_description,module_link, module_icon FROM app_module`
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Errorln("Unable to execute the query. %v", err)
	}

	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var module ModuleModel.Module

		err = rows.Scan(&module.ID, &module.Name, &module.Description, &module.Link, &module.Icon)

		if err != nil {
			log.Errorln("Unable to scan the row. %v", err)
		}

		modules = append(modules, module)

	}

	return modules, err
}
