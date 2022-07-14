# Working

## 1. Read Execute Print Loop

we read the user input, decide what commands to run and then print the output.

Using an infinite loop, we prompt the user to enter the query command

```
for{
   print.PrintPrompt()
   read.ReadInput(reader)
  }
```

The prompt message looks like the following

```
func PrintPrompt() {
	fmt.Print("db > ")
}
```

The read function looks like the following: 


```
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
  // if the command starts with ".", then it is a meta command
  if input[0] == '.' {
		msgCode = do_meta_command(input)
	} else {
		msgCode=do_sql_command(input)
	}
	fmt.Println("The statement executed, messageCode: ", msgCode)
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

```

Here we have two types of commands:- 
 - Meta_Command: for operations like exit,user_login etc.
 - Sql_Command: for operations using sql query.


Meta_Commands are executed as follows:

```

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

```

The Sql_Commands are executed as follows:
```
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

```
Here, we first prepare the statement and then execute it upon getting the success_code PREPARE_SUCCESS.

PrepareStatement is responsible for verifying whether the statement that we are going to execute is valid or not. It sends the prepMsgCode and params. prepMsgCode can have two values: constants.PREPARE_SUCCESS or constants.PREPARE_UNRECOGNIZED_STATEMENT.

```
func PrepareStatement(input string, statementType *int) (msgCode int, params []string) {

	if strings.Compare("insert", input[:6]) == 0 {
		*statementType = constants.STATEMENT_INSERT
	} else if strings.Compare("select", input[:6]) == 0 {
		*statementType = constants.STATEMENT_SELECT
	} else if strings.Compare("delete", input[:6]) == 0 {
		*statementType = constants.STATEMENT_DELETE
	} else if strings.Compare("update", input[:6]) == 0 {
		*statementType = constants.STATEMENT_UPDATE
	} else if strings.Compare("create", input[:6]) == 0 {
		*statementType = constants.STATEMENT_CREATE
	} else {
		msgCode = constants.PREPARE_UNRECOGNIZED_STATEMENT
		return msgCode, []string{}
	}

	params, err := perform_operation(input[6:], *statementType)
	if err != nil {
		msgCode = constants.PREPARE_SYNTAX_ERROR
		return msgCode, []string{}
	}
	fmt.Println("PARAMETERS DEFINED: ", params)
	msgCode = constants.PREPARE_SUCCESS
	return msgCode, params
}

```
