package models

import (
	"fmt"
	"packs/db"
)

type Customer struct {
	IdCustomer     int64
	Name           string
	Address        string
	PhoneNumber    string
	PhoneNumberAux string
}

type Customers []Customer

const SchemaCustomer string = `CREATE TABLE Customers(
	idCustomer int(11) unsigned not null primary key auto_increment,
	name varchar(50) not null,
	address varchar(100) not null,
	number1 int(11) not null,
	number2 int(11),
	current_Data TIMESTAMP default current_timestamp
);`

//FUNC TO BUILD THE STRUCT OF CUSTOMER
func buildStructCustomer(name, address, phone1, phone2 string) *Customer {
	cust1 := &Customer{
		Name:           name,
		Address:        address,
		PhoneNumber:    phone1,
		PhoneNumberAux: phone2,
	}

	return cust1
}

//METHOD OF CUSTOMER TO INSERT THE CUSTOMER ON THE DB
func (c *Customer) insertCustomer() {
	sql := "INSERT customers SET name=?, address=?, number1=?, number2=?"
	rows, err := db.Exec(sql, c.Name, c.Address, c.PhoneNumber, c.PhoneNumberAux)
	if err != nil {
		panic(err)
	}
	c.IdCustomer, _ = rows.LastInsertId()
	fmt.Println("Registro Insertado")
}

//FUNC TO MANAGE THE FUNCTIONS TO INSERT THE CUSTOMER
func CreateCustomer(name, address, phone1, phone2 string) *Customer {
	customer := buildStructCustomer(name, address, phone1, phone2)
	customer.insertCustomer()
	return customer
}

//FUNC TO LIST THE CUSTOMERS
func ListCustomer() Customers {
	sql := "SELECT idCustomer, name, address, number1, number2 From customers"
	customers := Customers{}
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		customer := Customer{}
		//WE scan in the object of Customer the data that we recover of the DB
		rows.Scan(&customer.IdCustomer, &customer.Name, &customer.Address, &customer.PhoneNumber, &customer.PhoneNumberAux)
		customers = append(customers, customer)
	}
	return customers
}

func SearchCustomerById(idCustomer int) Customer {
	sql := "Select idCustomer, name, address, number1,number2 from Customers where idCustomer=?"
	rows, err := db.Query(sql, idCustomer)
	customer := Customer{}
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		rows.Scan(&customer.IdCustomer, &customer.Name, &customer.Address, &customer.PhoneNumber, &customer.PhoneNumberAux)

	}
	return customer
}

func (cu *Customer) DeleteCustomerById(idCustomer int) {
	sql := "Delete from customers where idCustomer=?"
	_, err := db.Query(sql, idCustomer)
	if err != nil {
		panic(err)
	}
	mssg := fmt.Sprintf("Elemento con id %d eliminado", idCustomer)
	fmt.Println(mssg)
}

func (cu *Customer) UpdatedCustomer(idCustomer int) Customer {
	customer := SearchCustomerById(idCustomer)
	sql := "UPDATE customers SET name=?, address=?, number1=?, number2=?"
	_, err := db.Exec(sql, cu.Name, cu.Address, cu.PhoneNumber, cu.PhoneNumberAux)
	if err != nil {
		panic(err)
	}
	return customer
}
