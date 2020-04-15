package main

import (
	"log"
	"os"
	"strconv"

	"github.com/Kotyarich/find-your-pet/api/handlers"
	"github.com/Kotyarich/find-your-pet/db"
	"github.com/Kotyarich/find-your-pet/managers"
	"github.com/Kotyarich/find-your-pet/store/db/pg"
	"github.com/spf13/viper"
)

func PrepareConfig() {
	debug, err := strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		log.Fatalln(err)
	}
	if debug {
		viper.SetConfigFile("config.yml")
	} else {
		viper.SetConfigFile(os.Args[1])
	}
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Cannot read a config file: %v\n", err)
	}
}

func PrepareHandlerData() *handlers.HandlerData {
	db, err := db.ConnectToDB(viper.GetString("db.host"),
		viper.GetString("db.user"), viper.GetString("db.password"), viper.GetString("db.database"))
	if err != nil {
		log.Fatalf("Error with database: %v\n", err)
	}
	db.SetMaxOpenConns(viper.GetInt("db.max_connections"))
	lostController := pg.NewLostControllerPg(viper.GetInt("lost.itemsPerPage"), db)
	FileController := pg.NewFileControllerPg(db)
	lostAddingManager :=
		managers.NewLostAddingManager(db, lostController, FileController)

	foundController := pg.NewFoundControllerPg(viper.GetInt("found.itemsPerPage"), db)
	foundAddingManager :=
		managers.NewFoundAddingManager(db, foundController, FileController)

	profileController := pg.NewProfileControllerPg(
		viper.GetInt("profile.lost.itemsPerPage"),
		db)
	debug, err := strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		log.Fatalln(err)
	}
	return handlers.NewHandlerData(lostController, FileController,
		lostAddingManager, foundController, foundAddingManager,
		profileController, debug)
}
