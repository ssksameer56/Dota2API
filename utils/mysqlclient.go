package utils

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/ssksameer56/Dota2API/models/database"

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

func (conn *SqlConnection) QueryFavourites(pctx context.Context, query string) (*[]database.DBModelFavourites, error) {
	dbConn, err := sql.Open(conn.driverName, conn.connectionString)
	dbConn.Close()
	if err != nil {
		LogError("cannot open connection to DB: "+err.Error(), "QueryFavourites")
	}
	ctx, cancel := context.WithTimeout(pctx, time.Minute)
	defer cancel()
	rows, err := dbConn.QueryContext(ctx, query)
	if err != nil {
		LogError("erro reading from database: "+err.Error(), "QueryFavourites")
	}
	result := []database.DBModelFavourites{}
	for rows.Next() {
		newRow := database.DBModelFavourites{}
		err := rows.Scan(&newRow)
		if err != nil {
			LogError("erro reading entry: "+err.Error(), "QueryFavourites")
			return nil, err
		}
		result = append(result, newRow)
	}
	LogInfo(fmt.Sprintf("%d read from database complete", len(result)), "QueryFavourites")
	return &result, nil
}

//Function to update on the Favourites Table
func (conn *SqlConnection) ModifyFavourites(pctx context.Context, query string) (int64, error) {
	dbConn, err := sql.Open(conn.driverName, conn.connectionString)
	dbConn.Close()
	if err != nil {
		LogError("cannot open connection to DB: "+err.Error(), "ModifyFavourites")
		return -1, err
	}
	ctx, cancel := context.WithTimeout(pctx, time.Minute)
	defer cancel()
	result, err := dbConn.ExecContext(ctx, query)
	if err != nil {
		LogError("erro updating database: "+err.Error(), "ModifyFavourites")
		return -1, err
	}
	var rows int64
	rows, _ = result.RowsAffected()
	LogInfo("Executed the query: "+query, "ModifyFavourites")
	return rows, nil
}
