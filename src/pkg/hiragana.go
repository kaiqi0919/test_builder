package hiragana

func Kanjiconv(in string) string {
	var out string
	for _, r := range []rune(in) {
		out = out + string(Kanjiruneconv(r))
	}
	return out
}

func Kanjiruneconv(r rune) rune {
	if Ishiragana(r) {
		return r
	}
	return '*'
}

func Ishiragana(r rune) bool {
	Lo := 0x3041
	Hi := 0x3093
	if (Lo <= int(r)) && (int(r) <= Hi) {
		return true
	}
	return false
}


