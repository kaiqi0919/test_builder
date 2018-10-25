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
	ATitle		Title		`json:"タイトル"`
	Problems	[]Problem	`json:"問"`
}

// Title hoge
type Title struct {
	MainTitle	string	`json:"表題"`
}

// Problem hoge
type Problem struct {
	Number	int			`json:"大問番号"`
	Format	string		`json:"大問フォーマット"`
	Total	int			`json:"問題数"`
	Ranges	RangeP	`json:"出題範囲"`
}

// RangeP hoge
type RangeP struct {
	Level	string	`json:"レベル"`
	Kamoku	string	`json:"科目"`
	Daimon	string	`json:"大問"`
	Section	string	`json:"章番号"`
}

func main() {
	fmt.Println("loading format...")
	fmt.Println("以下の内容でテストを作成します。")
	fmt.Println("")
	path := "./ユーザー/期末テストセット.json"
	raw, err := ioutil.ReadFile(path)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

	var ts TestSet
	json.Unmarshal(raw, &ts)

	fmt.Printf("タイトル: %s\n", ts.ATitle.MainTitle)

    for _, p := range ts.Problems {
		fmt.Printf("大問%d\n", p.Number)
		fmt.Printf("  小問数: %d\n", p.Total)
		fmt.Printf("  出題範囲\n")
		fmt.Printf("    レベル: %s\n", p.Ranges.Level)
		fmt.Printf("    科目: %s\n", p.Ranges.Kamoku)
		fmt.Printf("    大問: %s\n", p.Ranges.Daimon)
		fmt.Printf("    章番号: %s\n", p.Ranges.Section)
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
			os.Exit(1)
		}

		fmt.Println("converting doc to docm...")
		err3 := exec.Command("openfile.exe").Run()
		if err3 != nil {
			os.Exit(1)
		}

		fmt.Println("done")
		fmt.Println("テストは正常に作成されました。")
		fmt.Println("プログラムを終了します。")
		time.Sleep(1*time.Second)
	}
}

