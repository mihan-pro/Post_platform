// Database Configuration
package Database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/prometheus/common/log"
)

func CreateConnection() *sql.DB {

	if os.Getenv("MOD") != "develop" {
		var (
			host     = os.Getenv("DBHOST")
			port     = 5432
			user     = os.Getenv("DBUSER")
			password = os.Getenv("DBPASS")
			dbname   = os.Getenv("DBNAME")
		)
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)
		db, err := sql.Open("postgres", psqlInfo)

		if err != nil {
			//panic(err)
			log.Errorln(err)
		}

		// check the connection
		err = db.Ping()

		if err != nil {
			log.Errorln(err)
		}
		//log.Info("Successfully connected!", db)

		// return the connection
		return db
	} else {
		const (
			host     = "localhost"
			port     = 5432
			user     = "developer"
			password = "developer"
			dbname   = "app"
		)
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)
		db, err := sql.Open("postgres", psqlInfo)

		if err != nil {
			//panic(err)
			log.Errorln(err)
		}

		// check the connection
		err = db.Ping()

		if err != nil {
			log.Errorln(err)
		}
		//log.Info("Successfully connected!", db)

		// return the connection
		return db
	}

}
