package main

import (
	"fmt"

	"github.com/cocolabo/hands-on-go/db"
	"github.com/cocolabo/hands-on-go/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully connected to database.")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}
