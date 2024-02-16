package database

import (
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func PostgresInit() *gorm.DB {
	var DB *gorm.DB
	var err error
	DB, err = gorm.Open(postgres.New(postgres.Config{
		DriverName: "postgres",
		DSN:        os.Getenv("POSTGRES_URL"),
	}))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	log.Println("Database connection established.")

	//defer func() {
	//	connect, _ := DB.DB()
	//	err := connect.Close()
	//	if err != nil {
	//		panic("Failed to close database connection.")
	//	}
	//}()

	return DB
}
