package main

import (
	"os"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"os/exec"
	"time"
)

// TestSet hoge
type TestSet struct {
	Date		string		`json:"実施日"`
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
	Number	int			`json:"大問番号"`
	Format	string		`json:"大問フォーマット"`
	Total	int			`json:"問題数"`
	Ranges	[]RangeP	`json:"出題範囲"`
}

// RangeP hoge
type RangeP struct {
	Level	string	`json:"レベル"`
	Kamoku	string	`json:"科目"`
	Daimon	string	`json:"大問"`
	Section	string	`json:"章番号"`
	Ease	string	`json:"難易度"`
}

func main() {
	fmt.Println("loading format...")
	fmt.Println("以下の内容でテストを作成します。")
	fmt.Println("")
	path := "./ユーザー/期末テストセット_3a.json"
	raw, err := ioutil.ReadFile(path)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

	var ts TestSet
	json.Unmarshal(raw, &ts)

	fmt.Printf("タイトル: %s\n", ts.ATitle.MainTitle)
	fmt.Printf("実施日: %s\n", ts.Date)
	
    for _, p := range ts.Problems {
		fmt.Printf("大問%d\n", p.Number)
		fmt.Printf("  小問数: %d\n", p.Total)
		for i, r := range p.Ranges {
			if len(p.Ranges) == 1 {
				fmt.Printf("  出題範囲\n")
			} else {
				fmt.Printf("  出題範囲 %d\n", i + 1)
			}
			fmt.Printf("    レベル: %s\n", r.Level)
			fmt.Printf("    科目: %s\n", r.Kamoku)
			fmt.Printf("    大問: %s\n", r.Daimon)
			fmt.Printf("    章番号: %s\n", r.Section)
			fmt.Printf("    難易度: %s\n\n", r.Ease)
		}
	}
	fmt.Println("以上\n")
	fmt.Println("1 + Enter を入力して続行してください。")
	fmt.Println("0 + Enter を入力するとプログラムを終了します。")
	var toggle int
	fmt.Scanf("%d", &toggle)
	if toggle == 1 {
//		for _, q := range ts.Problems{
//			
//		}
		os.Chdir("./bin")

		fmt.Println("creating CSV...")
		err1 := exec.Command("csvwriter.exe").Run()
		if err1 != nil {
			os.Exit(1)
		}
		
		fmt.Println("building doc...")
		pathfromsrc := "." + path
		err2 := exec.Command("test_builder.exe", pathfromsrc).Run()
		if err2 != nil {
			fmt.Println("error")
			fmt.Println(pathfromsrc)
			time.Sleep(1*time.Second)
			
			os.Exit(1)
		}

		fmt.Println("converting doc to docm...")
		err3 := exec.Command("openfile_comprehensive.exe").Run()
		if err3 != nil {
			os.Exit(1)
		}

		fmt.Println("done")
		fmt.Println("テストは正常に作成されました。")
		fmt.Println("プログラムを終了します。")
		time.Sleep(1*time.Second)
	}
}

