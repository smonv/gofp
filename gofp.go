package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	defer os.Exit(1)

	flag.Parse()
	src_dir := strings.TrimSpace(flag.Arg(0))

	if src_dir == "" {
		log.Println("Please input path to check!")
	} else {
		src, err := os.Stat(src_dir)
		if err != nil {
			log.Println(err)
		}

		if src.IsDir() {
			log.Println("Checking: " + src_dir)
			CheckPermission(src_dir)
		} else {
			log.Println("Source path not is directory")
		}
	}
}

func CheckPermission(dirPath string) {
	fullPath, err := filepath.Abs(dirPath)

	if err != nil {
		panic(err)
	}

	err = filepath.Walk(fullPath, walkFunc)

	if err != nil {
		panic(err)
	}
}

func walkFunc(path string, file os.FileInfo, err error) error {
	filename := file.Name()

	if file.IsDir() {
		if file.Mode().String() != "drwxrwxr-x" {
			err = os.Chmod(path, 0775)
			displayLog(&filename, err)
		}
	} else {
		if file.Mode().String() != "-rw-rw-r--" {
			err = os.Chmod(path, 0664)
			displayLog(&filename, err)
		}
	}
	return nil
}

func displayLog(filename *string, err error) {
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Fixed: ", *filename)
	}
}
