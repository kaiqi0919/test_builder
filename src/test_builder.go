package main

import (
	"io"
    "os"
	"fmt"
	"strconv"
	"strings"
	"math/rand"
	"time"
	"io/ioutil"
	"errors"
	"encoding/json"
	"./pkg"
)

// TestSet hoge
type TestSet struct {
	ATitle		Title		`json:"タイトル"`
	TestType	string		`json:"テスト種類"`
	Problems	[]Problem	`json:"問"`
}

// Title hoge
type Title struct {
	MainTitle	string	`json:"表題"`
	SubTitle	string	`json:"サブタイトル"`
}

// Problem hoge
type Problem struct {
	Number	int		`json:"大問番号"`
	Format	string	`json:"大問フォーマット"`
	Total	int		`json:"問題数"`
	Ranges	RangeP	`json:"出題範囲"`
}

// RangeP hoge
type RangeP struct {
	Level	string	`json:"レベル"`
	Kamoku	string	`json:"科目"`
	Daimon	string	`json:"大問"`
	Section	string	`json:"章番号"`
	Ease	string	`json:"難易度"`
}

// CSV_one_line hoge
type CSV_one_line struct {
	Level	string
	Kamoku	string
	Daimon	string
	Section	string
	Ease	string
	Problem	string
	ULine	string
	Correct	string
	Wrong_1	string
	Wrong_2	string
	Wrong_3 string
	Kana	string
	Rubis	[]RubiSet
}

type RubiSet struct {
	Kanji	string
	Rubi	string
}

func NewLine(linebuff string) (*CSV_one_line, error) {
	words := strings.Split(linebuff, ",")
	if len(words)<11 {
		return nil, errors.New("Wrong Style.")
	}

	line := &CSV_one_line{Level: words[0], Kamoku: words[1], Daimon: words[2], Section: words[3], Ease: words[4], 
		Problem: words[5], ULine: words[6], Correct: words[7], Wrong_1: words[8], Wrong_2: words[9], Wrong_3: words[10]}
	for i:=11; i<len(words); i++ {
		rubiset, err := NewRubiSet(words[i])
		if err == nil {
			line.Rubis = append(line.Rubis, *rubiset)
		}
	}
	return line, nil
}

func NewRubiSet(str string) (*RubiSet, error) {
	var halfspace, fullspace []string
	halfspace = strings.Split(str, " ")
	for _, halfwords := range halfspace {
		fullwords := strings.Split(halfwords, "　")
		fullspace = append(fullspace, fullwords...)
	}
	if len(fullspace) != 2 {
		return nil, errors.New("Wrong Rubi Style.")
	}
	return &RubiSet{Kanji: fullspace[0], Rubi: fullspace[1]}, nil
}

func (line *CSV_one_line) SprintWithChoice() string {
	var str string
	str = str + fmt.Sprintf(". %s\r\n", line.Problem)

	index := randompick(4, 4)
	var choice [4]string = [4]string{line.Correct, line.Wrong_1, line.Wrong_2, line.Wrong_3}
	str = str + fmt.Sprintf("  ① %s ② %s ③ %s ④ %s\r\n\r\n", PaddingSizeSprint(choice[index[0]], -17), 
		PaddingSizeSprint(choice[index[1]], -17), PaddingSizeSprint(choice[index[2]], -17), PaddingSizeSprint(choice[index[3]], -17))
	return str
}

func (line *CSV_one_line) SprintWithoutChoice() string {
	var str string
	str = str + fmt.Sprintf(". %s\r\n\r\n", line.Problem)
	fmt.Println(hiragana.Kanjiconv(line.Problem))
	return str
}

// 指定文字数だけ「半角」でスペースを取った文字列を返す
func PaddingSizeSprint(str string, size int) string {
	var padding_str, space string
	length := len(str) - len([]rune(str))
	// 4byteの漢字を使ったときにそろわない（修正の必要なし）
	if size < -length {
		for i:=0; i< -size-length ; i++ {
			space = space + " "
		}
		padding_str = fmt.Sprintf("%s%s", str, space)
	}else if size > length {
		for i:=0; i<size-length; i++ {
			space = space + " "
		}
		padding_str = fmt.Sprintf("%s%s", space, str)
	}else {
		return str
	}
	return padding_str
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

// go run test_builder.go "../ユーザー/アチーブメントテストセット.json"

func main() {
	path := os.Args[1]
	raw, err := ioutil.ReadFile(path)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

	var ts TestSet
	json.Unmarshal(raw, &ts)

	csvname := "../etc/problems.csv"
	allbuff, err := ioutil.ReadFile(csvname)
	if err != nil {
		os.Exit(0)
	}
	linebuffs := strings.Split(string(allbuff), "\n")

	str, uline, returncount, rubis, kana_yomi, kana_kaki := strcreate(ts, linebuffs)

	docname := "../問題用紙"
	for i:=1; i<100; i++ {
		if Exists(docname + strconv.Itoa(i) + ".doc") {
			continue
		}
		docname = docname + strconv.Itoa(i) + ".doc"
		break
	}
	CreateDoc(docname, str)

	str = "" //初期化
	for i:=0; i<returncount; i++ {
		var fullsplit []string
		halfsplit := strings.Split(uline[i], " ")
		for _, word := range halfsplit {
			fullsplit = append(fullsplit, strings.Split(word, "　")...)
		}
		for _, word := range fullsplit {
			str = str + word + ","
		}
		str = str + "\r\n"
	}
	csvname = "../etc/uline.csv"
	filecreate(str, csvname)

	str = "" //初期化
	for _, rubi := range rubis {
		str = str + rubi.Rubi
		str = str + "\r\n"
	}
	csvname = "../etc/rubi.csv"
	filecreate(str, csvname)

	str = ""
	for _, a := range kana_yomi {
		str = str + a
		str = str + "\r\n"
	}
	csvname = "../etc/figure1.csv"
	filecreate(str, csvname)

	str = ""
	for _, a := range kana_kaki {
		str = str + a
		str = str + "\r\n"
	}
	csvname = "../etc/figure2.csv"
	filecreate(str, csvname)

	t := time.Now()
	const layout = "2006-01-02"
	str = ts.ATitle.MainTitle + "\r\n" + t.Format(layout)
	headername := "../etc/header.txt"
	filecreate(str, headername)
}

func strcreate(ts TestSet, linebuffs []string) (string, []string, int, []RubiSet, []string, []string) {
	var str string
	var uline []string = make([]string, 1024)
	var returncount int
	var rubis []RubiSet = make([]RubiSet, 10)
	var kana_yomi []string
	var kana_kaki []string

	str = str + ts.ATitle.SubTitle + ts.Problems[0].Ranges.Level + " " + ts.Problems[0].Ranges.Section + "\r\n\r\n"
	if ts.TestType == "記述式" {
		str = str + "クラス　　　　　なまえ　　　　　　　　　　　　　　　あ" + "\r\n\r\n"
		returncount = 3
	} else {
		returncount = 1
	}

	for _, p := range ts.Problems {
		r := p.Ranges
		if r.Daimon == "漢字読み" {
			kana_yomi = make([]string, p.Total)
		}else if r.Daimon == "表記" {
			kana_kaki = make([]string, p.Total)
		}


		var lines []CSV_one_line
		for i, linebuff := range linebuffs {
			line, err := NewLine(linebuff)
			if (err == nil) && (line.Level == r.Level || r.Level == "") &&
			(line.Kamoku == r.Kamoku || r.Kamoku == "") && (line.Daimon == r.Daimon || r.Daimon == "") && 
			(line.Section == r.Section || r.Section == "") && (line.Ease == r.Ease || r.Ease == "") {
				lines = append(lines, *line)
				fmt.Println(i)
			}else {
				continue
			}
		}

		amount := len(lines)
		fmt.Println(amount)
		fmt.Println(p.Total)

		rand.Seed(time.Now().UnixNano())

		numset := randompick(amount, p.Total)
		fmt.Println(numset)

		str = str + p.Format + "\r\n"
		returncount++
		for j, i := range numset {
			str = str + strconv.Itoa(j+1)
			if ts.TestType == "選択式" {
				str = str + lines[i].SprintWithChoice()
				returncount++
				uline[returncount] = lines[i].ULine
				returncount = returncount + 2
	
			} else if ts.TestType == "記述式" {
				str = str + lines[i].SprintWithoutChoice()
				for _, rubi := range lines[i].Rubis {
					fmt.Println(rubi.Kanji, rubi.Rubi)
				}
				returncount++
				uline[returncount] = lines[i].ULine
				returncount = returncount + 1
				if r.Daimon == "漢字読み" {
					kana_yomi[j] = outhiragana(lines[i].ULine)
				}else if r.Daimon == "表記" {
					kana_kaki[j] = outhiragana(lines[i].Correct)
				}
			}
		}
		str = str + "\r\n"
		returncount = returncount + 1
	}
	return str, uline, returncount, rubis, kana_yomi, kana_kaki
}

// 解答用紙出力用の文字列データ
// 合計で半角スペース40個分以内
func outhiragana(uline string) string {
	var str string
	var runeuline []rune = []rune(uline)
	hiraganacount := HiraganaCount(uline)
	kanji := len(runeuline) - hiraganacount
	aspace := (20 - hiraganacount * 2) / kanji
	
	for _, r := range runeuline {
		if hiragana.Ishiragana(r) {
			str = str + string(r)
		} else {
			for i:=0; i<aspace; i++ {
				str = str + " "
			}
		}
	}
	return str
}

func HiraganaCount(str string) int {
	count := 0
	for _, r := range []rune(str) {
		if hiragana.Ishiragana(r) {
			count++
		}
	}
	return count
}

func filecreate(str string, filename string) {
	ioutil.WriteFile(filename, []byte(str), 0666)
}

func filecopy(srcname string, dstname string) {
	src, err := os.Open(srcname)
	if err != nil {
		panic(err)
	}
	defer src.Close()

	dst, err := os.Create(dstname)
	if err != nil {
		panic(err)
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		panic(err)
	}
}

func randompick(amount int, total int) []int {
	if amount < total {
		fmt.Printf("Number of problem should not be more than %d\n", amount)
		fmt.Println("Abort this pick")
		return []int{}
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
