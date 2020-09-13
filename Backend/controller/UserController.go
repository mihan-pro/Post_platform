package Controller

import (
	"database/sql"

	"fmt"

	// import userModel
	Database "go-postgres/config"
	UserModel "go-postgres/models"

	// Postgres Driver
	"github.com/gofrs/uuid"
	//// Postgres Driver
	_ "github.com/lib/pq"

	"github.com/prometheus/common/log"
)

//------------------------- Controlling Services Between Repository and Database  ----------------

// insert one user in the DB
func InsertUser(user UserModel.User) uuid.UUID {

	// Define the postgres db connection
	db := Database.CreateConnection()
	defer db.Close()

	// create the insert sql query
	sqlStatement := `INSERT INTO users (name, location, age, password ) VALUES ($1, $2, $3, $4) RETURNING userid`

	// the inserted id will store in this id
	var id uuid.UUID

	err := db.QueryRow(sqlStatement, user.Name, user.Location, user.Age, user.Password).Scan(&id)

	if err != nil {
		log.Errorln("Failed to fetch URL:", err)
	}
	log.Info("Inserted a single record:", id)

	// return the inserted id
	return id
}

// get one user from the DB by its userid
func GetUser(id string) (UserModel.User, error) {

	db := Database.CreateConnection()
	defer db.Close()

	// create a user of models.User type
	var user UserModel.User

	// create the select sql query
	sqlStatement := `SELECT userid, name, age,location FROM users WHERE userid=$1`

	// execute the sql statement
	row := db.QueryRow(sqlStatement, id)

	// unmarshal the row object to user
	err := row.Scan(&user.ID, &user.Name, &user.Age, &user.Location)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return user, nil
	case nil:
		return user, nil
	default:
		log.Errorln("Unable to scan the row. %v", err)
	}

	// return empty user on error
	return user, err
}

// get one user from the DB by its userid
func GetAllUsers() ([]UserModel.User, error) {

	db := Database.CreateConnection()
	defer db.Close()

	var users []UserModel.User

	sqlStatement := `SELECT userid, "name", age, "location"  FROM users`
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Errorln("Unable to execute the query. %v", err)
	}

	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var user UserModel.User

		err = rows.Scan(&user.ID, &user.Name, &user.Age, &user.Location)

		if err != nil {
			log.Errorln("Unable to scan the row. %v", err)
		}

		users = append(users, user)

	}

	return users, err
}

// update user in the DB
func UpdateUser(id int64, user UserModel.User) int64 {

	db := Database.CreateConnection()
	defer db.Close()

	sqlStatement := `UPDATE users SET name=$2, location=$3, age=$4 WHERE userid=$1`
	res, err := db.Exec(sqlStatement, id, user.Name, user.Location, user.Age)

	if err != nil {
		log.Errorln("Unable to execute the query. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Errorln("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}

// delete user in the DB
func DeleteUser(id int64) int64 {

	db := Database.CreateConnection()
	defer db.Close()

	sqlStatement := `DELETE FROM users WHERE userid=$1`
	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Errorln("Unable to execute the query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Errorln("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}
