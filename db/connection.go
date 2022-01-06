package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const connection = "root:@tcp(localhost:3306)/goweb"

var db *sql.DB

func Open() {
	con, err := sql.Open("mysql", connection)
	if err != nil {
		panic(err)
	}
	fmt.Println("connection succesful")
	db = con
}

func Ping() {
	err := db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("connection active")
}

func Close() {
	err := db.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println("connection close")
}

func ExistsTable(nameTable string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", nameTable)
	result, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	return result.Next()
}

func InsertTable(schema, nameTable string) {
	if !ExistsTable(nameTable) {
		result, err := db.Exec(schema)
		if err != nil {
			panic(err)
		}
		fmt.Println("Table created: ", result)
	} else {
		errMsg := fmt.Sprintf("Tabla %s existente", nameTable)
		fmt.Println(errMsg)
	}

}

func AddForeignKey(nameTable, nameConstraint, nameForeignKey, nameTablePrimary, namePrimaryKey string) {
	sql := fmt.Sprintf("alter table %s add constraint %s foreign key (%s) references %s (%s);", nameTable, nameConstraint, nameForeignKey, nameTablePrimary, namePrimaryKey)

	err, result := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	fmt.Println("Foreign Key Succeful: ", result)
}

func Truncate(nameTable string) {
	sql := fmt.Sprintf("TRUNCATE TABLE %s", nameTable)
	err, _ := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	sucMsg := fmt.Sprintf("Datos de la tabla %s eliminados", nameTable)
	fmt.Println(sucMsg)
}
