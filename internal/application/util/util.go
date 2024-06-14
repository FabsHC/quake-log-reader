package util

import (
	"regexp"
	"strings"
)

func IsGameStarting(logMessage string) bool {
	return containsIgnoreCase(logMessage, INIT_GAME)
}

func IsKillLog(logMessage string) bool {
	return containsIgnoreCase(logMessage, KILL)
}

func containsIgnoreCase(s, substr string) bool {
	s = strings.ToLower(s)
	substr = strings.ToLower(substr)
	return strings.Contains(s, substr)
}

func ExtractKillData(logMessage string) (string, string, string) {
	re := regexp.MustCompile(`\d{1,2}:\d{2} Kill: (\d+) (\d+) (\d+): (.+?) killed (.+?) by ([^ ]+)`)
	match := re.FindStringSubmatch(logMessage)
	return strings.TrimSpace(match[4]), strings.TrimSpace(match[5]), strings.TrimSpace(match[6])
}
