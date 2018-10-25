package main

import (
	"strings"
    "fmt"
    "io/ioutil"
	"path/filepath"
	"os/exec"
)

func main() {
	slice := dirwalk("../ユーザー")
	fmt.Println(slice)
	count := 0
	filestyle := "UTF-8"

	for _, file := range slice {
		out, err := exec.Command("nkf", "-g", file).Output()
		if err != nil {
			fmt.Println("command exec error.")
			fmt.Println(err)
        }
		fmt.Printf("%s のファイルは %s です。\n", file, out)
		if !strings.HasPrefix(string(out), filestyle) {
			fmt.Println(len(out))
			count++
		}
	}
	fmt.Printf("%d 個のファイルが %s ではありません。\n", count, string(filestyle))
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
