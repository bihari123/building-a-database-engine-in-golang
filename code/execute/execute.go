package execute

import (
	"fmt"
	"strings"

	"github.com/bihari123/building-a-database-in-golang/constants"
)

func PrepareStatement(input string, statementType *int) int {


	if strings.Compare("insert", input[:6]) == 0 {
		*statementType = constants.STATEMENT_INSERT
	} else if strings.Compare("select", input[:6]) == 0 {
		*statementType = constants.STATEMENT_SELECT
	} else if strings.Compare("delete", input[:6]) == 0 {
		*statementType = constants.STATEMENT_DELETE
	} else if strings.Compare("update", input[:6]) == 0 {
		*statementType = constants.STATEMENT_UPDATE
	} else {
		return constants.PREPARE_UNRECOGNIZED_STATEMENT
	}
  
  params,err:=perform_operation( input[6:],*statementType)
  if err!=nil{
  	return constants.PREPARE_SYNTAX_ERROR                                                                                                     
  }
  fmt.Println("PARAMETERS DEFINED: ",params)
  return constants.PREPARE_SUCCESS 
}

func ExecuteStatement(statementType int) {
	switch statementType {
	case constants.STATEMENT_INSERT:
		fmt.Println("this is where we will do an insert")
		break
	case constants.STATEMENT_SELECT:
		fmt.Println("this is where we will do a select")
		break
	case constants.STATEMENT_DELETE:
		fmt.Println("this is where we will do a delete")
		break
	case constants.STATEMENT_UPDATE:
		fmt.Println("this is where we will do an update")
		break

	}
}
