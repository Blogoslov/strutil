package strutil

import "strings"

//Wordwrap wraps the given string str based on colLen as max line width.
//It never breaks the words, even if it is longer than the colLen. It breaks
//the string on " ", ",", ".", ":", ";"
func Wordwrap(str string, colLen int) string {
	if colLen == 0 {
		return str
	}
	runes := []rune(str)
	var buff strings.Builder
	var prevChar string
	l := len(runes)

	//left trim
	startOfLine := findRealStartOfLine(runes, 0)
	var endOfLine int

	for i := startOfLine; i < l; i++ {
		c := string(runes[i])

		switch {
		case c == "\n":
			buff.WriteString(string(runes[startOfLine : i+1]))
			startOfLine = findRealStartOfLine(runes, i+1)
			endOfLine = startOfLine

		case i-startOfLine > colLen && endOfLine != startOfLine:
			buff.WriteString(string(runes[startOfLine:endOfLine]))
			buff.WriteString("\n")
			startOfLine = findRealStartOfLine(runes, endOfLine)
			endOfLine = startOfLine

		case c == " " && prevChar != " ":
			endOfLine = i
		case c == "," || c == "." || c == ":" || c == ";":
			endOfLine = i + 1
		}

		if startOfLine >= len(runes) {
			break
		}
		if startOfLine > i {
			i = startOfLine
		}

		//index (i) may be changed in the loop, so do not use "c" directly
		prevChar = string(runes[i])
	}
	if startOfLine < len(runes) {
		buff.WriteString(string(runes[startOfLine:]))
	}

	return buff.String()
}

//findRealStartOfLine is a helper function for Wordwrap. In practice it is used
//to trim the left side of the runes. It gets startOfLine as beginning index
//and returns new index if there are spaces after startOfLine
func findRealStartOfLine(runes []rune, startOfLine int) int {
	for ; startOfLine < len(runes) && string(runes[startOfLine]) == " "; startOfLine++ {
	}

	return startOfLine
}