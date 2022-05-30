package main

import (
	"bufio"
	"os"

	"github.com/bihari123/building-a-database-in-golang/utils"
)

func main(){
  reader:=bufio.NewReader(os.Stdin)

  for{
   utils.PrintPrompt()
   utils.ReadInput(reader)
  }                                                                                            
}
