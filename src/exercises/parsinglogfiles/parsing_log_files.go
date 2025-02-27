package parsinglogfiles

import (
	"fmt"
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=\-]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	pattern := `"[^"]*password[^"]*"`
	re := regexp.MustCompile(pattern)
	count := 0

	for _, line := range lines {
		lowercaseLine := strings.ToLower(line)
		matches := re.FindAllString(lowercaseLine, -1)
		count += len(matches)
	}

	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\w+)`)

	taggedLines := make([]string, len(lines))

	for i, line := range lines {
		match := re.FindStringSubmatch(line)
		if match != nil {
			userName := match[1]
			taggedLines[i] = fmt.Sprintf("[USR] %s %s", userName, line)
		} else {
			taggedLines[i] = line
		}
	}

	return taggedLines
}
