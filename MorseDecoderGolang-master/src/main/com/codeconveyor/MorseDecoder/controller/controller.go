package controller

import (
	"main/com/codeconveyor/MorseDecoder/model"
	"fmt"
	"time"
	"math/big"
	"log"
)

func Decode() {
	morseCode := model.GetStringFromFile("src/main/com/codeconveyor/MorseDecoder/resources/morse.txt")//view.GetMorseCodeFromConsole()
	var stringArray []string = model.SeparateLeters(morseCode)


	decodedStringParalel := Decode_paralel(stringArray)
	model.WriteStringToFile(decodedStringParalel)

	decodedStringSequential := Decode_Sequential(stringArray)
	fmt.Println(decodedStringSequential)

}

func Decode_paralel(stringArray []string) string{
	start := time.Now()

	r := new(big.Int)
	fmt.Println(r.Binomial(1000, 10))

	var c1 chan string = make(chan string)
	var c2 chan string = make(chan string)

	stringArrayForGorutine1 := stringArray[0:(len(stringArray)/2)]
	stringArrayForGorutine2 := stringArray[(len(stringArray)/2) + 1:]

	go model.DecodeArrayOfStrings(stringArrayForGorutine1, c1)
	go model.DecodeArrayOfStrings(stringArrayForGorutine2, c2)
	msg1 := <- c1
	msg2 := <- c2

	elapsed := time.Since(start)
	log.Printf("Paralel handle time %s", elapsed)
	return msg1+msg2
}

func Decode_Sequential(stringArray []string) string{
	start := time.Now()

	r := new(big.Int)
	fmt.Println(r.Binomial(1000, 10))

	result := model.DecodeArrayOfStrings(stringArray,nil)

	elapsed := time.Since(start)
	log.Printf("Sequential handle time %s", elapsed)
	return result
}
