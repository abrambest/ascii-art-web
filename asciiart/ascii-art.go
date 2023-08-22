package asciiart

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func readAscii(font string) []string {
	files := font + ".txt"

	file, err := os.ReadFile("asciiart/banners/" + files)
	if err != nil {
		fmt.Println(err)
	}
	str0 := ""
	file = []byte(strings.ReplaceAll(string(file), "\r", ""))
	arrSplit := strings.Split(string(file), "\n\n")
	for i, v := range arrSplit[0] {
		if i > 0 {
			str0 += string(v)
		}

	}
	arrSplit[0] = str0

	return arrSplit

}

func printAsciiArt(txt, arrSplit []string) string {

	str := ""
	for _, word := range txt {

		if word == "\n" {
			str += "\n"
			continue
		}

		for k := 0; k < 8; k++ {

			for _, s := range word {
				if s >= 32 {
					strCut := strings.Split(arrSplit[s-32], "\n")
					for _, ascii := range strCut[k] {

						str += string(ascii)
					}
				}

			}
			str += "\n"

		}

	}
	return str

}

func checkTxt(str string) error {

	noChar := ""
	for _, s := range str {
		if s < 0 || s > 127 {
			noChar += string(s)
		}

	}
	if noChar != "" {
		return errors.New(fmt.Sprintf("a character \"%v\" is not available.", noChar))
	}
	return nil
}
func check(errMsg string, err error) {
	if err != nil {
		fmt.Println(errMsg, err)
		os.Exit(1)
	}
}

func AsciiFunc(txt, font string) (string, error) {

	if txt == "" {
		return "", nil
	}

	err := checkTxt(txt)
	if err != nil {

		return "", err
	}

	arrTxt := strings.Split(txt, "\n")

	arrSplit := readAscii(font)
	return printAsciiArt(arrTxt, arrSplit), nil

}
