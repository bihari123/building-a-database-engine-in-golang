package main

import (
	"bufio"
	"os"

	"github.com/bihari123/building-a-database-in-golang/config"
	"github.com/bihari123/building-a-database-in-golang/print"
	"github.com/bihari123/building-a-database-in-golang/read"
	"github.com/bihari123/building-a-database-in-golang/utils/loghelper"
)

func main(){
  reader:=bufio.NewReader(os.Stdin)
  
  appConfig:=config.Appconfig{
  	Prod: false,
  }
  loghelper.SetPrint(appConfig.Prod)
  loghelper.LogInfo("print is working")


	for{
   print.PrintPrompt()
   read.ReadInput(reader)
  }                                                                                            
}
