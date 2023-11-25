package encoding

import (
	"encoding/json"
	"github.com/Yandex-Practicum/final-project-encoding-go/models"
	"gopkg.in/yaml.v3"
	"os"
)

// JSONData тип для перекодирования из JSON в YAML
type JSONData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// YAMLData тип для перекодирования из YAML в JSON
type YAMLData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// MyEncoder интерфейс для структур YAMLData и JSONData
type MyEncoder interface {
	Encoding() error
}

// Encoding перекодирует файл из JSON в YAML
func (j *JSONData) Encoding() error {
	var jsonFile []byte
	var yamlData []byte
	var yamlFile *os.File
	var err error

	if jsonFile, err = os.ReadFile(j.FileInput); err != nil {
		return err
	}

	if err = json.Unmarshal(jsonFile, &j.DockerCompose); err != nil {
		return err
	}

	if yamlData, err = yaml.Marshal(&j.DockerCompose); err != nil {
		return err
	}

	if yamlFile, err = os.Create(j.FileOutput); err != nil {
		return err
	}

	defer yamlFile.Close()

	if _, err = yamlFile.Write(yamlData); err != nil {
		return err
	}

	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	var yamlFile []byte
	var jsonData []byte
	var jsonFile *os.File
	var err error

	if yamlFile, err = os.ReadFile(y.FileInput); err != nil {
		return err
	}

	if err = yaml.Unmarshal(yamlFile, &y.DockerCompose); err != nil {
		return err
	}

	if jsonData, err = json.Marshal(&y.DockerCompose); err != nil {
		return err
	}

	if jsonFile, err = os.Create(y.FileOutput); err != nil {
		return err
	}

	defer jsonFile.Close()

	if _, err = jsonFile.Write(jsonData); err != nil {
		return err
	}

	return nil
}
