package main

import (
    "fmt"
    "io/ioutil"
	"path/filepath"
	"os/exec"
)

func main() {
	slice := dirwalk("../ユーザー")
	fmt.Println(slice)

	for _, file := range slice {
		err := exec.Command("nkf", "-w", "--overwrite", file).Run()
		if err != nil {
			fmt.Println("command exec error.")
			fmt.Println(err)
        }
        fmt.Printf("%s は変換されました\n", file)
	}
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
