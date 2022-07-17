package dbutils

var (
	DBName    string
	TableName string

	PathToDB string
)

var Databases = make(map[string]interface{}) // Databases:=map[string]interface{ "database1":table1{col:val}}

type Field struct {
	Name string
	Val  interface{}
	Type int
}

type Row struct {
	Id     uint
	Fields []Field
}

type Statement struct {
	Type        int
	RowToInsert Row
}

type Table struct {
	TableName string
	NumRows   uint
	Rows      []Row
	// will use pages in teh later section of the dev
	Pages []Page
}

type Page struct {
	Rows []Row
}

type Schema struct {
	Tables []Table
}

type Database struct {
	Name   string
	Tables []Table
}
