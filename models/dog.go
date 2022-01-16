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

type Dogs []Dog

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
		fmt.Println("There's no such ID Customer")
	}
	dog.IdCustomer, _ = result.LastInsertId()
	fmt.Println("Dog Save")
}

func CreateDog(name, age, breed string, idCustomer int64) *Dog {

	dog := buildDog(name, age, breed, idCustomer)
	dog.insertDog()
	return dog
}

func ListDogs() Dogs {
	sql := "Select idDog, name, age, breed, idCustomer from dogs"
	rows, err := db.Query(sql)
	dogs := Dogs{}

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		dog := Dog{}
		rows.Scan(&dog.IdDog, &dog.Name, &dog.Age, &dog.Breed, &dog.IdCustomer)
		dogs = append(dogs, dog)
	}

	return dogs
}

func existsDog(dog Dog) bool {
	return dog.IdDog != 0 && dog.Name != ""
}

func (dog *Dog) DeleteDogByID(idDog int) {
	sql := "Delete from dogs where idDog=?"
	_, err := db.Query(sql, idDog)
	if err != nil {
		fmt.Println("Id inexistente")
	}
	fmt.Println("Registro eliminado")

}

func SearchDogById(idDog int) (Dog, bool) {
	sql := "Select idDog, name,age,breed, idCustomer from dogs where idDog=?"
	row, err := db.Query(sql, idDog)
	if err != nil {
		panic(err)
	}
	dog := Dog{}
	for row.Next() {
		row.Scan(&dog.IdDog, &dog.Name, &dog.Age, &dog.Breed, &dog.IdCustomer)
	}
	existsDog := existsDog(dog)
	return dog, existsDog
}

func SerchDogByName(nameDog string) Dog {
	sql := "Select idDog, name,age,breed, idCustomer from dogs where name=?"
	row, err := db.Query(sql, nameDog)
	if err != nil {
		panic(err)
	}
	dog := Dog{}
	for row.Next() {
		row.Scan(&dog.IdDog, &dog.Name, &dog.Age, &dog.Breed, &dog.IdCustomer)
	}
	return dog
}

func (dog *Dog) UpdatedDog(idDog int) *Dog {
	SearchDogById(idDog)
	sql := "UPDATE dogs set name=?, age=?, breed=? where idDog=?"
	_, erro := db.Exec(sql, dog.Name, dog.Age, dog.Breed, idDog)
	if erro != nil {
		panic(erro)
	}

	return dog
}
