package execute

import (
	"errors"
	"strings"

	"github.com/bihari123/building-a-database-in-golang/constants"
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

func databaseNotSelected()bool{
	return false //for now
}


func validateStatement(input string, statementType int) (params []string, err error) {
	// figure out a way to find out whether  the database is selected or not
	if databaseNotSelected() {
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
	}
	return

}
