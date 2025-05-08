package config

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
)

type ConnectionParams struct {
	Server   string
	Database string
	LoginID  string
	Password string
}

func setupConnection(conn ConnectionParams) (*sql.DB, error) {
	var err error
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s;encrypt=disable", conn.Server, conn.LoginID, conn.Password, conn.Database)

	// Open the database connection
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		return nil, fmt.Errorf("⛔️ Error opening database: %v", err)
	}

	// Set a timeout for the connection
	db.SetConnMaxLifetime(time.Minute * 15)
	db.SetMaxOpenConns(15)
	db.SetMaxIdleConns(5)

	return db, nil
}

func CheckServer(conn ConnectionParams) bool {
	db, err := setupConnection(conn)
	if err != nil {
		return false
	}

	errPing := db.Ping()
	if errPing != nil {
		return false
	}

	return true
}

func ConnectDB(conn ConnectionParams) *sqlx.DB {
	db, err := setupConnection(conn)
	if err != nil {
		log.Fatalf("⛔️ Failed to connect to the database: %v", err)
	}
	// defer db.Close()

	dbx := sqlx.NewDb(db, "sqlserver")
	return dbx
}

func ExecQuery(dbx *sqlx.DB, query string, isReturnData bool) *sqlx.Rows {
	rowsData, err := dbx.Queryx(query)
	if err != nil {
		log.Fatalf("⛔️ ERROR executing query: %v", err)
	}

	if !isReturnData {
		rowsData.Close()
	}

	return rowsData
}
