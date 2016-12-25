package model
import (
	"testing"
	"strings"
)


func TestDecode(t *testing.T){
	strFromFile := GetStringFromFile(
		"src/main/com/codeconveyor/MorseDecoder/resources/test.txt")

	strDecoded :=
		SeparateLeters(
			GetStringFromFile(
				"src/main/com/codeconveyor/MorseDecoder/resources/morse.txt"));
	result := DecodeArrayOfStrings(strDecoded,nil)

	if(strings.EqualFold(result, strFromFile)){
		t.Error("Error comparing")
	}
}