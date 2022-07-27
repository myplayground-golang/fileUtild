package main

import (
	"fmt"
	"github.com/myplayground-golang/fileUtild"
)

func main() {
	home, err := fileUtild.Home()
	if err != nil {
		panic(err)
	}
	fmt.Println(home)
	fmt.Println(fileUtild.GetCurrentWorkingFolder())
	fmt.Println(fileUtild.IsDirExist("/home/DOCKER"))

	size, err := fileUtild.DirSizeMB("/home/DOCKER/project_golang_2")
	if err != nil {
		panic(err)
	}
	fmt.Println(size)

	fmt.Println(fileUtild.RandomString(24))
	fmt.Println(fileUtild.RandomInt(10))
}
