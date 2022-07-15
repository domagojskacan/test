package delete

import (
	"os"
	"path/filepath"
)

func DelJson(username string, filename string) error {
	path := filepath.Join("../../test", username, filename+".json")
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}

func DelYaml(username string, filename string) error {
	path := filepath.Join("../../test", username, filename+".yaml")
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}

func DelCsv(username string, filename string) error {
	path := filepath.Join("../../test", username, filename+".csv")
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}
