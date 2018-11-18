package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// This will be responsible for taking data files, parsing them, and storing the parsed data in an in memory mapping.
// Data files will contain information such as an item identifier, name, description, and systems that need to be
// attached to the entity, color, stats, etc.

type FileLoader struct {
	dataFilesLocation string
}

// LoadDataFromFile takes a single filename (located in DataLoader.dataFilesLocation), and parses it, returning a
// map representation of the data contained in the file
func (fl *FileLoader) LoadDataFromFile(fileName string) map[string]interface{} {
	filePath := filepath.FromSlash(fl.dataFilesLocation + "/" + fileName)
	jsonFile, err := os.Open(filePath)

	if err != nil {
		// TODO: Proper error handling here would be ideal
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}

	json.Unmarshal([]byte(byteValue), &result)

	return result
}

// LoadAllFromFiles will walk the data directory provided to the FileLoader, and load into dictionaries any data it
// finds, and return these as a map, whose keys are the filenames, and the values the data loaded from those files.
func (fl * FileLoader) LoadAllFromFiles() map[string]map[string]interface{} {
	data := make(map[string]map[string]interface{})

	var files []string

	err := filepath.Walk(fl.dataFilesLocation, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})

	if err != nil {
		// TODO: proper error handling
		panic(err)
	}

	for _, file := range files {
		loadedData := fl.LoadDataFromFile(file)
		fileName := strings.TrimSuffix(file, path.Ext(file))
		data[fileName] = loadedData
	}

	return data
}
