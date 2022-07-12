package fileio

func SaveFile(username string, filename string, data string) error {
	err := saveJson(username, filename, data)
	// ...
	return err
}

func saveJson(username string, filename string, data string) error {
	return nil
}

func saveCsv(username string, filename string, data string) error {
	return nil
}

func saveYaml(username string, filename string, data string) error {
	return nil
}

func ReadYaml(username string, filename string) (string, error) {
	return "", nil
}

func ReadJson(username string, filename string) (string, error) {
	return "", nil
}

func ReadCsv(username string, filename string) (string, error) {
	return "", nil
}
