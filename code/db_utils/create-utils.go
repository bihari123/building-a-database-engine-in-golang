package dbutils

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/bihari123/building-a-database-in-golang/utils/loghelper"
)

func CreateDB(dbName string) (err error) {

	if err = CheckDatabase(dbName, true); err != nil {
		loghelper.LogError(err.Error())
		return
	}

	home, _ := os.UserHomeDir()
	pathToDBEngineFiles := filepath.Join(home, "VectorDB")
	newDB := filepath.Join(pathToDBEngineFiles, dbName)

	if err = os.MkdirAll(newDB, os.ModePerm); err != nil {
		loghelper.LogError(err.Error())
		return
	}

	db := Database{
		Name: dbName,
	}

	Databases[dbName] = db

	loghelper.LogInfo(fmt.Sprintf("Database %v created", dbName))
	return
}
func DropDB(dbName string) (err error) {

	if err = CheckDatabase(dbName, false); err != nil {
		loghelper.LogError(err.Error())
		return
	}

	home, _ := os.UserHomeDir()
	pathToDBEngineFiles := filepath.Join(home, "VectorDB")
	newDB := filepath.Join(pathToDBEngineFiles, dbName)

	var confirmation string

	fmt.Printf("Are you sure to delete the database %v (y/n): ", dbName)
	fmt.Scan(&confirmation)

	if confirmation == "n\n" || confirmation == "N\n" {
		errMsg := "Aborting operation...."
		loghelper.LogError(errMsg)
		err = errors.New(errMsg)
		return
	}

	if err = os.RemoveAll(newDB); err != nil {
		loghelper.LogError(err.Error())
		return
	}
  
  // Refresh database 

	loghelper.LogInfo(fmt.Sprintf("Database %v deleted", dbName))
	return
}
