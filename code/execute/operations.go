package execute

import (
	"errors"
	"strings"

	"github.com/bihari123/building-a-database-in-golang/constants"
)

// input string is the string after "insert" statement
func insert_operation(input string) (params []string, err error) {

	if len(input) == 0 {
		err = errors.New("parameters are empty")
		return
	}
	params = strings.Split(input, " ")

	return
}

// input string is the string after "select" statement
func select_operation(input string) (params []string, err error) {

	if len(input) == 0 {
		err = errors.New("parameters are empty")
		return
	}
	params = strings.Split(input, " ")

	return
}

// input string is the string after "delete" statement
func delete_operation(input string) (params []string, err error) {

	if len(input) == 0 {
		err = errors.New("parameters are empty")
		return
	}
	params = strings.Split(input, " ")

	return
}

// input string is the string after "update" statement
func update_operation(input string) (params []string, err error) {

	if len(input) == 0 {
		err = errors.New("parameters are empty")
		return
	}
	params = strings.Split(input, " ")

	return
}

func perform_operation(input string, statementType int) (params []string, err error) {

	switch statementType {
	case constants.STATEMENT_SELECT:
		params, err = select_operation(input)
		break
	case constants.STATEMENT_INSERT:
		params, err = insert_operation(input)
		break
	case constants.STATEMENT_DELETE:
		params, err = delete_operation(input)
	case constants.STATEMENT_UPDATE:
		params, err = update_operation(input)
	}
return

}

