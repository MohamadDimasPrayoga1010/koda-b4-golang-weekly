package utils

import (
	"fmt"
	"strings"
)

func FormatInt64ToRp(amount int64) string {
	s := fmt.Sprintf("%d", amount)
	n := len(s)

	if n <= 3 {
		return "Rp " + s
	}

	var result strings.Builder
	remainder := n % 3

	if remainder > 0 {
		result.WriteString(s[:remainder])
		if n > 3 {
			result.WriteString(".")
		}
	}

	for i := remainder; i < n; i += 3 {
		result.WriteString(s[i : i+3])
		if i+3 < n {
			result.WriteString(".")
		}
	}

	return "Rp " + result.String()
}
