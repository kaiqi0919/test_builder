テスト作成ソフト
====

概要
小テストや期末テストを自動生成する教師向けソフトです。

## 説明
問題を登録して、テスト範囲を指定して実行すると、テストの問題用紙と解答用紙が自動生成されます。
外国人に日本語を教える日本語学校の教師向けに、現在開発中です。

## 使用環境
* OS・ソフトウェア: 
- Windows7以上, Mac(未確認)
- **Microsoft Word (2007以降)**

## インストール
* Githubからインストールする場合
1. **release** にZipファイルがあるので、最新版をダウンロードします。
2. 保存したい場所にフォルダを展開します。

## HowToUse
現在は「アチーブメントテスト作成」と「期末テスト作成」と「問題登録」の機能があります。
* 準備
0. **現在開いているWord文書を、すべて閉じてください。**

* アチーブメントテスト作成
1. テスト作成ソフトのフォルダから、 **アチーブメントテスト作成.exe** を実行します。初回のみ、セキュリティブロックされる可能性があるので、「詳細」から「実行」をします。
2. テストの内容確認画面が出るので、確認したら **1** を入力して **Enter** を押します。テストの内容を変えたい場合は **0** を入力して **Enter** を押して、プログラムを一度終了します。
3. Wordファイル`docm_generator.docm`が開かれます。（30秒～1分ほどかかる場合があります。）
4. メッセージボックスに **〇〇が正常に生成されました。** と表示されたら成功です。 **OK** を押すとプログラムが終了します。（10秒ほどかかる場合があります。）
5. 元フォルダに、2つのファイル **問題用紙x.docm、解答用紙x.docm** があります。

* 期末テスト作成
アチーブメントテスト作成とほぼ同じなので割愛します。

* 問題登録
テスト作成ソフトのフォルダから、 **ユーザー** フォルダの中にある **データベース** フォルダの中に、問題データを入れてください。

* 問題データについて
現在、以下のフォーマットで作っています。


> 漢字読み書きの問題文および解答のセットを登録する方法について

> <データセットの作り方>
> 〇作成環境
> テキストファイル（メモ帳）が操作できれば特に指定はありません。
> フォルダ管理も適当で大丈夫（誰がいつどのファイルを作ったかわかるといいかも）

> 〇ファイル名
> ファイルの形式は「***.txt」で保存
> ファイルタイトルは、出題箇所がわかるようなタイトルでお願いします。（「N5_1章_読み1.txt」など）

> 〇内容
> 1行目にレベルを入力して改行
> 2行目に問題科目を入力して改行
> 3行目に大問項目を入力して改行
> 4行目に章番号を入力して改行
> 5行目に難易度を入力して改行
> 6行目に問題文
> 7行目に下線箇所
> 8行目に解答
> 9～11行目に誤答
> 12行目に空行
> 13行目以降にルビセット
> ルビセットの後に空行。その空行以降は読み取りません。備考などあれば自由に書いてください。

> 〇レベル
> N1/N2/N3/N4/N5から1つ選択

> "〇問題科目
文字/語彙/文法/読解/聴解から1つ選択

〇大問項目（文字）
漢字読み/表記から1つ選択

〇章番号
「1章」、「2章」など章単位で記入

〇難易度
難/易から1つ選択

〇ルビについて
ルビをつける場合は、章番号から1行空けて、ルビをつける漢字とルビを以下のように入力してください。
例．
漢字　　かんじ
読　　　よ


<例>
N5
文字
漢字読み
8章
難
今日は　お父さんと　電話　します。
お父さん
おとうさん
おふさん
おとさん
おじいさん

今日　きょう
電話　でんわ"

## Licence

## Author
[鈴木海地](https://github.com/kaiqi0919)
