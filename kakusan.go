package kakusan

import "unicode"
import "strings"

var (
	hankakuToZenkaku = map[string]string{
		"ｦ": "ヲ",
		"ｧ": "ァ",
		"ｨ": "ィ",
		"ｩ": "ゥ",
		"ｪ": "ェ",
		"ｫ": "ォ",
		"ｬ": "ャ",
		"ｭ": "ュ",
		"ｮ": "ョ",
		"ｯ": "ッ",
		"ｰ": "ー",
		"ｱ": "ア",
		"ｲ": "イ",
		"ｳ": "ウ",
		"ｴ": "エ",
		"ｵ": "オ",
		"ｶ": "カ",
		"ｷ": "キ",
		"ｸ": "ク",
		"ｹ": "ケ",
		"ｺ": "コ",
		"ｻ": "サ",
		"ｼ": "シ",
		"ｽ": "ス",
		"ｾ": "セ",
		"ｿ": "ソ",
		"ﾀ": "タ",
		"ﾁ": "チ",
		"ﾂ": "ツ",
		"ﾃ": "テ",
		"ﾄ": "ト",
		"ﾅ": "ナ",
		"ﾆ": "ニ",
		"ﾇ": "ヌ",
		"ﾈ": "ネ",
		"ﾉ": "ノ",
		"ﾊ": "ハ",
		"ﾋ": "ヒ",
		"ﾌ": "フ",
		"ﾍ": "ヘ",
		"ﾎ": "ホ",
		"ﾏ": "マ",
		"ﾐ": "ミ",
		"ﾑ": "ム",
		"ﾒ": "メ",
		"ﾓ": "モ",
		"ﾔ": "ヤ",
		"ﾕ": "ユ",
		"ﾖ": "ヨ",
		"ﾗ": "ラ",
		"ﾘ": "リ",
		"ﾙ": "ル",
		"ﾚ": "レ",
		"ﾛ": "ロ",
		"ﾜ": "ワ",
		"ﾝ": "ン",
	}
	dakuten = map[string]string{
		"ｶ": "ガ",
		"ｷ": "ギ",
		"ｸ": "グ",
		"ｹ": "ゲ",
		"ｺ": "ゴ",
		"ｻ": "ザ",
		"ｼ": "ジ",
		"ｽ": "ズ",
		"ｾ": "ゼ",
		"ｿ": "ゾ",
		"ﾀ": "ダ",
		"ﾁ": "ヂ",
		"ﾂ": "ヅ",
		"ﾃ": "デ",
		"ﾄ": "ド",
		"ﾊ": "バ",
		"ﾋ": "ビ",
		"ﾌ": "ブ",
		"ﾍ": "ベ",
		"ﾎ": "ボ",
		"ｳ": "ヴ",
	}
	handakuten = map[string]string{
		"ﾊ": "パ",
		"ﾋ": "ピ",
		"ﾌ": "プ",
		"ﾍ": "ペ",
		"ﾎ": "ポ",
	}
)

// ConvertHankakuToZenkaku convert hankaku katakana to zenkaku.
func ConvertHankakuToZenkaku(nameChan chan string, name string) {
	var preview, now, result string
	for _, c := range name {
		if unicode.In(c, unicode.Katakana) || strings.EqualFold(string(c), "ﾞ") || strings.EqualFold(string(c), "ﾟ") {
			if strings.EqualFold(string(c), "ﾞ") {
				now = dakuten[preview]
				preview = ""
			} else if strings.EqualFold(string(c), "ﾟ") {
				now = handakuten[preview]
				preview = ""
			} else if strings.EqualFold(hankakuToZenkaku[string(c)], "") == false {
				if strings.EqualFold(dakuten[string(c)], "") == false || strings.EqualFold(handakuten[string(c)], "") == false {
					if strings.EqualFold(preview, "") == false {
						result += string(hankakuToZenkaku[preview])
					}
					preview = string(c)
					continue
				}
				if strings.EqualFold(preview, "") == false {
					now = hankakuToZenkaku[preview]
					preview = ""
				}
				now += hankakuToZenkaku[string(c)]
			} else {
				now = string(c)
			}
			result += now
			now = ""
		} else {
			result += string(c)
		}
	}
	if strings.EqualFold(preview, "") == false {
		result += hankakuToZenkaku[preview]
	}
	nameChan <- result
}
