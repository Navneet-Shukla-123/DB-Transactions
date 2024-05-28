package main

import (
	"db-transactions/database"
	"log"
)

func init() {
	database.ConnectToDB()
}

func main() {

	err:=database.InsertIntoTableWithTX()
	if err!=nil{
		log.Println("Error in database operation with transaction: ",err)
		return
	}

	//database.InsertIntoTable()

	log.Println("Database operation is successfull.")

}
