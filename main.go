package main

import (
	"fmt"
	"packs/db"
	"packs/models"
)

func main() {
	db.Open()
	db.Ping()

	//FUNCTIONS ABOUT DB
	// db.InsertTable(models.SchemaCustomer, "Customers")
	// db.InsertTable(models.SchemaVet, "Veterinary")
	// db.AddForeignKey("veterinary", "fk_Customer", "idCustomer", "customers", "idCustomer")
	// db.DeleteForeignKey("veterinary", "fk_Customer")
	// db.Truncate("Customers")

	//FUNCTIONS ABOUT MODELS
	// customer := models.CreateCustomer("Roberto", "Calle coloquial", "2461231231", "2461283515")
	// fmt.Println(customer)
	listCustomer := models.ListCustomer()
	fmt.Println(listCustomer)
	db.Close()
}
