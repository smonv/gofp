package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Result struct {
	Dir  int
	File int
}

func main() {
	defer os.Exit(1)

	flag.Parse()
	dirPath := strings.TrimSpace(flag.Arg(0))
	result := &Result{0, 0}
	result = CheckPermission(dirPath, result)
	if result != nil {
		fmt.Println("Total Fixed Dir: ", result.Dir)
		fmt.Println("Total Fixed File: ", result.File)
	}
}

func CheckPermission(dirPath string, r *Result) (result *Result) {
	if dirPath == "" {
		log.Println("Please input path to check!")
	} else {
		src, err := os.Stat(dirPath)
		if err != nil {
			log.Println(err)
			return nil
		}
		if src.IsDir() {
			log.Println("Checking: " + dirPath)
			fullPath, _ := filepath.Abs(dirPath)
			if err = filepath.Walk(fullPath, r.Visit); err != nil {
				log.Println(err)
			} else {
				return r
			}
		} else {
			log.Println("Source path not is directory")
		}
	}
	return nil
}

func (r *Result) Visit(path string, file os.FileInfo, err error) error {
	if file.IsDir() {
		if file.Mode().String() != "drwxr-xr-x" {
			if err = os.Chmod(path, 0755); err == nil {
				log.Println("Fixed: ", file.Name())
				r.Dir++
			} else {
				log.Println(err)
			}
		}
	} else {
		if file.Mode().String() != "-rw-r--r--" {
			if err = os.Chmod(path, 0644); err == nil {
				log.Println("Fixed: ", file.Name())
				r.File++
			} else {
				log.Println(err)
			}
		}
	}
	return nil
}
