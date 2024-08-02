package str

import (
	"strconv"
	"unicode"
)

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

func String2Int64(s *string, intVal *int64) (bool, error) {
	i, err := strconv.ParseInt(*s, 10, 64)
	if err != nil {
		return false, err
	}
	if intVal != nil {
		*intVal = i
	}
	return true, nil
}

func String2Float64(s *string, intVal *float64) (bool, error) {
	i, err := strconv.ParseFloat(*s, 64)
	if err != nil {
		return false, err
	}
	if intVal != nil {
		*intVal = i
	}
	return true, nil
}
