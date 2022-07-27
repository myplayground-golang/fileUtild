package fileUtild

import (
	"github.com/pkg/errors"
	"os"
	"path/filepath"
	"strings"
)

func GetCurrentWorkingFolder() string {
	currentWorkingDirectory, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return currentWorkingDirectory
}

func GetCurrentProjectRootFolder(projectName string) (string, error) {
	currentWorkingDir := GetCurrentWorkingFolder()
	parts := strings.Split(currentWorkingDir, "/")

	found := false
	newParts := make([]string, 0)
	for _, part := range parts {
		if part == projectName {
			found = true
			newParts = append(newParts, part)
			break
		} else {
			newParts = append(newParts, part)
		}
	}

	if !found {
		return "", errors.New("Unable to identify the project root folder")
	}

	return "/" + filepath.Join(newParts...), nil
}