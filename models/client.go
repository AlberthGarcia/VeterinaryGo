package models

type Customer struct {
	IdCustomer     int64
	Name           string
	Address        string
	PhoneNumber    int
	PhoneNumberAux int
}

const SchemaCustomer string = `CREATE TABLE Customers(
	idCustomer int(11) unsigned not null primary key auto_increment,
	name varchar(50) not null,
	address varchar(100) not null,
	number1 int(11) not null,
	numer2 int(11),
	current_Data TIMESTAMP default current_timestamp
);`
