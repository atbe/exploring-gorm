package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/matryer/try"
	postgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DB_SERVICE_DIALECT           = "DB_SERVICE_DIALECT"
	DB_SERVICE_CONNECTION_STRING = "DB_SERVICE_CONNECTION_STRING"
)

var db *gorm.DB

func init() {
	connectionString, ok := os.LookupEnv(DB_SERVICE_CONNECTION_STRING)
	if !ok {
		panic(DB_SERVICE_CONNECTION_STRING + " environment variable not set")
	}

	const attempts = 5
	err := try.Do(func(attempt int) (bool, error) {
		fmt.Printf("Connecting to db, attempt %v\n", attempt)
		var err error
		db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
		fmt.Printf("failed to connect database, attempt #%v", attempt)
		fmt.Println(fmt.Errorf("error: %w", err))
		if err == nil {
			return true, nil
		}
		sleepTime := attempt * attempts * 2
		time.Sleep(time.Second * time.Duration(sleepTime))
		return attempt < attempts, err
	})
	if err != nil {
		panic("failed to connect database")
	}

	log.Print("Successfully connected to postgres!")
}

func main() {
}
