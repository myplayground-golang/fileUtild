package fileUtild

import (
	"os"
	"path/filepath"
)

func ListFiles(targetPath string) []string {
	files := make([]string, 0)

	var walkFunc filepath.WalkFunc = func(currentPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() { // file
			files = append(files, currentPath)
		}
		if info.IsDir() {
			// skip
		}

		return nil
	}

	err := filepath.Walk(targetPath, walkFunc)
	if err != nil {
		panic(err)
	}

	return files
}


func ListDirectories(targetPath string) []string {
	directories := make([]string, 0)

	var walkFunc filepath.WalkFunc = func(currentPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() { // file
			// skip
		}

		if info.IsDir() {
			directories = append(directories, currentPath)
		}

		return nil
	}

	err := filepath.Walk(targetPath, walkFunc)
	if err != nil {
		panic(err)
	}

	return directories
}
