package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v3"
)

type Config struct {
	DataPath string
	Linux    bool
}

const (
	DAILYNOTE_ENV = "DAILYNOTE"
	DEFAULT_PERM  = 0644
)

func LoadConfig() (Config, bool) {
	var cfg Config
	cfg.Linux = (runtime.GOOS != "windows")

	path := os.Getenv(DAILYNOTE_ENV)
	if path == "" {
		homeDir, err := os.UserHomeDir()

		if err != nil {
			log.Fatal(err)
		}

		if cfg.Linux {
			cfg.DataPath = fmt.Sprintf("%s/.dailynote/", homeDir)
		} else {
			cfg.DataPath = fmt.Sprintf("%s\\.dailynote\\", homeDir)
		}

	} else {
		cfg.DataPath = path
	}

	_, err := os.Stat(cfg.DataPath)

	return cfg, !os.IsNotExist(err)
}

func CreateConfigFolder(cfg Config) {
	err := os.MkdirAll(cfg.DataPath, DEFAULT_PERM)
	if err != nil {
		log.Fatal(err)
	}
	createTagsFile(cfg)
}

func createTagsFile(cfg Config) {
	tagspath := filepath.Join(cfg.DataPath, "tags.yml")

	_, err := os.Create(tagspath)

	if err != nil {
		log.Fatal(err)
	}

}

func AddTag(cfg Config, tag, id string) {
	tagspath := filepath.Join(cfg.DataPath, "tags.yml")
	file, err := ioutil.ReadFile(tagspath)

	if err != nil {
		log.Fatal(err)
	}

	readData := make(map[string][]string)
	yamlReadErr := yaml.Unmarshal(file, &readData)

	if yamlReadErr != nil {
		log.Fatal(yamlReadErr)
	}

	currents := readData[tag]

	currents = append(currents, id)
	readData[tag] = currents

	writeData, yamlWriteErr := yaml.Marshal(readData)

	if yamlWriteErr != nil {
		log.Fatal(yamlWriteErr)
	}

	fileWriteErr := ioutil.WriteFile(tagspath, writeData, DEFAULT_PERM)

	if fileWriteErr != nil {
		log.Fatal(fileWriteErr)
	}

	log.Printf("Add id='%s' to tag='%s'.\n", id, tag)
}
