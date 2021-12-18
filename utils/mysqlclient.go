package utils

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type SqlConnection struct {
	connectionString string
	databaseName     string
	driverName       string
}

func InitalizeDBConnection(connString string, dbName string) *SqlConnection {
	return &SqlConnection{
		databaseName:     dbName,
		connectionString: connString,
		driverName:       "mysql",
	}
}

func (conn *SqlConnection) Query(pctx context.Context, query string, dataObject interface{}) error {
	dbConn, err := sql.Open(conn.driverName, conn.connectionString)
	dbConn.Close()
	if err != nil {
		LogError("cannot open connection to DB: "+err.Error(), "Query")
	}
	ctx, cancel := context.WithTimeout(pctx, time.Minute)
	defer cancel()
	rows, err := dbConn.QueryContext(ctx, query)
	if err != nil {
		LogError("erro reading from database: "+err.Error(), "Query")
	}
	for rows.Next() {
		//TODO: Modify result scan function to be as generic as possible
		err := rows.Scan(&dataObject)
		if err != nil {
			LogError("erro reading entry: "+err.Error(), "Query")
			return err
		}
	}
	LogInfo("Read data from database complete", "Query")
	return nil
}
