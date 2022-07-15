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
	params = strings.Split(input, " ")

	return
}

// input string is the string after "select" statement
func validate_select_operation(input string) (params []string, err error) {

	if len(input) == 0 {
		err = errors.New("parameters are empty")
		return
	}
	params = strings.Split(input, " ")

	return
}

// input string is the string after "delete" statement
func validate_delete_operation(input string) (params []string, err error) {

	if len(input) == 0 {
		err = errors.New("parameters are empty")
		return
	}
	params = strings.Split(input, " ")

	return
}

// input string is the string after "update" statement
func validate_update_operation(input string) (params []string, err error) {

	if len(input) == 0 {
		err = errors.New("parameters are empty")
		return
	}
	params = strings.Split(input, " ")

	return
}

// input string is the string after "update" statement
func validate_create_operation(input string) (params []string, err error) {

	if len(input) == 0 {
		err = errors.New("parameters are empty")
		return
	}
	params = strings.Split(input, " ")

	return
}

// input string is the string after "use" statement
func validate_use_operation(input string) (params []string, err error) {

	if len(input) == 0 {
		err = errors.New("parameters are empty")
		return
	}
	// input[1:] eliminates the space
	params = strings.Split(input[1:], " ")

	if len(params) > 1 {
		errMsg := fmt.Sprintf("syntax error at the end of : use %v", input)
		err = errors.New(errMsg)
		loghelper.LogError(errMsg)
		return
	}

	if err = dbutils.CheckDatabase(params[0]); err != nil {
		loghelper.LogError(err.Error())
		return
	}

	return
}

func validateStatement(input string, statementType int) (params []string, err error) {
	// figure out a way to find out whether  the database is selected or not
	if dbutils.DatabaseNotSelected() && statementType != constants.STATEMENT_USE {
		err = errors.New("No Database Selected")
		return
	}

	switch statementType {
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
		params, err = validate_create_operation(input)
		break
	case constants.STATEMENT_USE:
		params, err = validate_use_operation(input)
		break
	}
	return

}
