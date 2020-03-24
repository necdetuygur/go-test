package main

import (
	"os"
	"io/ioutil"
)

func main() {
	RootFolder := "./htdocs"
	files, _ := ioutil.ReadDir(RootFolder)
	for _, f := range files {
		SubFolder := f.Name()
		Files, _ := ioutil.ReadDir(RootFolder + "/" + SubFolder)
		for _, f2 := range Files {
			FileName := f2.Name()
			FilePath := RootFolder + "/" + SubFolder + "/" + FileName
			TargetPath := RootFolder + "/" + SubFolder + "_" + FileName
			os.Rename(FilePath, TargetPath)
		}
	}
}

go build renamer.go && ./renamer*
