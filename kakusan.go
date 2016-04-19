package kakusan

import "unicode"

var (
	hankakuToZenkaku = map[string]int{
		"ｦ": 0x30F2,
		"ｧ": 0x30A1,
		"ｨ": 0x30A3,
		"ｩ": 0x30A5,
		"ｪ": 0x30A7,
		"ｫ": 0x30A9,
		"ｬ": 0x30E3,
		"ｭ": 0x30E5,
		"ｮ": 0x30E7,
		"ｯ": 0x30C3,
		"ｰ": 0x30FC,
		"ｱ": 0x30A2,
		"ｲ": 0x30A4,
		"ｳ": 0x30A6,
		"ｴ": 0x30A8,
		"ｵ": 0x30AA,
		"ｶ": 0x30AB,
		"ｷ": 0x30AD,
		"ｸ": 0x30AF,
		"ｹ": 0x30B1,
		"ｺ": 0x30B3,
		"ｻ": 0x30B5,
		"ｼ": 0x30B7,
		"ｽ": 0x30B9,
		"ｾ": 0x30BB,
		"ｿ": 0x30BD,
		"ﾀ": 0x30BF,
		"ﾁ": 0x30C1,
		"ﾂ": 0x30C4,
		"ﾃ": 0x30C6,
		"ﾄ": 0x30C8,
		"ﾅ": 0x30CA,
		"ﾆ": 0x30CB,
		"ﾇ": 0x30CC,
		"ﾈ": 0x30CD,
		"ﾉ": 0x30CE,
		"ﾊ": 0x30CF,
		"ﾋ": 0x30D2,
		"ﾌ": 0x30D5,
		"ﾍ": 0x30D8,
		"ﾎ": 0x30DB,
		"ﾏ": 0x30DE,
		"ﾐ": 0x30DF,
		"ﾑ": 0x30E0,
		"ﾒ": 0x30E1,
		"ﾓ": 0x30E2,
		"ﾔ": 0x30E4,
		"ﾕ": 0x30E6,
		"ﾖ": 0x30E8,
		"ﾗ": 0x30E9,
		"ﾘ": 0x30EA,
		"ﾙ": 0x30EB,
		"ﾚ": 0x30EC,
		"ﾛ": 0x30ED,
		"ﾜ": 0x30EF,
		"ﾝ": 0x30F3,
	}
	dakuten = map[string]int{
		"ｶ": 0x30AC,
		"ｷ": 0x30AE,
		"ｸ": 0x30B0,
		"ｹ": 0x30B2,
		"ｺ": 0x30B4,
		"ｻ": 0x30B6,
		"ｼ": 0x30B9,
		"ｽ": 0x30BA,
		"ｾ": 0x30BC,
		"ｿ": 0x30BE,
		"ﾀ": 0x30C0,
		"ﾁ": 0x30C2,
		"ﾂ": 0x30C5,
		"ﾃ": 0x30C7,
		"ﾄ": 0x30C9,
		"ﾊ": 0x30D0,
		"ﾋ": 0x30D3,
		"ﾌ": 0x30D6,
		"ﾍ": 0x30D9,
		"ﾎ": 0x30DC,
		"ｳ": 0x30F4,
	}
	handakuten = map[string]int{
		"ﾊ": 0x30D1,
		"ﾋ": 0x30D4,
		"ﾌ": 0x30D7,
		"ﾍ": 0x30DA,
		"ﾎ": 0x30DD,
	}
)

// ConvertHankakuToZenkaku convert hankaku katakana to zenkaku.
func ConvertHankakuToZenkaku(nameChan chan string, name string) {
	var preview, now, result string
	for _, c := range name {
		if unicode.In(c, unicode.Katakana) || string(c) == "ﾞ" || string(c) == "ﾟ" {
			if string(c) == "ﾞ" {
				now = string(dakuten[preview])
				preview = ""
			} else if string(c) == "ﾟ" {
				now = string(handakuten[preview])
				preview = ""
			} else if hankakuToZenkaku[string(c)] != 0 {
				if dakuten[string(c)] != 0 || handakuten[string(c)] != 0 {
					if preview != "" {
						result += string(hankakuToZenkaku[preview])
					}
					preview = string(c)
					continue
				}
				if preview != "" {
					now = string(hankakuToZenkaku[preview])
					preview = ""
				}
				now += string(hankakuToZenkaku[string(c)])
			} else {
				now = string(c)
			}
			result += now
			now = ""
		} else {
			result += string(c)
		}
	}
	if preview != "" {
		result += string(hankakuToZenkaku[preview])
	}
	nameChan <- result
}
