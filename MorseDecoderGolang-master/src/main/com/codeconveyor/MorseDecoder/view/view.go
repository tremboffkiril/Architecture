package view


import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

const headerStr = "\t\t----------.... MORSE DECODER ....----------\n"
const firstStr = "Enter your Morse code using '.' '-' and single space" +
	" for separating letters and double or triple space" +
	" for separating words : "
const errStr = "Error : your code is incorrect you have to use " +
	"'.' '-' and single space for separating letters and double" +
	" or triple space for separating words!\n Repeat please : "

func isCorrectMorseCode(entry ...string) bool {
	for _, str := range entry {
		if strings.Count(str, "-")+strings.Count(str, ".")+strings.Count(str, " ") != len(str) - 1 {
			return false
		}
	}
	return true
}
func GetMorseCodeFromConsole() string {
	fmt.Println(headerStr, firstStr)
	var res string
	reader := bufio.NewReader(os.Stdin)
	res, _ = reader.ReadString('\n')
	for !isCorrectMorseCode(res) {
		fmt.Println(errStr)
		res, _ = reader.ReadString('\n')
	}

	return res
}
