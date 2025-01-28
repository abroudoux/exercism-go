package techpalace

import (
	"fmt"
	"strings"
)

func WelcomeMessage(customer string) string {
	return fmt.Sprintf("Welcome to the Tech Palace, %s", strings.ToUpper(customer))
}

func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	border := strings.Repeat("*", numStarsPerLine)
	return fmt.Sprintf("%s\n%s\n%s", border, welcomeMsg, border)
}

func CleanupMessage(oldMsg string) string {
	cleanedMessage := strings.ReplaceAll(oldMsg, "*", "")
    cleanedMessage = strings.ReplaceAll(cleanedMessage, "\n", "")
	cleanedMessage = strings.TrimSpace(cleanedMessage)
    return cleanedMessage
}

