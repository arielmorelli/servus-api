package filehandler

import (
	"encoding/json"
	"io/ioutil"
	"os"

	models "github.com/arielmorelli/servus-api/models"
	route "github.com/arielmorelli/servus-api/route"
)

func getFileData(filename string) ([]byte, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	byteValue, _ := ioutil.ReadAll(file)

	return byteValue, nil
}

// LoadRoutesFromFile loads routers from a JSON file
func LoadRoutesFromFile(filename string) error {
	var registerSchemaArray []models.RegisterSchema

	fileContent, err := getFileData(filename)
	if err != nil {
		return err
	}

	json.Unmarshal(fileContent, &registerSchemaArray)

	for _, schema := range registerSchemaArray {
		route.RegisterRoute(schema)
	}

	return nil
}
