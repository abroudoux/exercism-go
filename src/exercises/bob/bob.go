package bob

import "strings"

func Hey(remark string) string {
	trimmedRemark := strings.TrimFunc(remark, func(r rune) bool {
		return r == ' ' || r == '\\' || r == '\n' || r == '\t' || r == '\r'
	})
	if trimmedRemark == "" {
		return "Fine. Be that way!"
	}

	if isRemarkAQuestion(trimmedRemark) && isRemarkYelled(trimmedRemark) {
		return "Calm down, I know what I'm doing!"
	}

	if isRemarkAQuestion(trimmedRemark) {
		return "Sure."
	}

	if isRemarkYelled(trimmedRemark) {
		return "Whoa, chill out!"
	}

	return "Whatever."
}

func isRemarkAQuestion(remark string) bool {
	return strings.HasSuffix(remark, "?")
}

func isRemarkYelled(remark string) bool {
	hasValidLetter := false
	for _, char := range remark {
		if strings.ContainsRune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", char) {
			hasValidLetter = true
		}
	}
	return hasValidLetter && remark == strings.ToUpper(remark)
}
