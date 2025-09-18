package parsinglogfiles

import (
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`\[[A-Z]{3}\]`)
	indexes := re.FindStringIndex(text)
	if indexes != nil && indexes[0] == 0 {
		match := text[indexes[0]:indexes[1]]
		if match == "[TRC]" || match == "[DBG]" || match == "[INF]" || match == "[WRN]" || match == "[ERR]" || match == "[FTL]" {
			return true
		}
	}
	return false
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~|*|=|-]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	pattern := `".*(?i)password.*"`
	re := regexp.MustCompile(pattern)
	var count int
	for _, line := range lines {
		if re.MatchString(line) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line[0-9]+`)
	matches := re.FindAllString(text, -1)
	s := text
	for _, match := range matches {
		s = strings.ReplaceAll(s, match, "")
	}
	return s
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User[ ]+`)
	for i, line := range lines {
		if re.MatchString(line) {
			username := strings.Split(re.Split(line, 2)[1], " ")
			lines[i] = "[USR] " + username[0] + " " + line
		}
	}
	return lines
}
