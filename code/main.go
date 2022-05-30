package main

import (
	"bufio"
	"os"

	"github.com/bihari123/building-a-database-in-golang/read"
	"github.com/bihari123/building-a-database-in-golang/print"

)

func main(){
  reader:=bufio.NewReader(os.Stdin)

  for{
   print.PrintPrompt()
   read.ReadInput(reader)
  }                                                                                            
}
