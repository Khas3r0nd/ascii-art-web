package function

import (
	"fmt"
	"os"
	"strings"
)

func ReadFont(banner string) (map[rune][]string, error) {
	input, err := os.ReadFile(banner)
	if err != nil {
		return nil, err
	}

	standardList := strings.Split(string(input), "\n")
	s := make(map[rune][]string)
	var startPoint rune = ' '
	var temp []string
	for i, v := range standardList {
		i++
		if i%9 == 0 {
			temp = append(temp, v)
			s[startPoint] = temp
			startPoint++
			temp = nil
		} else if v != "" {
			temp = append(temp, v)
		}
	}

	return s, nil
}

func PrintFormat(text string, format map[rune][]string) (string, error) {
	text = strings.ReplaceAll(text, "\\n", "\n")
	answer := ""
	if len(text) < 1 {
		return "", nil
	}

	line := strings.Split(text, "\n")
	for _, word := range line {
		if len(word) < 1 {
			answer += string('\n')
			continue
		}
		for i := 0; i < 8; i++ {
			for _, char := range word {
				if v, exists := format[char]; exists {
					answer += v[i]
				} else {
					if char != '\n' && char != 13 {
						return "", fmt.Errorf("Incorrect symbol found %c", char)
					}
				}
			}
			answer += string('\n')
		}
		answer += string('\n')
	}
	return answer, nil
}
