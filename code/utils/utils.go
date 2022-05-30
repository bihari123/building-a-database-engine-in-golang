package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
const(
  SUCCESS = 0
)
func PrintPrompt() {
	fmt.Print("db > ")
}

func ReadInput() {

	// ReadString reads until the first occurrence of delim in the input,
	// returning a string containing the data up to and including the delimiter.
	// If ReadString encounters an error before finding a delimiter,
	// it returns the data read before the error and the error itself (often io.EOF).
	// ReadString returns err != nil if and only if the returned data does not end in
	// delim.
	// For simple uses, a Scanner may be more convenient.
	//func (b *Reader) ReadString(delim byte) (string, error) {

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)

	if strings.Compare(".exit", input) == 0 {
		fmt.Println("Exiting....bye bye")
	  os.Exit(SUCCESS)
	}

}
