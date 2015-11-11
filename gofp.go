package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	defer os.Exit(1)

	cPath := flag.String("p", "", "check path")

	flag.Parse()

	if *cPath != "" {
		r := &Result{0, 0}
		err := r.checkPermission(*cPath)
		if err == nil {
			fmt.Println("Total Fixed Dir: ", r.Dir)
			fmt.Println("Total Fixed File: ", r.File)
		} else {
			log.Fatal(err)
		}
	} else {
		log.Panic("Please input path to check!")
	}
}
