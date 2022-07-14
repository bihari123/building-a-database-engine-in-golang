package execute

import (
	"fmt"
	"strings"

	"github.com/bihari123/building-a-database-in-golang/constants"
	dbutils "github.com/bihari123/building-a-database-in-golang/db_utils"
)

func PrepareStatement(input string, statementType *int) (msgCode int, params []string) {

	if strings.Compare("insert", input[:6]) == 0 {
		*statementType = constants.STATEMENT_INSERT
	} else if strings.Compare("select", input[:6]) == 0 {
		*statementType = constants.STATEMENT_SELECT
	} else if strings.Compare("delete", input[:6]) == 0 {
		*statementType = constants.STATEMENT_DELETE
	} else if strings.Compare("update", input[:6]) == 0 {
		*statementType = constants.STATEMENT_UPDATE
	} else if strings.Compare("create", input[:6]) == 0 {
		*statementType = constants.STATEMENT_CREATE
	} else {
		msgCode = constants.PREPARE_UNRECOGNIZED_STATEMENT
		return msgCode, []string{}
	}

	params, err := perform_operation(input[6:], *statementType)
	if err != nil {
		msgCode = constants.PREPARE_SYNTAX_ERROR
		return msgCode, []string{}
	}
	fmt.Println("PARAMETERS DEFINED: ", params)
	msgCode = constants.PREPARE_SUCCESS
	return msgCode, params
}

func ExecuteStatement(statementType int,params []string) {
	switch statementType {
	case constants.STATEMENT_INSERT:
		fmt.Println("this is where we will do an insert")
		dbutils.InsertRow()
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
	case constants.STATEMENT_CREATE:
		fmt.Println("this is where we will do a create")
		break

	}
}
