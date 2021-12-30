package handlers

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/ssksameer56/Dota2API/database"
	"github.com/ssksameer56/Dota2API/utils"
)

type FavouritesHandler struct {
	mysqlConn       database.SqlConnection
	FavouritesTable string
}

var mysqlhandler FavouritesHandler

//Queries the favourite heros of a certain user
func (handler *FavouritesHandler) QueryFavouritesOfAUser(pctx context.Context, userID int) ([]int, error) {
	ctx, cancel := context.WithCancel(pctx)
	defer cancel()

	dbConn, err := sql.Open(handler.mysqlConn.DriverName, handler.mysqlConn.ConnectionString)
	if err != nil {
		utils.LogError("cannot open connection to DB: "+err.Error(), "QueryFavouritesTable")
	}
	defer dbConn.Close()

	query := fmt.Sprintf("SELECT * FROM %s WHERE USERID = %d", handler.FavouritesTable, userID)
	data, err := mysqlhandler.mysqlConn.QueryFavouritesTable(ctx, dbConn, query)
	if err != nil {
		utils.LogInfo("Error in getting data from database", "QueryFavouritesOfAUser")
		return []int{}, err
	}
	row := (*data)[0]
	rowData := strings.Split(row.HeroIDs, ",")
	heroIDs := []int{}
	for _, id := range rowData {
		id_int, _ := strconv.Atoi(id)
		heroIDs = append(heroIDs, id_int)
	}
	return heroIDs, nil
}

//inserts favourite hero of a new user
func (handler *FavouritesHandler) InsertFavouritesForAUser(pctx context.Context, userID int, heroIDs []int) (bool, error) {
	formattedIDs := []string{}
	if len(heroIDs) == 0 {
		utils.LogInfo("no hero ids provided", "InsertFavouritesForAUser")
		return false, nil
	}
	for _, id := range heroIDs {
		formattedIDs = append(formattedIDs, fmt.Sprintf("%d", id))
	}
	data := strings.Join(formattedIDs, ",")

	ctx, cancel := context.WithCancel(pctx)
	defer cancel()

	dbConn, err := sql.Open(handler.mysqlConn.DriverName, handler.mysqlConn.ConnectionString)
	if err != nil {
		utils.LogError("cannot open connection to DB: "+err.Error(), "QueryFavouritesTable")
	}
	defer dbConn.Close()

	query := fmt.Sprintf("INSERT INTO %s VALUES (%d, '%s')", handler.FavouritesTable, userID, data)
	res, err := mysqlhandler.mysqlConn.ModifyFavouritesTable(ctx, dbConn, query)
	if err != nil {
		utils.LogError("error inserting data: "+err.Error(), "InsertFavouritesForAUser")
		return false, err
	}
	utils.LogInfo(fmt.Sprintf("Inserted %d rows for %d userID", res, userID), "InsertFavouritesForAUser")
	return res > 0, nil
}

//Updates the favourites of a certain userID
func (handler *FavouritesHandler) UpdateFavouritesForAUser(pctx context.Context, userID int, heroIDs []int) (bool, error) {
	formattedIDs := []string{}
	for _, id := range heroIDs {
		formattedIDs = append(formattedIDs, fmt.Sprintf("%d", id))
	}
	data := strings.Join(formattedIDs, ",")

	ctx, cancel := context.WithCancel(pctx)
	defer cancel()

	dbConn, err := sql.Open(handler.mysqlConn.DriverName, handler.mysqlConn.ConnectionString)
	if err != nil {
		utils.LogError("cannot open connection to DB: "+err.Error(), "QueryFavouritesTable")
	}
	defer dbConn.Close()

	query := fmt.Sprintf("UPDATE %s SET HeroID='%s' where UserID=%d", handler.FavouritesTable, data, userID)
	res, err := mysqlhandler.mysqlConn.ModifyFavouritesTable(ctx, dbConn, query)
	if err != nil {
		utils.LogError("error inserting data: "+err.Error(), "UpdateFavouritesForAUser")
		return false, err
	}
	utils.LogInfo(fmt.Sprintf("Updated %d rows for %d userID", res, userID), "UpdateFavouritesForAUser")
	return res > 0, nil
}

//Marks favourites of the user - appends if favourites already exist or inserts new records
func (handler *FavouritesHandler) MarkFavouritesForAUser(pctx context.Context, userID int, newHeroIDs []int) (bool, error) {
	ctx, cancel := context.WithCancel(pctx)
	defer cancel()
	existingData, err := handler.QueryFavouritesOfAUser(ctx, userID)
	if err != nil {
		return false, err
	}
	if len(existingData) == 0 {
		_, err := handler.InsertFavouritesForAUser(ctx, userID, newHeroIDs)
		if err != nil {
			return false, err
		}
	} else {
		updatedIDs := append(existingData, newHeroIDs...)
		_, err := handler.UpdateFavouritesForAUser(ctx, userID, updatedIDs)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

//Gets the next UserID
func (handler *FavouritesHandler) GetNextUserID(pctx context.Context) (int, error) {
	query := fmt.Sprintf("SELECT MAX(UserID),HeroIDs FROM %s", handler.FavouritesTable)
	ctx, cancel := context.WithCancel(pctx)
	defer cancel()

	dbConn, err := sql.Open(handler.mysqlConn.DriverName, handler.mysqlConn.ConnectionString)
	if err != nil {
		utils.LogError("cannot open connection to DB: "+err.Error(), "QueryFavouritesTable")
	}
	defer dbConn.Close()

	data, err := mysqlhandler.mysqlConn.QueryFavouritesTable(ctx, dbConn, query)
	if err != nil {
		utils.LogInfo("Error in getting max ID from database", "GetNextUserID")
		return -1, err
	}
	row := (*data)[0]
	id := row.UserID
	return id + 1, nil
}
