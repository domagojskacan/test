package delete

import (
	"os"
	"path/filepath"
)

func DelJson(username string, filename string) error {
	return delFile(username, filename, ".json")
}

func DelYaml(username string, filename string) error {
	return delFile(username, filename, ".yaml")

}

func DelCsv(username string, filename string) error {
	return delFile(username, filename, ".csv")

}

func delFile(username string, filename string, tip string) error {
	path := filepath.Join("../../test", username, filename, tip)
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil

}
