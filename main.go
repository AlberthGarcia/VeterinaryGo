package main

import (
	"packs/db"
)

func main() {
	db.Open()
	db.Ping()

	//FUNCTIONS ABOUT DB
	// db.InsertTable(models.SchemaCustomer, "Customers")
	// db.InsertTable(models.SchemaVet)
	// db.Truncate("Customers")
	// db.AddForeignKey("veterinary", "fk_Customer", "idCustomer", "customers", "idCustomer")

	//FUNCTIONS ABOUT MODELS

	db.Close()
}
