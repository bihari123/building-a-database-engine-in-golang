package execute

import (
	"bufio"
	"strings"

	"github.com/bihari123/building-a-database-in-golang/constants"
)

func PrepareStatement(input string, reader bufio.Reader, statement *uint) int {

	if strings.Compare("insert", input[:6]) == 0 {
		*statement = constants.STATEMENT_INSERT
		return constants.PREPARE_SUCCESS
	} else if strings.Compare("select", input[:6]) == 0 {
		*statement = constants.STATEMENT_SELECT
		return constants.PREPARE_SUCCESS
	}

	return constants.PREPARE_UNRECOGNIZED_STATEMENT
}
