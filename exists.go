package fileUtild

import "os"

// IsDirExist verify path folder exists
func IsDirExist(path string) bool {
	fileinfo, err := os.Stat(path)
	if fileinfo != nil && !fileinfo.IsDir() {
		panic(path + " is not a valid path")
	}
	if err != nil {
		return os.IsExist(err)
	} else {
		return true
	}
}

// MakeDirExist make sure this path folder will be created if not existing
func MakeDirExist(path string) error {
	if IsDirExist(path) {
		return nil
	} else {
		return os.MkdirAll(path, 0755)
	}
}

// IsFileExist verify file exists
func IsFileExist(file string) bool {
	_, err := os.Stat(file)
	if err != nil {
		return os.IsExist(err)
	} else {
		return true
	}
}

// MakeFileExist make sure this file will be created if not existing
func MakeFileExist(file string) error {
	if IsFileExist(file) {
		return nil
	} else {
		f, err := os.Create(file)
		if err != nil {
			return err
		}
		f.Close()
		return nil
	}
}
