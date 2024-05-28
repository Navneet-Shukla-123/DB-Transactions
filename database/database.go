package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func ConnectToDB() {
	username := "posttest"
	password := "test1234"
	host := "localhost"
	port := "5432"
	databaseName := "transactions"

	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		username, password, host, port, databaseName)

	var err error

	db, err = sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to the database!")
}

// InsertIntoTable will use transaction mechanism to insert into two table
func InsertIntoTableWithTX() error {

	tx, err := db.Begin()
	if err != nil {
		log.Println("Error in starting the transaction ", err)
		return err
	}

	query := `Insert into address(email,city,state) values($1,$2,$3);`
	_, err = tx.Exec(query, "c@c.com", "bsb", "up")

	if err != nil {
		log.Println("Error in inserting to address table ", err)
		tx.Rollback()
		return err
	}
	query = `Insert into users(name,email) values($1,$2);`

	_, err = tx.Exec(query, "c", "c@c.com")
	if err != nil {
		log.Println("Error in inserting to users table ", err)
		tx.Rollback()
		return err
	}

	err = tx.Commit()

	if err != nil {
		log.Println("Error  in doing the commit ", err)
		tx.Rollback()
		return err
	}

	return nil

}

func InsertIntoTable() error {

	tx := db

	query := `Insert into address(email,city,state) values($1,$2,$3);`
	_, err := tx.Exec(query, "c@c.com", "bsb", "up")

	if err != nil {
		log.Println("Error in inserting to address table ", err)

		return err
	}
	query = `Insert into users(name,email) values($1,$2);`

	_, err = tx.Exec(query, "c", "c@c.com")
	if err != nil {
		log.Println("Error in inserting to users table ", err)
		return err
	}

	return nil

}
