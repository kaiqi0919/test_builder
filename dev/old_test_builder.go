package main

import (
	"flag"
    "os"
	"fmt"
	"strconv"
	"strings"
	"math/rand"
	"time"
	"io/ioutil"
	"errors"
)

type CSV_one_line struct {
	Level	string
	Kamoku	string
	Daimon	string
	Section	string
	Problem	string
	ULine	string
	Correct	string
	Wrong_1	string
	Wrong_2	string
	Wrong_3 string
}

func NewLine(linebuff string) (*CSV_one_line, error) {
	words := strings.Split(linebuff, ",")
	if len(words)<10 {
		return nil, errors.New("Wrong Style.")
	}

	line := &CSV_one_line{Level: words[0], Kamoku: words[1], Daimon: words[2], Section: words[3], 
		Problem: words[4], ULine: words[5], Correct: words[6], Wrong_1: words[7], Wrong_2: words[8], Wrong_3: words[9]}
	return line, nil
}

func (line *CSV_one_line) Sprint() string {
	var str string
	str = str + fmt.Sprintf("問題. %s\n", line.Problem)

	index := randompick(4, 4)
	var choice [4]string = [4]string{line.Correct, line.Wrong_1, line.Wrong_2, line.Wrong_3}
	str = str + fmt.Sprintf("A.%s B.%s C.%s D.%s\n\n", choice[index[0]], choice[index[1]], choice[index[2]], choice[index[3]])
	return str
}

func CreateDoc(filename string, str string) {
    file, err := os.Create(filename)
    if err != nil {
        // Openエラー処理
    }
    defer file.Close()

    file.Write(([]byte)(str))
}


// go run testmaker_sample.go -L N5 -K 文字 -T 10
// -D 漢字読み -S 1章 

func main() {
	var level, kamoku, daimon, section, totalnumber string

	flag.StringVar(&level, "L", "", "level")
	flag.StringVar(&kamoku, "K", "", "kamoku")
	flag.StringVar(&daimon, "D", "", "daimon")
	flag.StringVar(&section, "S", "", "section")
	flag.StringVar(&totalnumber, "T", "", "totalnumber")
	flag.Parse()

	total, _ := strconv.Atoi(totalnumber)

	csvname := "../etc/problems.csv"

	var lines []CSV_one_line
	allbuff, err := ioutil.ReadFile(csvname)
	if err != nil {
		os.Exit(0)
	}

	linebuffs := strings.Split(string(allbuff), "\n")
	for _, linebuff := range linebuffs {
		line, err := NewLine(linebuff)
		if err == nil {
			if (line.Level == level || level == "") && (line.Kamoku == kamoku || kamoku == "") && 
			(line.Daimon == daimon || daimon == "") && (line.Section == section || section == ""){
				lines = append(lines, *line)
			}
		}
	}

	amount := len(lines)
	fmt.Println(amount)
	fmt.Println(total)

	rand.Seed(time.Now().UnixNano())

	numset := randompick(amount, total)
	fmt.Println(numset)
	var str string

	for _, i := range numset {
		str = str + lines[i].Sprint()
	}

	docname := "../Sample"
	for i:=1; i<100; i++ {
		if Exists(docname + strconv.Itoa(i) + ".doc") {
			continue
		}
		docname = docname + strconv.Itoa(i) + ".doc"
		break
	}
	CreateDoc(docname, str)
}

func randompick(amount int, total int) []int {
	if amount < total {
		fmt.Printf("Number of problem should not be more than %d\n", amount)
		fmt.Println("Program Abort")
		os.Exit(0)
	}

	var slice, tank []int
	var i int
	for i=0; i<amount; i++ {
		tank = append(tank, i)
	}

	for i=0; i<total; i++ {
		j := rand.Intn(amount-i)
		slice = append(slice, tank[j])
		tank = append(tank[:j], tank[j+1:]...)
	}
	return slice
}

func shuffle(data []string) {
    n := len(data)
    for i := n - 1; i >= 0; i-- {
        j := rand.Intn(i + 1)
        data[i], data[j] = data[j], data[i]
    }
}

func Exists(filename string) bool {
    _, err := os.Stat(filename)
    return err == nil
}
