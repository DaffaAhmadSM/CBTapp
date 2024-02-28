package main

import (
	"github.com/DaffaAhmadSM/CBTapp/internal/http/database"
	"github.com/DaffaAhmadSM/CBTapp/internal/http/routers"
	"github.com/joho/godotenv"
	"net/http"
)

func main() {
	// LOAD .ENV FILE
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// INIT DATABASE
	db := database.PostgresInit()
	defer func() {
		debe, _ := db.DB()
		err := debe.Close()
		if err != nil {
			panic(err)
		}
	}()
	// INIT SERVER
	mux := routers.RouterInit(db)
	serverInit := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	err = serverInit.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
