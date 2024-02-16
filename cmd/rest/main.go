package main

import (
	"github.com/DaffaAhmadSM/CBTapp/internal/http/database"
	"github.com/DaffaAhmadSM/CBTapp/internal/http/routers"
	"github.com/joho/godotenv"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	db := database.PostgresInit()

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
