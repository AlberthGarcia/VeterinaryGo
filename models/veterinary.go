package models

type Veterinary struct {
	IdVet    int64
	Name     string
	Address  string
	IdClient int64
}

const SchemaVet string = `CREATE TABLE veterinary(
	idVet int(11) unsigned not null primary key auto_increment,
	name varchar(50) not null,
	address varchar(100) not null,
	idCustomer int(11) unsigned not null,
	current_Data TIMESTAMP default current_timestamp
);`

func buildVeterinary(name, address string, idCustomer int64) *Veterinary {
	veterinary := &Veterinary{
		Name:     name,
		Address:  address,
		IdClient: idCustomer,
	}
	return veterinary
}
