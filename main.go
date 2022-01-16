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
		db.InsertTable(models.SchemaDog, "dogs")
	case number == 3:
		db.Truncate("Customers")
	case number == 4:
		db.Delete("veterinary")
	case number == 5:
		db.AddForeignKey("dogs", "fk_Customer", "idCustomer", "customers", "idCustomer")
	case number == 6:
		db.DeleteForeignKey("dogs", "fk_Customer")
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
	fmt.Println("Return to the main menu----0")
	fmt.Scanln(&number)
	switch {
	case number == 1:
		var name, AP, AM, nameCom, address, phone1, phone2 string
		fmt.Println("Name")
		fmt.Scan(&name)

		fmt.Println("First Name")
		fmt.Scan(&AP)

		fmt.Println("last Name")
		fmt.Scan(&AM)

		nameCom = fmt.Sprintf("%s %s %s", name, AP, AM)

		fmt.Println("address")
		fmt.Scan(&address)

		fmt.Println("cellphone")
		fmt.Scan(&phone1)

		fmt.Println("emergency phone (optional)")
		fmt.Scan(&phone2)

		customer := models.CreateCustomer(nameCom, address, phone1, phone2)
		fmt.Println(customer)
	case number == 2:
		listCustomer := models.ListCustomer()
		fmt.Println("List of Customers", listCustomer)
	case number == 3:
		var id int
		fmt.Println("Write the ID to delete")
		fmt.Scan(&id)
		customer, exists := models.SearchCustomerById(id)
		if exists {
			fmt.Println(customer)
		} else {
			fmt.Println("There is no such ID")
		}
	case number == 4:
		customer, exists := models.SerchCustomerByName("Alberth")
		if exists {
			fmt.Println(customer)
		} else {
			fmt.Println("There is no such name")
		}
	case number == 5:
		var id int
		fmt.Println("Write the ID to delete")
		fmt.Scan(&id)
		customer, exists := models.SearchCustomerById(id)
		if exists {
			fmt.Println("Data to delete", customer)
			customer.DeleteCustomerById(id)
		} else {
			fmt.Println("There is no such ID")
		}

	case number == 6:
		var id int
		fmt.Println("Write the ID to updated")
		fmt.Scan(&id)
		customer, exists := models.SearchCustomerById(id)
		fmt.Println("Customer to updated", customer)

		if exists {
			var fname, AP, AM, name, address, phone1, phone2 string
			fmt.Println("Name")
			fmt.Scan(&fname)

			fmt.Println("First Name")
			fmt.Scan(&AP)

			fmt.Println("last Name")
			fmt.Scan(&AM)

			name = fmt.Sprintf("%s %s %s", fname, AP, AM)
			customer.Name = name

			fmt.Println("Customer address")
			fmt.Scan(&address)
			customer.Address = address

			fmt.Println("Customer number")
			fmt.Scan(&phone1)
			customer.PhoneNumber = phone1

			fmt.Println("Customer number2 (optional)")
			fmt.Scan(&phone2)
			customer.PhoneNumberAux = phone2

			customerUpdated := customer.UpdatedCustomer(id)
			fmt.Println("Customer updated", customerUpdated)
		} else {
			fmt.Println("There is no such Name")
		}
	case number == 0:
		main()
	default:
		fmt.Print("Invalid option")
	}

	db.Close()
}

func dog() {
	var number int
	db.Open()
	fmt.Println("-----------OPTIONS--------")
	fmt.Println("Insert Dog----1")
	fmt.Println("List Dogs----2")
	fmt.Println("Delete Dog----3")
	fmt.Println("Search Dog ID----4")
	fmt.Println("Search Dog name----5")
	fmt.Println("Update Dog----6")
	fmt.Println("Return to the main menu----0")
	fmt.Scanln(&number)
	switch {
	case number == 1:
		var name, age, breed string
		var idCustomer int64
		fmt.Println("Name")
		fmt.Scan(&name)

		fmt.Println("Age")
		fmt.Scan(&age)

		fmt.Println("Breed")
		fmt.Scan(&breed)

		fmt.Println("Id Owner")
		fmt.Scan(&idCustomer)

		dog := models.CreateDog(name, age, breed, idCustomer)
		fmt.Println(dog)

	case number == 2:
		dogs := models.ListDogs()
		fmt.Println(dogs)
	case number == 3:
		var id int
		fmt.Println("ID to delete")
		fmt.Scan(&id)
		dog, exists := models.SearchDogById(id)
		if exists {
			dog.DeleteDogByID(id)
		} else {
			fmt.Println("There is no such ID")
		}
	case number == 4:
		var id int
		fmt.Println("Id Dog")
		fmt.Scan(&id)
		dog, existDog := models.SearchDogById(id)
		if existDog {
			fmt.Println(dog)
		} else {
			fmt.Println("There is no such ID")
		}
	case number == 5:
		var nameDog string
		fmt.Println("name Dog")
		fmt.Scan(&nameDog)
		dog := models.SerchDogByName(nameDog)
		fmt.Println(dog)
	case number == 6:
		var idDog int
		var name, age, breed string
		fmt.Println("Id of the Dog")
		fmt.Scanln(&idDog)

		dog, existDog := models.SearchDogById(idDog)
		if existDog {
			fmt.Println(dog)

			fmt.Println("Name")
			fmt.Scanln(&name)
			dog.Name = name

			fmt.Println("Age")
			fmt.Scanln(&age)
			dog.Age = age

			fmt.Println("Breed")
			fmt.Scanln(&breed)
			dog.Breed = breed

			dog := dog.UpdatedDog(idDog)
			fmt.Println("Updated dog", dog)
		} else {
			fmt.Println("There is no such Id")
		}
	case number == 0:
		main()
	default:
		fmt.Print("Invalid option")
	}

	db.Close()

}

func main() {
	var number int
	fmt.Println("-----------OPTIONS--------")
	fmt.Println("Functions about Database----1")
	fmt.Println("Functions about Customers----2")
	fmt.Println("Functions about Dogs----3")
	fmt.Scanln(&number)

	switch {
	case number == 1:
		dataBase()
	case number == 2:
		customer()
	case number == 3:
		dog()
	default:
		fmt.Print("Invalid option")
	}

}
