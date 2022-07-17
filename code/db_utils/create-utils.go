package dbutils

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/bihari123/building-a-database-in-golang/constants"
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

func CreateTable(pathToDB string, param []string) (err error) {

	// table = param[1]
	if err = CheckTable(pathToDB, param[1]); err != nil {
		loghelper.LogError(err.Error())
		return
	}

	pathToTable := filepath.Join(pathToDB, param[1]+".csv")

	var tableFile *os.File

	if tableFile, err = os.OpenFile(pathToTable, os.O_CREATE, os.ModePerm); err != nil {
		loghelper.LogError(err.Error())
		return
	}

	// param[2] = "(",  param[len(param)-1]=")"

	table := Table{
		TableName: param[1],
		Rows: []Row{
			{
				Id: 1,
				Fields: []Field{
					{
						Name: param[3],
						Type: constants.INT,
					},
					{
						Name: param[5],
						Type: constants.INT,
					},
				},
			},
		},
	}

	if err != nil {
		loghelper.LogError(err.Error())
		return
	}

	records := [][]string{
		{table.Rows[0].Fields[0].Name, table.Rows[0].Fields[1].Name},
	}

	w := csv.NewWriter(tableFile)

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatal("error writing record to file", err)
		}
	}

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
