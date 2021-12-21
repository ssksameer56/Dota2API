package handlers

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/ssksameer56/Dota2API/utils"
)

type FavouritesHandler struct {
	mysqlConn       utils.SqlConnection
	FavouritesTable string
}

var mysqlhandler FavouritesHandler

//Queries the favourite heros of a certain user
func (handler *FavouritesHandler) QueryFavouritesOfAUser(pctx context.Context, userID int) ([]int, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE USERID = %d", handler.FavouritesTable, userID)
	ctx, cancel := context.WithCancel(pctx)
	defer cancel()
	data, err := mysqlhandler.mysqlConn.QueryFavourites(ctx, query)
	if err != nil {
		utils.LogInfo("Error in getting data from database", "QueryFavourites")
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

//Marks favourite hero of certain user
func (handler *FavouritesHandler) MarkFavouritesForAUser(pctx context.Context, userID int, heroIDs []int) (bool, error) {
	formattedIDs := []string{}
	for _, id := range heroIDs {
		formattedIDs = append(formattedIDs, fmt.Sprintf("%d", id))
	}
	data := strings.Join(formattedIDs, ",")
	query := fmt.Sprintf("INSERT INTO %s VALUES (%d, '%s')", handler.FavouritesTable, userID, data)
	ctx, cancel := context.WithCancel(pctx)
	defer cancel()
	res, err := mysqlhandler.mysqlConn.ModifyFavourites(ctx, query)
	if err != nil {
		return false, err
	}
	return res > 0, nil
}

func (handler *FavouritesHandler) UpdateFavouritesForAUser(pctx context.Context, userID int, heroIDs []int) (bool, error) {
	formattedIDs := []string{}
	for _, id := range heroIDs {
		formattedIDs = append(formattedIDs, fmt.Sprintf("%d", id))
	}
	data := strings.Join(formattedIDs, ",")
	query := fmt.Sprintf("UPDATE %s SET HeroID='%s' where UserID=%d", handler.FavouritesTable, data, userID)
	ctx, cancel := context.WithCancel(pctx)
	defer cancel()
	res, err := mysqlhandler.mysqlConn.ModifyFavourites(ctx, query)
	if err != nil {
		return false, err
	}
	return res > 0, nil
}

func (handler *FavouritesHandler) GetNextUserID(pctx context.Context) (int, error) {
	query := fmt.Sprintf("SELECT MAX(UserID),HeroIDs FROM %s", handler.FavouritesTable)
	ctx, cancel := context.WithCancel(pctx)
	defer cancel()
	data, err := mysqlhandler.mysqlConn.QueryFavourites(ctx, query)
	if err != nil {
		utils.LogInfo("Error in getting max ID from database", "GetNextUserID")
		return -1, err
	}
	row := (*data)[0]
	id := row.UserID
	return id + 1, nil
}
