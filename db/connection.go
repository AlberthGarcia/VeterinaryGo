package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const connection = "root:@tcp(localhost:3306)/goweb"

var db *sql.DB

//FUNC TO OPEN THE CONNECTION WITH THE DB AND ASSIGN THIS CONECCTION TO A VARIABLE
func Open() {
	con, err := sql.Open("mysql", connection)
	if err != nil {
		panic(err)
	}
	fmt.Println("connection succesful")
	db = con
}

//FUNC TO VERIFY A CONNECTION WITH THE DB
func Ping() {
	err := db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("connection active")
}

//FUNC TO CLOSE DE CONNECTION WITH THE DB
func Close() {
	err := db.Close()
	if err != nil {
		panic(err)
	}
}

//FUNCT TO VERIFY IF A TABLE IN THE DB EXISTS
func existsTable(nameTable string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", nameTable)
	result, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	return result.Next()
}

//FUNC TO INSERT A TABLE IF THIS DOES NOT EXISTS
func InsertTable(schema, nameTable string) {
	if !existsTable(nameTable) {
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

//FUNC TO ADD A FOREIGN KEY AND HIS CONSTRAINT IN A TABLE IF THIS DOES NOT HAVE ONE
func AddForeignKey(nameTable, nameConstraint, nameForeignKey, nameTablePrimary, namePrimaryKey string) {
	sql := fmt.Sprintf("ALTER TABLE %s ADD CONSTRAINT %s FOREIGN KEY (%s) REFERENCES %s (%s);", nameTable, nameConstraint, nameForeignKey, nameTablePrimary, namePrimaryKey)
	// fmt.Println(sql)
	result, err := Exec(sql)
	if err != nil {
		panic(err)
	}
	fmt.Println("Foreign Key Succeful: ", result)
}

//FUNC TO TRUNCATE A TABLE OR DELETE ALL REGISTERS IN THIS
func Truncate(tableName string) {
	sql := fmt.Sprintf("TRUNCATE TABLE %s;", tableName)
	_, err := Exec(sql)
	if err != nil {
		panic(err)
	}
	sucMsg := fmt.Sprintf("Datos de la tabla %s eliminados", tableName)
	fmt.Println(sucMsg)
}

//FUNC TO DELETE A TABLE
func Delete(tableName string) {
	sql := fmt.Sprintf("DROP TABLE IF EXISTS %s", tableName)
	_, err := Exec(sql)
	if err != nil {
		panic(err)
	}
	sucMsg := fmt.Sprintf("Tabla %s eliminada", tableName)
	fmt.Println(sucMsg)
}

//FUNCT TO ELIMINATE A FOREGIN KEY IN A TABLE
func DeleteForeignKey(nameTable, nameConstraint string) {
	sql := fmt.Sprintf("ALTER TABLE %s DROP FOREIGN KEY %s;", nameTable, nameConstraint)
	err, _ := Exec(sql)
	if err != nil {
		panic(err)
	}
	sucMsg := fmt.Sprintf("Clave foranea de la tabla %s eliminada", nameTable)
	fmt.Println(sucMsg)
}

//FUNC EXEC
func Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := db.Exec(query, args...)
	if err != nil {
		panic(err)
	}
	return result, err
}

//FUNC QUERY
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		panic(err)
	}
	return rows, err
}
