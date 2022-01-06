package main

import (
	"packs/db"
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

	db.Close()
}
