package execute

import (
	"fmt"
	"strings"

	"github.com/bihari123/building-a-database-in-golang/constants"
	dbutils "github.com/bihari123/building-a-database-in-golang/db_utils"
	"github.com/bihari123/building-a-database-in-golang/utils/loghelper"
)

func PrepareStatement(input string, statementType *int) (msgCode int, params []string) {
	var inputParams string
	if strings.Compare("insert", input[:6]) == 0 {
		*statementType = constants.STATEMENT_INSERT
		inputParams = input[6:]
	} else if strings.Compare("select", input[:6]) == 0 {
		*statementType = constants.STATEMENT_SELECT
		inputParams = input[6:]
	} else if strings.Compare("delete", input[:6]) == 0 {
		*statementType = constants.STATEMENT_DELETE
		inputParams = input[6:]
	} else if strings.Compare("update", input[:6]) == 0 {
		*statementType = constants.STATEMENT_UPDATE
		inputParams = input[6:]
	} else if strings.Compare("create", input[:6]) == 0 {
		*statementType = constants.STATEMENT_CREATE
		inputParams = input[6:]
	} else if strings.Compare("use", input[:3]) == 0 {
		*statementType = constants.STATEMENT_USE
		inputParams = input[3:]
	}else if strings.Compare("drop",input[:4])==0{
    *statementType =constants.STATEMENT_DROP 
    inputParams=input[4:]
	}	else {
		msgCode = constants.PREPARE_UNRECOGNIZED_STATEMENT
		return msgCode, []string{}
	}
	// verify whether the statement is valid. Ex. if it contains the right number of params or not
	params, err := validateStatement(inputParams, statementType)
	if err != nil {
		loghelper.LogError(fmt.Sprintf("error validating the statement: %v\n\n", err))
		msgCode = constants.PREPARE_SYNTAX_ERROR
		return msgCode, []string{}
	}
	fmt.Println("PARAMETERS DEFINED: ", params)
	msgCode = constants.PREPARE_SUCCESS
	return msgCode, params
}

func ExecuteStatement(statementType int, params []string) (err error) {
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
	case constants.STATEMENT_CREATE:
		fmt.Println("this is where we will do a create")
		break
	case constants.STATEMENT_CREATE_DB:
		fmt.Println("this is where we will do a createDB")
		err = dbutils.CreateDB(params[1])
		break
	case constants.STATEMENT_USE:
		fmt.Println("this is where we will do a use")
		break
	case constants.STATEMENT_DROP_DB:
		err = dbutils.DropDB(params[1])
		break

	}
	return
}
