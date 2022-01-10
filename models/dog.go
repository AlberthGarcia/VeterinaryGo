package models

import (
	"fmt"
	"packs/db"
)

type Dog struct {
	IdDog      int64
	Name       string
	Age        string
	Breed      string
	IdCustomer int64
}

const SchemaDog string = `CREATE TABLE dogs(
	idDog int(11) unsigned not null primary key auto_increment,
	name varchar(50) not null,
	age varchar(20) not null,
	breed varchar(30) not null,
	idCustomer int(11) unsigned not null,
	current_Data TIMESTAMP default current_timestamp
);`

func buildDog(name, age, breed string, idCustomer int64) *Dog {
	dog := &Dog{
		Name:       name,
		Age:        age,
		Breed:      breed,
		IdCustomer: idCustomer,
	}

	return dog
}

func (dog *Dog) insertDog() {
	sql := "INSERT dogs SET name=?, age=?, breed=?, idCustomer=?"
	result, err := db.Exec(sql, dog.Name, dog.Age, dog.Breed, dog.IdCustomer)
	if err != nil {
		panic(err)
	}
	fmt.Println("Registro insertado", result)
}

func CreateDog(name, age, breed string, idCustomer int64) *Dog {

	dog := buildDog(name, age, breed, idCustomer)
	dog.insertDog()
	return dog
}
