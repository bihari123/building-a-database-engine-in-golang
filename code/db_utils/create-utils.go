package dbutils

import (
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

	if err = os.RemoveAll(newDB); err != nil {
		loghelper.LogError(err.Error())
		return
	}
	loghelper.LogInfo(fmt.Sprintf("Database %v deleted", dbName))
	return
}

