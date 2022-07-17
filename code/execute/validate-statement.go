package execute

import (
	"errors"
	"fmt"
	"strings"

	"github.com/bihari123/building-a-database-in-golang/constants"
	dbutils "github.com/bihari123/building-a-database-in-golang/db_utils"
	"github.com/bihari123/building-a-database-in-golang/utils/loghelper"
)

// input string is the string after "insert" statement
func validate_insert_operation(input string) (params []string, err error) {

	if len(input) == 0 {
		err = errors.New("parameters are empty")
		return
	}
	params = strings.Split(strings.TrimSpace(input), " ")

	return
}

// input string is the string after "select" statement
func validate_select_operation(input string) (params []string, err error) {

	if len(input) == 0 {
		err = errors.New("parameters are empty")
		return
	}
	params = strings.Split(strings.TrimSpace(input), " ")

	return
}

// input string is the string after "delete" statement
func validate_delete_operation(input string) (params []string, err error) {

	if len(input) == 0 {
		err = errors.New("parameters are empty")
		return
	}
	params = strings.Split(strings.TrimSpace(input), " ")

	return
}

// input string is the string after "update" statement
func validate_update_operation(input string) (params []string, err error) {

	if len(input) == 0 {
		err = errors.New("parameters are empty")
		return
	}
	params = strings.Split(strings.TrimSpace(input), " ")

	return
}

// input string is the string after "update" statement
func validate_create_operation(input string, statementType *int) (params []string, err error) {

	var errMsg string

	if len(input) == 0 {
		err = errors.New("parameters are empty")
		return
	}
	params = strings.Split(strings.TrimSpace(input), " ")

	switch params[0] {
	//create database DBNAME
	case "database":
		if len(params) > 2 {
			errMsg = fmt.Sprintf("Error in the number of params: %v", params)
			loghelper.LogError(errMsg)
			err = errors.New(errMsg)
			return
		}
		*statementType = constants.STATEMENT_CREATE_DB
	
	case "table":
		if dbutils.DatabaseNotSelected() {
			errMsg = fmt.Sprint("Database not selected")
			loghelper.LogError(errMsg)
			err = errors.New(errMsg)
			return
		}

		// tableName := params[1]

		if params[2] != "(" || params[len(params)-1] != ")" {
			errMsg = fmt.Sprintf("Syntax error near: %v\n", input)
			loghelper.LogError(errMsg)
			err = errors.New(errMsg)
			return
		}

		if len(params) < 6 {
			errMsg = fmt.Sprintf("Error in the number of params: %v", params)
			loghelper.LogError(errMsg)
			err = errors.New(errMsg)
			return
		}

		*statementType = constants.STATEMENT_CREATE_TABLE
	
	default:
		errMsg = fmt.Sprintf("Unrecognized Keyword %s", params[0])
		loghelper.LogError(errMsg)
		err = errors.New(errMsg)
	}

	return
}

// input string is the string after "use" statement
func validate_use_operation(input string) (params []string, err error) {

	if len(input) == 0 {
		err = errors.New("parameters are empty")
		return
	}
	// input[1:] eliminates the space
	params = strings.Split(strings.TrimSpace(input), " ")

	if len(params) > 1 {
		errMsg := fmt.Sprintf("syntax error at the end of : use %v", input)
		err = errors.New(errMsg)
		loghelper.LogError(errMsg)
		return
	}

	if err = dbutils.CheckDatabase(params[0], false); err != nil {
		loghelper.LogError(err.Error())
		return
	}

	return
}

// input string is the string after "drop" statement
func validate_drop_operation(input string, statementType *int) (params []string, err error) {
	var errMsg string
	if len(input) == 0 {
		err = errors.New("parameters are empty")
		return
	}
	params = strings.Split(strings.TrimSpace(input), " ")

	switch params[0] {
	//create database DBNAME
	case "database":
		if len(params) > 2 {
			errMsg = fmt.Sprintf("Error in the number of params: %v", params)
			loghelper.LogError(errMsg)
			err = errors.New(errMsg)
		}
		*statementType = constants.STATEMENT_DROP_DB
	case "table":
		loghelper.LogInfo("Need to figure this out")

		if dbutils.DatabaseNotSelected() {
			errMsg = fmt.Sprint("Database not selected")
		}
		// if dbutils.TableNotPresentInDB(){
		// 	errMsg+=fmt.Sprintf("Table: %v not present in Database: %v",table,database)
		// }

		*statementType = constants.STATEMENT_DROP_TABLE
	}

	return
}

func validateStatement(input string, statementType *int) (params []string, err error) {
	// figure out a way to find out whether  the database is selected or not
	if dbutils.DatabaseNotSelected() && *statementType != constants.STATEMENT_USE && *statementType != constants.STATEMENT_CREATE && *statementType != constants.STATEMENT_DROP {
		err = errors.New("No Database Selected")
		return
	}

	switch *statementType {
	case constants.STATEMENT_SELECT:
		params, err = validate_select_operation(input)
		break
	case constants.STATEMENT_INSERT:
		params, err = validate_insert_operation(input)
		break
	case constants.STATEMENT_DELETE:
		params, err = validate_delete_operation(input)
		break
	case constants.STATEMENT_UPDATE:
		params, err = validate_update_operation(input)
		break
	case constants.STATEMENT_CREATE:
		params, err = validate_create_operation(input, statementType)
		break
	case constants.STATEMENT_USE:
		params, err = validate_use_operation(input)
		break
	case constants.STATEMENT_DROP:
		params, err = validate_drop_operation(input, statementType)
		break

	}
	return

}
