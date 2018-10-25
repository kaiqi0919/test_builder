package main

import (
	"fmt"
	"io/ioutil"
	"bufio"
	"os"
	"path/filepath"
)

func main() {
	var foldername string
	foldername = "../ユーザー/データベース"
	database := dirwalk(foldername)
	var buff string
	for _, filename := range database {
		fmt.Println(filename)
		fp, err := os.Open(filename)
		if err != nil {
			fmt.Println("error\n")
			continue
		}

		scanner := bufio.NewScanner(fp)
		for scanner.Scan() {
			buff = buff + scanner.Text() + ","
		}
		buff = buff + "\r\n"
	}
	csvcreate(buff, "../etc/problems.csv") 
}

func csvcreate(str string, filename string) {
	ioutil.WriteFile(filename, []byte(str), 0666)
}

func dirwalk(dir string) []string {
    files, err := ioutil.ReadDir(dir)
    if err != nil {
        panic(err)
    }

    var paths []string
    for _, file := range files {
        if file.IsDir() {
            paths = append(paths, dirwalk(filepath.Join(dir, file.Name()))...)
            continue
        }
        paths = append(paths, filepath.Join(dir, file.Name()))
    }

    return paths
}
