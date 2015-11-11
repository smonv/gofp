package main

import (
	"log"
	"os"
	"path/filepath"
)

type Result struct {
	Dir  int
	File int
}

func (r *Result) checkPermission(cPath string) error {
	src, err := os.Stat(cPath)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	fullPath, _ := filepath.Abs(cPath)
	if src.IsDir() {
		log.Println("Checking: " + cPath)
		if err = filepath.Walk(fullPath, r.visit); err != nil {
			return err
		}
	} else {
		err := r.checkFile(fullPath, src)

		if err != nil {
			log.Println(err)
		}
	}
	return nil
}

func (r *Result) checkDirectory(p string, dir os.FileInfo) error {
	if dir.Mode().String() != "drwxr-xr-x" {
		err := os.Chmod(p, 0755)
		if err == nil {
			log.Println("Fixed:", dir.Name())
			r.Dir++
		} else {
			return err
		}
	}
	return nil
}

func (r *Result) checkFile(cPath string, file os.FileInfo) error {
	if file.Mode().String() != "-rw-r--r--" {
		err := os.Chmod(cPath, 0644)
		if err == nil {
			log.Println("Fixed: ", file.Name())
			r.File++
		} else {
			return err
		}
	}
	return nil
}

func (r *Result) visit(cPath string, fileInfo os.FileInfo, err error) error {
	if fileInfo.IsDir() {
		err = r.checkDirectory(cPath, fileInfo)

		if err != nil {
			log.Println(err)
		}
	} else {
		err = r.checkFile(cPath, fileInfo)

		if err != nil {
			log.Println(err)
		}
	}
	return nil
}
