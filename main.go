package main

import (
	"context"
	"database/sql"

	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/recommendation-system/controller"
	middlewares "github.com/recommendation-system/middleware"
	"github.com/recommendation-system/model"
	"github.com/recommendation-system/router"
	corn "github.com/robfig/cron/v3"
)

func main() {
	ctx := context.Background()
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	dbinfo := os.Getenv("POSTGERS_API_LINE")
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal("There is an err in the conaction ")
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Printf("err type: %T\n", err)
		log.Printf("err is: %v\n", err)
		log.Fatal("There is an err in Ping")
	}
	log.Println("Database connection successful")

	add := ":" + os.Getenv("PORT")
	st_ore := model.NewStore(db)
	log.Println("Store has been opened")

	app := middlewares.Application{
		Address: add,
		Storge:  st_ore,
	}

	ctrlApp := &controller.Application{
		Application: app,
	}

	Traker := controller.TrakerCron{
		Cron:        corn.New(),
		Application: &app,
	}
	
	Traker.UsersRatingsCSV(ctx)
	Traker.UsersCSV(ctx)
	Traker.Cron.Start()

	cntrolObj := &router.Control{
		Controller: ctrlApp,
	}
	appRouter := &router.Application{
		Application: app,
	}

	mux := cntrolObj.Moul()

	if err := appRouter.Run(mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
