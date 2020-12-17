package advent

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// PasswordIsOk evaluates the advent password rule
func PasswordIsOk(minChars int, maxChars int, char string, input string) bool {
	count := strings.Count(input, char)
	return count >= minChars && count <= maxChars
}

// PasswordIsOkv2 evaluates the advent password rule for round 2
func PasswordIsOkv2(aIndex int, bIndex int, char string, input string) bool {
	runeinput := []rune(input)
	runechar := []rune(char)[0]
	fmt.Println(runeinput, runechar, input, char, aIndex, bIndex)
	return (runeinput[aIndex-1] == runechar || runeinput[bIndex-1] == runechar) && (runeinput[aIndex-1] != runeinput[bIndex-1])
}

// StringsFromFile gets a list of input cases from a file
func StringsFromFile(filename string) (cases []string) {
	file, err := os.Open(filename)
	check(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		cases = append(cases, scanner.Text())
	}
	return
}

// ExtractPieces get the components of a input case
func ExtractPieces(testcase string) (int, int, string, string) {
	testcase = strings.Replace(testcase, "-", " ", -1)
	testcase = strings.Replace(testcase, ":", "", -1)
	var minChars, maxChars int
	var char, input string
	_, err := fmt.Sscanf(testcase, "%d %d %v %v", &minChars, &maxChars, &char, &input)
	check(err)
	return minChars, maxChars, char, input
}

// ValidPasswordCount checks a passwords file and returns the count of valid ones
func ValidPasswordCount(filename string) (count int) {
	cases := StringsFromFile(filename)
	for _, testcase := range cases {
		if PasswordIsOk(ExtractPieces(testcase)) {
			count++
		}
	}
	return
}

// NewValidPasswordCount finished advent 2 pt 2
func NewValidPasswordCount(filename string) (count int) {
	cases := StringsFromFile(filename)
	for _, testcase := range cases {
		if PasswordIsOkv2(ExtractPieces(testcase)) {
			count++
		}
	}
	return
}
