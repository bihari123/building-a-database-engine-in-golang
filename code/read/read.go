package read

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/bihari123/building-a-database-in-golang/constants"
	"github.com/bihari123/building-a-database-in-golang/execute"
)

// in our case, reader = os.Stdin
func ReadInput(reader *bufio.Reader) {

	// ReadString reads until the first occurrence of delim in the input,
	// returning a string containing the data up to and including the delimiter.
	// If ReadString encounters an error before finding a delimiter,
	// it returns the data read before the error and the error itself (often io.EOF).
	// ReadString returns err != nil if and only if the returned data does not end in
	// delim.
	// For simple uses, a Scanner may be more convenient.
	//func (b *Reader) ReadString(delim byte) (string, error) {

	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)
	var msgCode int
	if input[0] == '.' {
		msgCode = do_meta_command(input)
	} else {
		msgCode=do_sql_command(input)
	}
	fmt.Println("The statement executed, messageCode: ", msgCode)
}

func do_meta_command(input string) int {

	// Compare returns an integer comparing two strings lexicographically.
	// The result will be 0 if a == b, -1 if a < b, and +1 if a > b.
	//
	// Compare is included only for symmetry with package bytes.
	// It is usually clearer and always faster to use the built-in
	// string comparison operators ==, <, >, and so on.
	// func Compare(a, b string) int {

	if strings.Compare(".exit", input) == 0 {
		fmt.Println("Exiting....bye bye")
		os.Exit(constants.EXIT_SUCCESS)
	} // fmt.Println("unrecognized command : ", input)
	fmt.Println("the meta command was not recognized")
	return constants.META_COMMAND_UNRECOGNIZED_COMMAND

}

func do_sql_command(input string) (msgCode int) {
	var statement int
	prepMsgCode,params:=execute.PrepareStatement(input, &statement)
	switch prepMsgCode  {
	case constants.PREPARE_SUCCESS:
		msgCode = constants.PREPARE_SUCCESS
		execute.ExecuteStatement(statement,params)
		break
	case constants.PREPARE_UNRECOGNIZED_STATEMENT:
		msgCode = constants.PREPARE_UNRECOGNIZED_STATEMENT
		fmt.Printf("unrecognized keyword at the start of %s. \n", input)
	}
	return
}
