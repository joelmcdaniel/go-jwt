package driver

import (
	"database/sql"
	"jwt-auth-restapi/utils"
	"os"

	"github.com/lib/pq"
)

var db *sql.DB

// ConnectDB ...
func ConnectDB() *sql.DB {
	pgURL, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))

	utils.LogFatal(err)

	db, err = sql.Open("postgres", pgURL)

	utils.LogFatal(err)

	err = db.Ping()

	utils.LogFatal(err)

	return db
}
