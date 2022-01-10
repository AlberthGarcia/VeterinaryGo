package main

import (
	"fmt"
	"packs/db"
	"packs/models"
)

func dataBase() {
	db.Open()
	var number int
	fmt.Println("-----------OPTIONS--------")
	fmt.Println("Verify Connection----1")
	fmt.Println("Create table----2")
	fmt.Println("Empty table----3")
	fmt.Println("Delete table----4")
	fmt.Println("Create Foreign Key----5")
	fmt.Println("Delete foreign key----6")
	fmt.Println("Close Connection----7")
	fmt.Scanln(&number)

	switch {
	case number == 1:
		db.Ping()
	case number == 2:
		db.InsertTable(models.SchemaVet, "veterinary")
	case number == 3:
		db.Truncate("Customers")
	case number == 4:
		db.Delete("veterinary")
	case number == 5:
		db.AddForeignKey("veterinary", "fk_Customer", "idCustomer", "customers", "idCustomer")
	case number == 6:
		db.DeleteForeignKey("veterinary", "fk_Customer")
	case number == 7:
		db.Close()
	default:
		fmt.Print("Invalid option")
	}
	db.Close()
}

func customer() {
	var number int
	db.Open()
	fmt.Println("-----------OPTIONS--------")
	fmt.Println("Insert Customer----1")
	fmt.Println("List Customers----2")
	fmt.Println("Search Customer ID----3")
	fmt.Println("Search Customer Name----4")
	fmt.Println("Delete Customer ID----5")
	fmt.Println("Update Customer----6")
	fmt.Scanln(&number)
	switch {
	case number == 1:
		customer := models.CreateCustomer("Roberto", "Calle coloquial", "2461231231", "2461283515")
		fmt.Println(customer)
	case number == 2:
		listCustomer := models.ListCustomer()
		fmt.Println("List of Customers", listCustomer)
	case number == 3:
		customer := models.SearchCustomerById(4)
		fmt.Println(customer)
	case number == 4:
		customer := models.SerchCustomerByName("Alberth")
		fmt.Println(customer)
	case number == 5:
		// customer.DeleteCustomerById(1)

	case number == 6:
		customer := models.SearchCustomerById(2)
		fmt.Println("Customer to updated", customer)
		customer.Name = "pepe"
		customer.Address = "asdasdfasdf"
		customer.PhoneNumber = "12341234"
		customerUpdated := customer.UpdatedCustomer(2)
		fmt.Println("Customer updated", customerUpdated)

	default:
		fmt.Print("Invalid option")
	}

	db.Close()
}

func main() {
	var number int
	fmt.Println("-----------OPTIONS--------")
	fmt.Println("Functions about database----1")
	fmt.Println("Functions about customers----2")
	fmt.Println("Functions about Veterinary----3")
	fmt.Scanln(&number)

	switch {
	case number == 1:
		dataBase()
	case number == 2:
		customer()
	default:
		fmt.Print("Invalid option")
	}

}
