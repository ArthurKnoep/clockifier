package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

type (
	Clockify struct {
		ApiKey      string `json:"api_key"`
		WorkspaceId string `json:"workspace_id"`
	}
	Toggl struct {
		ApiKey      string `json:"api_key"`
		WorkspaceId string `json:"workspace_id"`
	}
	File struct {
		Clockify       Clockify          `json:"clockify"`
		Toggl          Toggl             `json:"toggl"`
		ProjectMapping map[string]string `json:"project_mapping"`
	}
)

var (
	NoConfigPresent = errors.New("no config present")
)

func ensureDirectory(fileName string) error {
	dirName := filepath.Dir(fileName)
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		if err := os.MkdirAll(dirName, os.ModePerm); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	return nil
}

func ensureFile(fileName string) error {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		if f, err := os.Create(fileName); err != nil {
			return err
		} else {
			f.Close()
		}
	} else if err != nil {
		return err
	}
	return nil
}

func ensure(path string) error {
	if err := ensureDirectory(path); err != nil {
		return err
	}
	if err := ensureFile(path); err != nil {
		return err
	}
	return nil
}

func readFile(path string) ([]byte, error) {
	if err := ensure(path); err != nil {
		return nil, err
	}
	if file, err := ioutil.ReadFile(path); err != nil {
		return nil, err
	} else {
		return file, nil
	}
}

func LoadConfig(path string) (*File, error) {
	fileContent, err := readFile(path)
	if err != nil {
		return nil, err
	}
	if len(fileContent) == 0 {
		return nil, NoConfigPresent
	}
	var cfg File
	if err := json.Unmarshal(fileContent, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func SaveConfig(path string, cfg *File) error {
	if err := ensure(path); err != nil {
		return err
	}
	f, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	if err := json.NewEncoder(f).Encode(cfg); err != nil {
		return err
	}
	return nil
}
