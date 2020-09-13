package Controller

import (

	// import userModel
	"database/sql"
	"errors"
	Database "go-postgres/config"
	Helpers "go-postgres/helpers"
	ModuleProduct "go-postgres/models"

	// Postgres Driver
	_ "github.com/lib/pq"
	"github.com/prometheus/common/log"
)

//------------------------- Controlling Services Between Repository and Database  ----------------
func GetSingleProduct(id string) (ModuleProduct.Product, error) {

	db := Database.CreateConnection()
	defer db.Close()

	// create a user of models.User type
	var product ModuleProduct.Product

	// create the select sql query
	sqlStatement := `select
	app_product.product_id,
	app_product.product_name,
	app_product.product_price,
	app_product.delivery_time,
	app_product.product_rating,
	app_product.product_description,
	app_product.product_image,
	app_product.organisation_id,
	app_organisation.organisation_name,
	app_category.category_name
from
	app_product
left join app_organisation on
	app_product.organisation_id = app_organisation.organisation_id
left join app_category on
	app_product.category_id = app_category.category_id 
where
	app_product.product_id =$1`

	// execute the sql statement
	row := db.QueryRow(sqlStatement, id)

	// unmarshal the row object to user
	err := row.Scan(&product.Pid, &product.Name, &product.Price, &product.Dtime, &product.Rating, &product.Description, &product.Image, &product.Orgid, &product.Orgname, &product.Catname)

	switch err {
	case sql.ErrNoRows:
		err = errors.New("Нет данных")
	case nil:
		return product, nil
	default:
		err = errors.New("Ошибка запроса")
	}

	// return empty module on error
	return product, err

}

// GetHot all
func GetHot(module_id string, sort_type string) ([]ModuleProduct.Product, error) {

	db := Database.CreateConnection()
	defer db.Close()
	// var cat, rat, org = false, false, false
	var products []ModuleProduct.Product
	var sqlStatement = `select
	app_product.product_id,
	app_product.product_name,
	app_product.product_price,
	app_product.delivery_time,
	app_product.product_rating,
	app_product.product_description,
	app_product.product_image,
	app_product.organisation_id,
	app_organisation.organisation_name,
	app_category.category_name
from
	app_product
left join app_organisation on
	app_product.organisation_id = app_organisation.organisation_id
left join app_category on
	app_product.category_id = app_category.category_id 
where
	app_product.module_id =$1 and app_product.product_rating = $2 order by random()`

	if Helpers.IsValidUUID(sort_type) {
		sqlStatement = `select
	app_product.product_id,
	app_product.product_name,
	app_product.product_price,
	app_product.delivery_time,
	app_product.product_rating,
	app_product.product_description,
	app_product.product_image,
	app_product.organisation_id,
	app_organisation.organisation_name,
	app_category.category_name
from
	app_product
left join app_organisation on
	app_product.organisation_id = app_organisation.organisation_id
left join app_category on
	app_product.category_id = app_category.category_id
where
	app_product.module_id = $1
	and (app_product.category_id = $2
	or app_product.organisation_id = $2) order by random()`
	}

	rows, err := db.Query(sqlStatement, module_id, sort_type)

	if err != nil {
		log.Errorln("Unable to execute the query. %v", err)
	}

	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var product ModuleProduct.Product

		err = rows.Scan(&product.Pid, &product.Name, &product.Price, &product.Dtime, &product.Rating, &product.Description, &product.Image, &product.Orgid, &product.Orgname, &product.Catname)

		if err != nil {
			log.Errorln("Unable to scan the row. %v", err)
		}

		products = append(products, product)

	}

	return products, err
}

// GetHotProductsCollection all
func GetHotProductsCollection(module_id string) ([]ModuleProduct.Product, error) {

	db := Database.CreateConnection()
	defer db.Close()
	// var cat, rat, org = false, false, false
	var products []ModuleProduct.Product

	var sqlStatement = `select
	app_product.product_id,
	app_product.product_name,
	app_product.product_price,
	app_product.delivery_time,
	app_product.product_rating,
	app_product.product_description,
	app_product.product_image,
	app_product.organisation_id,
	app_organisation.organisation_name,
	app_category.category_name
from
	app_product
left join app_organisation on
	app_product.organisation_id = app_organisation.organisation_id
left join app_category on
	app_product.category_id = app_category.category_id 
where
	app_product.module_id =$1`

	rows, err := db.Query(sqlStatement, module_id)

	if err != nil {
		log.Errorln("Unable to execute the query. %v", err)
	}

	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var product ModuleProduct.Product

		err = rows.Scan(&product.Pid, &product.Name, &product.Price, &product.Dtime, &product.Rating, &product.Description, &product.Image, &product.Orgid, &product.Orgname, &product.Catname)

		if err != nil {
			log.Errorln("Unable to scan the row. %v", err)
		}

		products = append(products, product)

	}

	return products, err
}
