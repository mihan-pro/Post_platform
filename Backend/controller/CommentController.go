package Controller

import (

	// import userModel
	Database "go-postgres/config"
	ModuleComment "go-postgres/models"

	// Postgres Driver
	_ "github.com/lib/pq"
	"github.com/prometheus/common/log"
)

//------------------------- Controlling Services Between Repository and Database  ----------------

// GetAllModules all
func GetAllComments(id string) ([]ModuleComment.Comment, error) {

	db := Database.CreateConnection()
	defer db.Close()

	var comments []ModuleComment.Comment

	sqlStatement := `SELECT * from app_comment where product_id =$1`
	rows, err := db.Query(sqlStatement, id)

	if err != nil {
		log.Errorln("Unable to execute the query. %v", err)
	}

	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var comment ModuleComment.Comment

		err = rows.Scan(&comment.Comid, &comment.Autor, &comment.Pid, &comment.Comment)

		if err != nil {
			log.Errorln("Unable to scan the row. %v", err)
		}

		comments = append(comments, comment)

	}

	return comments, err
}
