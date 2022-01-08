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
	fmt.Println("Lista de customers a modificar", listCustomer)
	customer := models.SearchCustomerById(2)
	// fmt.Println(customer)
	// customer = models.SearchCustomerById(1)
	// fmt.Println(customer)
	fmt.Println("Customer a Modificar", customer)

	customer.Name = "asd"
	customer.Address = "asd"
	customer.PhoneNumber = "ads"
	customer.UpdatedCustomer(2)
	listCustomer = models.ListCustomer()
	fmt.Println("Lista de Customers nueva", listCustomer)
	customer.DeleteCustomerById(1)

	db.Close()
}
