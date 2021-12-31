package tests

import (
	"context"
	"testing"

	"github.com/ssksameer56/Dota2API/database"
	"github.com/ssksameer56/Dota2API/handlers"
)

var connString string = "root:sameer123@tcp(127.0.0.1:3306)/DotaDatabase"
var driver string = "mysql"
var DatabaseName string = "DotaDatabase"

func TestQueryFavouritesWhereDataExists(t *testing.T) {
	conn := database.SqlConnection{
		ConnectionString: connString,
		DatabaseName:     DatabaseName,
		DriverName:       driver,
	}

	fh := handlers.FavouritesHandler{
		MysqlConn:       conn,
		FavouritesTable: "Favourites",
	}
	data, err := fh.QueryFavouritesOfAUser(context.Background(), -1)
	if err != nil {
		t.Errorf("couldn't get favourites:" + err.Error())
		t.FailNow()
	}
	if len(data) != 3 {
		t.Fail()
	}
	if data[0] != 122 {
		t.Fail()
	}
}

func TestQueryFavouritesWhereDataDoesntExists(t *testing.T) {
	conn := database.SqlConnection{
		ConnectionString: connString,
		DatabaseName:     DatabaseName,
		DriverName:       driver,
	}

	fh := handlers.FavouritesHandler{
		MysqlConn:       conn,
		FavouritesTable: "Favourites",
	}
	data, err := fh.QueryFavouritesOfAUser(context.Background(), 0)
	if err != nil {
		t.Errorf("couldn't get favourites:" + err.Error())
		t.FailNow()
	}
	if len(data) != 0 {
		t.FailNow()
	}
}

func TestMarkHeroesToFavourites(t *testing.T) {
	conn := database.SqlConnection{
		ConnectionString: connString,
		DatabaseName:     DatabaseName,
		DriverName:       driver,
	}

	fh := handlers.FavouritesHandler{
		MysqlConn:       conn,
		FavouritesTable: "Favourites",
	}
	res, err := fh.MarkFavouritesForAUser(context.Background(), -2, []int{1, 2, 3, 4})
	if err != nil {
		t.Errorf("couldn't get favourites:" + err.Error())
		t.FailNow()
	}
	if !res {
		t.FailNow()
	}
	data, err := fh.QueryFavouritesOfAUser(context.Background(), -2)
	if err != nil {
		t.Errorf("couldn't get favourites:" + err.Error())
	}
	if len(data) != 4 {
		t.FailNow()
	}
}
