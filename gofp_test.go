package main

import (
	"os"
	"path"
	"testing"
)

func TestCheckPermission(t *testing.T) {
	test_dir := "/tmp/gofp"
	os.Mkdir(test_dir, 0777)
	// because umask=0002 so cannot create dir with 0777 permission
	os.Chmod(test_dir, 0777)
	defer os.RemoveAll(test_dir)

	tmp := [][]string{{"dir1", "file1"}, {"dir2", "file2"}, {"dir3", "file3"}}
	for _, v := range tmp {
		dir_path := path.Join(test_dir, v[0])
		os.Mkdir(dir_path, 0777)
		os.Chmod(dir_path, 0777)
		file, _ := os.Create(path.Join(dir_path, v[1]))
		file.Chmod(0666)
	}
	r := &Result{0, 0}
	CheckPermission(test_dir, r)

	for _, v := range tmp {
		dir_path := path.Join(test_dir, v[0])
		dir, _ := os.Stat(dir_path)
		if dir.Mode().String() != "drwxr-xr-x" {
			t.Error(dir.Name(), "wrong permission: ", dir.Mode())
		}

		file, _ := os.Stat(path.Join(dir_path, v[1]))
		if file.Mode().String() != "-rw-r--r--" {
			t.Error(file.Name(), "wrong permission: ", file.Mode())
		}
	}
}
