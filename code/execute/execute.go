package execute

import (
	"fmt"
	"strings"

	"github.com/bihari123/building-a-database-in-golang/constants"
)

func PrepareStatement(input string,  statementType *int) int {

	if strings.Compare("insert", input[:6]) == 0 {
		*statementType = constants.STATEMENT_INSERT
		return constants.PREPARE_SUCCESS
	} else if strings.Compare("select", input[:6]) == 0 {
		*statementType = constants.STATEMENT_SELECT
		return constants.PREPARE_SUCCESS
	}

	return constants.PREPARE_UNRECOGNIZED_STATEMENT
}

func ExecuteStatement(statementType int) {
	switch statementType {
	case constants.STATEMENT_INSERT:
		fmt.Println("this is where we will do an insert")
		break
	case constants.STATEMENT_SELECT:
		fmt.Println("this is where we will do a select")
		break
	}
}
