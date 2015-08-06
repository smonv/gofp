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
	CheckPermission(src_dir)
}

func CheckPermission(dirPath string) {
	if dirPath == "" {
		log.Println("Please input path to check!")
	} else {
		src, err := os.Stat(dirPath)
		if err != nil {
			log.Println(err)
		}
		if src.IsDir() {
			log.Println("Checking: " + dirPath)
			fullPath, _ := filepath.Abs(dirPath)
			if err = filepath.Walk(fullPath, walkFunc); err != nil {
				log.Println(err)
			}
		} else {
			log.Println("Source path not is directory")
		}
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
