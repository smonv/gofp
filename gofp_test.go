package main

import (
	"os"
	"path"
	"testing"
)

var testDir = "/tmp/gofp"

func TestCheckDir(t *testing.T) {
	os.Mkdir(testDir, 0777)
	defer os.RemoveAll(testDir)

	dirPath := path.Join(testDir, "dir")
	os.Mkdir(dirPath, 0775)

	r := &Result{0, 0}
	dirInfo, _ := os.Stat(dirPath)
	err := r.checkDirectory(dirPath, dirInfo)
	if err != nil {
		t.Error(err)
	}

	dirInfo, _ = os.Stat(dirPath)
	if dirInfo.Mode().String() != "drwxr-xr-x" {
		t.Error(dirInfo.Name(), "wrong permission: ", dirInfo.Mode())
	}

}

func TestCheckFile(t *testing.T) {
	os.Mkdir(testDir, 0777)
	defer os.RemoveAll(testDir)

	filePath := path.Join(testDir, "file")
	os.Create(filePath)
	os.Chmod(filePath, 0666)

	r := &Result{0, 0}

	fileInfo, _ := os.Stat(filePath)
	err := r.checkFile(filePath, fileInfo)
	if err != nil {
		t.Error(err)
	}

	fileInfo, _ = os.Stat(filePath)
	if fileInfo.Mode().String() != "-rw-r--r--" {
		t.Error(fileInfo.Name(), "wrong permission: ", fileInfo.Mode())
	}
}

func TestCheckPermission(t *testing.T) {
	os.Mkdir(testDir, 0777)
	defer os.RemoveAll(testDir)

	tmp := [][]string{{"dir1", "file1"}, {"dir2", "file2"}, {"dir3", "file3"}}
	for _, v := range tmp {
		dirPath := path.Join(testDir, v[0])
		os.Mkdir(dirPath, 0777)
		os.Chmod(dirPath, 0777)
		file, _ := os.Create(path.Join(dirPath, v[1]))
		file.Chmod(0666)
	}

	r := &Result{0, 0}

	err := r.checkPermission(testDir)

	if err != nil {
		t.Error(err)
	} else {
		for _, v := range tmp {
			dirPath := path.Join(testDir, v[0])
			dir, _ := os.Stat(dirPath)
			if dir.Mode().String() != "drwxr-xr-x" {
				t.Error(dir.Name(), "wrong permission: ", dir.Mode())
			}

			file, _ := os.Stat(path.Join(dirPath, v[1]))
			if file.Mode().String() != "-rw-r--r--" {
				t.Error(file.Name(), "wrong permission: ", file.Mode())
			}
		}
	}
}
