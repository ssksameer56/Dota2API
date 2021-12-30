package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/ssksameer56/Dota2API/models/database"
	"github.com/ssksameer56/Dota2API/utils"

	_ "github.com/go-sql-driver/mysql"
)

type SqlConnection struct {
	ConnectionString string
	DatabaseName     string
	DriverName       string
}

func InitalizeDBConnection(connString string, dbName string) *SqlConnection {
	return &SqlConnection{
		DatabaseName:     dbName,
		ConnectionString: connString,
		DriverName:       "mysql",
	}
}

//Returns rows from the Favourites Table
func (conn *SqlConnection) QueryFavouritesTable(pctx context.Context, dbConn *sql.DB, query string) (*[]database.DBModelFavourites, error) {

	ctx, cancel := context.WithTimeout(pctx, time.Minute)
	defer cancel()
	rows, err := dbConn.QueryContext(ctx, query)
	if err != nil {
		utils.LogError("error reading from database: "+err.Error(), "QueryFavouritesTable")
	}
	result := []database.DBModelFavourites{}
	for rows.Next() {
		newRow := database.DBModelFavourites{}
		err := rows.Scan(&newRow)
		if err != nil {
			utils.LogError("erro reading entry: "+err.Error(), "QueryFavouritesTable")
			return nil, err
		}
		result = append(result, newRow)
	}
	utils.LogInfo(fmt.Sprintf("%d read from database complete", len(result)), "QueryFavouritesTable")
	return &result, nil
}

//Function to update on the Favourites Table
func (conn *SqlConnection) ModifyFavouritesTable(pctx context.Context, dbConn *sql.DB, query string) (int64, error) {
	ctx, cancel := context.WithTimeout(pctx, time.Minute)
	defer cancel()
	result, err := dbConn.ExecContext(ctx, query)
	if err != nil {
		utils.LogError("error updating database: "+err.Error(), "ModifyFavouritesTable")
		return -1, err
	}
	var rows int64
	rows, _ = result.RowsAffected()
	utils.LogInfo("Executed the query: "+query, "ModifyFavouritesTable")
	return rows, nil
}
