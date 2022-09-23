package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func isExist(_path string) bool {
	if _, err := os.Stat(_path); err != nil {
		return false
	} else {
		return true
	}
}

func createDesktopDir(dirName string) string {
	currentUser, err := user.Current()
	check(err)

	desktop := currentUser.HomeDir + "\\Desktop\\"
	newDir := desktop + dirName

	if isExist(newDir) {
		fmt.Println(newDir, "\t\t\t isAlreadyExist: true")
	} else {
		if err := os.MkdirAll(newDir, os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}

	return newDir
}

func createFile(fileName string, dirName string) string {
	dir := createDesktopDir(dirName)
	fileNamePath := dir + "\\" + fileName

	if isExist(fileNamePath) {
		fmt.Println(fileNamePath, "\t\t\t isAlreadyExist: true")
	} else {
		file, err := os.Create(fileNamePath)
		check(err)

		if err := file.Close(); err != nil {
			log.Fatal(err)
		}

		_filepath, _err := filepath.Abs(file.Name())
		check(_err)

		return _filepath
	}

	return fileNamePath
}

func writeData(filePath string, data []byte) {
	fileStream, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	check(err)

	if _, err := fileStream.Write(data); err != nil {
		log.Fatal(err)
	}

	if err := fileStream.Close(); err != nil {
		log.Fatal(err)
	}
}

func readData(filePath string) []byte {
	fileStream, err := os.ReadFile(filePath)
	check(err)

	return fileStream
}

func getFileInfo(filePath string) {
	if isExist(filePath) {
		fileInfo, err := os.Stat(filePath)
		check(err)

		fmt.Println("File Name:", fileInfo.Name())
		fmt.Println("Size:", fileInfo.Size(), "bytes")
		fmt.Println("Permissions:", fileInfo.Mode())
		fmt.Println("Last Modified:", fileInfo.ModTime())
		fmt.Println("Is Directory: ", fileInfo.IsDir())
	} else {
		fmt.Println("file dose not exist")
	}

}

func getAll(dir string) {
	files, err := os.ReadDir(dir)
	check(err)

	if len(files) == 0 {
		fmt.Println("empty dir")
	} else {
		for _, file := range files {
			fmt.Println(">", file.Name(), " \t ", "isDir:", file.IsDir())
		}
	}
}

func main() {
	filePath := createFile("go.txt", "GoResourceFolder")

	writeData(filePath, []byte("Hello, World!"))
	readData(filePath)
	getFileInfo(filePath)

	currentDir, err := filepath.Abs(filepath.Dir(filePath))
	check(err)
	getAll(currentDir)
}
