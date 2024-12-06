package utils

import "os"

func ReadData(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
