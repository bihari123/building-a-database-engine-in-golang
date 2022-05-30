package read

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/bihari123/building-a-database-in-golang/constants"
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


}

func do_meta_command(reader *bufio.Reader,input string)int{

	
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
	}	// fmt.Println("unrecognized command : ", input)
	  return constants.META_COMMAND_UNRECOGNIZED_COMMAND 
	

}
