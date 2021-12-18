package handlers

import "github.com/ssksameer56/Dota2API/utils"

var DatabaseConnection utils.SqlConnection

func init() {
	DatabaseConnection := utils.InitalizeDBConnection("test", "test")
}
