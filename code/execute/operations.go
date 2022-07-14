package execute

import (
	"errors"
	"strings"

	"github.com/bihari123/building-a-database-in-golang/constants"
)

// input string is the string after "insert" statement
func verify_insert_operation(input string) (params []string, err error) {

	if len(input) == 0 {
		err = errors.New("parameters are empty")
		return
	}
	params = strings.Split(input, " ")

	return
}

// input string is the string after "select" statement
func verify_select_operation(input string) (params []string, err error) {

	if len(input) == 0 {
		err = errors.New("parameters are empty")
		return
	}
	params = strings.Split(input, " ")

	return
}

// input string is the string after "delete" statement
func verify_delete_operation(input string) (params []string, err error) {

	if len(input) == 0 {
		err = errors.New("parameters are empty")
		return
	}
	params = strings.Split(input, " ")

	return
}

// input string is the string after "update" statement
func verify_update_operation(input string) (params []string, err error) {

	if len(input) == 0 {
		err = errors.New("parameters are empty")
		return
	}
	params = strings.Split(input, " ")

	return
}

// input string is the string after "update" statement
func verify_create_operation(input string) (params []string, err error) {

	if len(input) == 0 {
		err = errors.New("parameters are empty")
		return
	}
	params = strings.Split(input, " ")

	return
}


func validateStatement(input string, statementType int) (params []string, err error) {

	switch statementType {
	case constants.STATEMENT_SELECT:
		params, err = verify_select_operation(input)
		break
	case constants.STATEMENT_INSERT:
		params, err = verify_insert_operation(input)
		break
	case constants.STATEMENT_DELETE:
		params, err = verify_delete_operation(input)
	case constants.STATEMENT_UPDATE:
		params, err = verify_update_operation(input)
	case constants.STATEMENT_CREATE:
		params, err = verify_create_operation(input)

	}
	return

}
