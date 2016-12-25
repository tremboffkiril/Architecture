package model

import (
	"strings"
	//"fmt"
	"os"
)

var MORSE_TABLE = map[string]string{
	""	: " ",
	".-"	: "A",
	"-..."	: "B",
	"-.-."	: "C",
	"-.."	: "D",
	"."	: "E",
	"..-."	: "F",
	"--."	: "G",
	"...."	: "H",
	".."	: "I",
	".---"	: "J",
	"-.-"	: "K",
	".-.."	: "L",
	"--"	: "M",
	"-."	: "N",
	"---"	: "O",
	".--."	: "P",
	"--.-"	: "Q",
	".-."	: "R",
	"..."	: "S",
	"-"	: "T",
	"..-"	: "U",
	"...-"	: "V",
	".--"	: "W",
	"-..-"	: "X",
	"-.--"	: "Y",
	"--.."	: "Z",
	"-----"	: "0",
	".----"	: "1",
	"..---"	: "2",
	"...--"	: "3",
	"....-"	: "4",
	"....."	: "5",
	"-...."	: "6",
	"--..."	: "7",
	"---.."	: "8",
	"----."	: "9",
}

func SeparateLeters(morseCode string) []string{
	morseCode = strings.Replace(morseCode,"\n"," \n",-1)
	var str []string = strings.Split(morseCode," ")

	//WriteStringToFile(result)
	return str
}

func DecodeArrayOfStrings(str []string, c chan string) string{
	var result string
	for _, i:= range str{
		//fmt.Print(MORSE_TABLE[i])
		result += MORSE_TABLE[i]
	}
	if(c != nil) {
		c <- result
	}
	return result
}

func GetStringFromFile(fileName string) string{
	file, err := os.Open(fileName)
	if err != nil {
		// handle the error here
		return "error"
	}
	defer file.Close()

	// get the file size
	stat, err := file.Stat()
	if err != nil {
		return "error"
	}
	// read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return "error"
	}

	return string(bs)
}

func WriteStringToFile(str string){
	file, err := os.Create("src/main/com/codeconveyor/MorseDecoder/resources/test.txt")
	if err != nil {
		// handle the error here
		return
	}
	defer file.Close()

	file.WriteString(str)
}

