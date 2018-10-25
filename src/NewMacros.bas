Attribute VB_Name = "NewMacros"
Sub Style()
    Call ChangeMargin
    Call ChangeLineSpace
    Call ChangeLetterSize
    Call ChangeFontStyle
End Sub

Sub ChangeFontStyle()
    ActiveDocument.Range.Font.Name = "HG���ȏ���"
End Sub

Sub ChangePaperSizeToA3()
    With ActiveDocument.PageSetup
        .PaperSize = wdPaperA3
    End With
End Sub

Sub ChangeMargin()
    With ActiveDocument.PageSetup
        .LeftMargin = 40
        .RightMargin = 40
        .BottomMargin = 40
        .TopMargin = 40
    End With
End Sub

Sub ChangeColor()
    ActiveDocument.Range.Font.ColorIndex = wdBlack
End Sub

Sub ChangeLineSpace()
    With ActiveDocument.Paragraphs
        .LineSpacingRule = wdLineSpaceExactly
        .LineSpacing = 12
    End With
End Sub

Sub ChangeLetterSize()
    ActiveDocument.Range.Font.Size = 10.5
End Sub

Sub InputWord()
    Dim word(127) As String
    word(0) = "いきます"
    Call UnderlineOnSpecialWord(word)
End Sub

Sub UnderlineOnSpecialWord(str() As String)
'
' UnderlineOnSpecialWord Macro
    Dim myRange As Range
    Dim j As Long
    
    For j = LBound(str) To UBound(str)
       'myRange（オブジェクト変数）を設定
       Set myRange = ActiveDocument.Range(0, 0)
    
       '一括置換を実行（「検索と置換」ダイアログボックスの設定）
       With myRange.Find
         .Text = str(j)  '検索する文字列
         .Replacement.Text = ""    '置換後の文字列（空欄でOK）
         .Replacement.Font.Underline = wdUnderlineSingle  '一重下線
         .Forward = True
         .Wrap = wdFindStop
         .Format = True              '書式の設定をオン
         .MatchCase = False          '大文字と小文字の区別する
         .MatchWholeWord = False     '完全に一致する単語だけを検索する
         .MatchAllWordForms = False  '英単語の異なる活用形を検索する
         .MatchSoundsLike = False    'あいまい検索（英）
         .MatchFuzzy = False         'あいまい検索（日）
         .MatchWildcards = False     'ワイルドカードを使用する
         .MatchByte = True           '半角と全角を区別する
         .Execute Replace:=wdReplaceAll
       End With
    
       'myRangeを解放
       Set myRange = Nothing
    Next j
'
End Sub

Sub doctodocx()
    fpath = "C:\users\kaichi\Desktop\�e�X�g�쐬�\�t�g\"
    Set wd = CreateObject("Word.Application")
    wd.Visible = True
    FName = Dir(fpath & "*.doc")
    
    Do While (FName <> "")
        Set doc = wd.Documents.Open(fpath & FName)

        doc.SaveAs2 fpath & Mid(FName, 1, Len(FName) - 3) & "docm", 12
        doc.Close
        FName = Dir()
    Loop
    wd.Quit
End Sub
