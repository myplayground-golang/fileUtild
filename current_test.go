package fileUtild

import (
	"gotest.tools/v3/assert"
	"testing"
)

func TestProjectRootFolder(t *testing.T) {
	t.Run("testGetCurrentProjectRootFolderSuccess", testGetCurrentProjectRootFolderSuccess)
	t.Run("testGetCurrentProjectRootFolderFail", testGetCurrentProjectRootFolderFail)
}

func testGetCurrentProjectRootFolderSuccess(t *testing.T) {
	rootFolder, err := GetCurrentProjectRootFolder("fileUtild")
	if err != nil {
		panic(err)
	}
	assert.Assert(t, func() bool {
		return rootFolder == "/home/I309008/projects/polarisx_go/project_golang_module/fileUtild"
	}())
}

func testGetCurrentProjectRootFolderFail(t *testing.T) {
	rootFolder, err := GetCurrentProjectRootFolder("fileUtil")
	assert.Assert(t, func() bool{
		return rootFolder == ""
	}())
	assert.ErrorContains(t, err,"Unable to identify the project root folder")
}
