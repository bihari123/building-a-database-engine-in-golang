package golangbuffers

import (
	"bufio"
	"log"
	"os"
)

func WriteIntoFile(filePath string) (err error) {
	
	
// OpenFile is the generalized open call; most users will use Open
// or Create instead. It opens the named file with specified flag
// (O_RDONLY etc.). If the file does not exist, and the O_CREATE flag
// is passed, it is created with mode perm (before umask). If successful,
// methods on the returned File can be used for I/O.
// If there is an error, it will be of type *PathError.

	filePtr, err := os.OpenFile(filePath, os.O_WRONLY, 0666)
	if err != nil {
		log.Println("Error opening file ", err)
		return err
	}

	// create a buffer writer using gile pointer

// NewWriter returns a new Writer whose buffer has the default size.
// If the argument io.Writer is already a Writer with large enough buffer size,
// it returns the underlying Writer.
	bufferedWriter := bufio.NewWriter(filePtr)

	if _, err = bufferedWriter.WriteString("hello world"); err != nil {
		log.Println("error writing to file: ", err)
	}

	bufferedWriter.Flush()
	filePtr.Close()
	return
}

func ReadFile(filePath string)(err error){

	filePtr, err := os.Open("Demo.txt")
	if err != nil {
		log.Println(err)
		return 
	}

	bufferedReader := bufio.NewReader(filePtr)

	data := make([]byte, 12)
	len, err := bufferedReader.Read(data)
	if err != nil {
		log.Println(err)
		return 
	}
	log.Printf("Data Len: %d\n", len)
	log.Printf("Data    : %s\n", data)

	filePtr.Close()

	return 
}
