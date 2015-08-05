package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
)

func main() {
	flag.Parse()
	src_dir := flag.Arg(0)
	if src_dir == "" {
		src_dir = "/mnt"
	}

	log.Println("Checking: " + src_dir)

	src, err := os.Stat(src_dir)
	if err != nil {
		panic(err)
	}

	if src.IsDir() {
		CheckPermission(src_dir)
	} else {
		log.Println("Source path not is directory")
	}
	os.Exit(1)
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
