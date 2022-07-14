package dbutils

import "fmt"

type Row struct {
	Id       uint
	UserName string
	Email    string
}

type Statement struct {
	Type        int
	RowToInsert Row
}

type Table struct {
	TableName string
	NumRows   uint
	Pages     []Page
}

type Page struct {
	Rows []Row
}

type Schema struct{
	Tables []Table 
}

func RowSlot(table *Table, rowNum uint) {
	pageNum := rowNum / table.NumRows
	var page *Page
	page = &table.Pages[pageNum]
	if page == nil {
		// allocate memory to page if it is nil
		page = new(Page)
	}

	rowOffset := rowNum % table.NumRows

	// TODO get the rowSize
	rowSize := uint(128)
	byteOffset := rowOffset * rowSize
	//return page + byteOffset
	fmt.Println(byteOffset)
}

func SerializeRows(table Table) (result []Row) {

	for _, v := range table.Pages {
		result = append(result, v.Rows...)
	}

	return
}
func CreateTable(tableName string, numRows uint) Table {
	table := Table{
		TableName: tableName,
		NumRows:   numRows,
		Pages:     []Page{},
	}
	return table
}

func InsertRow(table *Table, rowData Row) {
	pageNum := len(table.Pages) - 1
	rowNum := len(table.Pages[pageNum].Rows) - 1
	table.Pages[pageNum].Rows[rowNum] = rowData
}

func SelectAllFromTable(table Table) {
	fmt.Printf("\n\nTable: %+v\n\n", table)
}

