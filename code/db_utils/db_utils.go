package dbutils

type Row struct {
	Id       uint
	UserName string
	Email    string
}


type Statement struct{
  Type int 
  RowToInsert Row 
}

type Table struct{
	NumRows uint 
	Pages []interface{}
}


func RowSlot(table *Table,rowNum uint){
	pageNum:=rowNum/table.NumRows 
	var page *interface{}
	page=&table.Pages[pageNum]
	if page==nil{
	// allocate memory to page if it is nil
	}

	//rowOffset:=rowNum%table.NumRows

	// TODO get the rowSize
	//rowSize:=uint(128)
	//byteOffset:= rowOffset*rowSize
//return page + byteOffset

}
