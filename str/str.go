package str

import "unicode"

// CamelCaseToUnderscore 驼峰单词转下划线单词
//
// e.g. "DoSomething" => "do_something"
func CamelCaseToUnderscore(str string) string {
	var output []rune
	for i, s := range str {
		if i == 0 {
			output = append(output, unicode.ToLower(s))
			continue
		}
		if unicode.IsUpper(s) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(s))
	}
	return string(output)
}
