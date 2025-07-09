package utils

import (
	"os"
	)



// Check if a file or directory exists
func Exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// Create a folder if it doesn't exist 
func CreateFolder(Path string) error {
	err := os.MkdirAll(Path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}


