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

func NewFileLoader(dataDir string) (*FileLoader, error) {
	fileLoader := FileLoader{}

	// Check if the directory exists. If not, raise an error
	if _, err := os.Stat(dataDir); err == nil {
		fileLoader.dataFilesLocation = dataDir
	} else if os.IsNotExist(err) {
		return nil, err
	}

	return &fileLoader, nil
}

// LoadDataFromFile takes a single filename (located in DataLoader.dataFilesLocation), and parses it, returning a
// map representation of the data contained in the file
func (fl *FileLoader) LoadDataFromFile(fileName string) (map[string]interface{}, error) {
	filePath := filepath.FromSlash(fl.dataFilesLocation + "/" + fileName)
	jsonFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}

	json.Unmarshal([]byte(byteValue), &result)

	return result, nil
}

// LoadAllFromFiles will walk the data directory provided to the FileLoader, and load into dictionaries any data it
// finds, and return these as a map, whose keys are the filenames, and the values the data loaded from those files.
func (fl *FileLoader) LoadAllFromFiles() (map[string]map[string]interface{}, error) {
	data := make(map[string]map[string]interface{})

	var files []string

	err := filepath.Walk(fl.dataFilesLocation, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})

	if err != nil {
		return nil, err
	}

	baseDir := fl.dataFilesLocation + "/"

	for _, file := range files {

		if file == fl.dataFilesLocation {
			continue
		}

		pathElements := strings.Split(file, baseDir)
		loadedData, err := fl.LoadDataFromFile(strings.Join(pathElements, ""))

		if err != nil {
			return nil, err
		}

		fileName := strings.TrimSuffix(file, path.Ext(file))
		data[fileName] = loadedData
	}

	return data, nil
}
