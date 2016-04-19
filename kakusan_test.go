package kakusan

import (
	"bufio"
	"os"
	"testing"
)

func TestConvertHankakuToZenkaku(t *testing.T) {
	hankaku := readHankakuFromFile("./fixtures/katakana.txt")
	nameChan := make(chan string)
	go ConvertHankakuToZenkaku(nameChan, hankaku)
	result := <-nameChan

	stringAssert(t, result, "アイウエオカキクケコサシスセソタチツテトナニヌネノハヒフヘホマミムメモヤユヨワヲン")
}

func TestConvertHankakuToZenkaku2(t *testing.T) {
	hankaku := readHankakuFromFile("./fixtures/katakana2.txt")
	nameChan := make(chan string)
	go ConvertHankakuToZenkaku(nameChan, hankaku)
	result := <-nameChan

	stringAssert(t, result, "アイウェオ")
}

func TestConvertHankakuToZenkaku3(t *testing.T) {
	hankaku := readHankakuFromFile("./fixtures/katakana3.txt")
	nameChan := make(chan string)
	go ConvertHankakuToZenkaku(nameChan, hankaku)
	result := <-nameChan

	stringAssert(t, result, "カキクゲコ")
}

func TestConvertHankakuToZenkaku4(t *testing.T) {
	hankaku := readHankakuFromFile("./fixtures/katakana4.txt")
	nameChan := make(chan string)
	go ConvertHankakuToZenkaku(nameChan, hankaku)
	result := <-nameChan

	stringAssert(t, result, "ハヒプベホ")
}

func TestConvertHankakuToZenkaku5(t *testing.T) {
	hankaku := readHankakuFromFile("./fixtures/katakana5.txt")
	nameChan := make(chan string)
	go ConvertHankakuToZenkaku(nameChan, hankaku)
	result := <-nameChan

	stringAssert(t, result, "ABアイEFG")
}

func TestConvertHankakuToZenkaku6(t *testing.T) {
	hankaku := readHankakuFromFile("./fixtures/katakana6.txt")
	nameChan := make(chan string)
	go ConvertHankakuToZenkaku(nameChan, hankaku)
	result := <-nameChan

	stringAssert(t, result, "アイウエオ")
}

func TestConvertHankakuToZenkaku7(t *testing.T) {
	hankaku := readHankakuFromFile("./fixtures/katakana7.txt")
	nameChan := make(chan string)
	go ConvertHankakuToZenkaku(nameChan, hankaku)
	result := <-nameChan

	stringAssert(t, result, "岡本コウイチロウ")
}

func readHankakuFromFile(fileName string) string {
	file, _ := os.OpenFile(fileName, os.O_RDONLY, 777)
	buffer, _ := bufio.NewReader(file).ReadBytes('\n')
	return string(buffer)
}

func stringAssert(t *testing.T, s1, s2 string) {
	if s1 != s2 {
		t.Errorf("expected is %s\nbut was %s", s2, s1)
		t.Fail()
	} else {
		t.Logf("%s and %s is equal", s1, s2)
	}
}
