package functions

import (
	"strings"
)

func AllFunc(content string) string {
	lines := strings.SplitAfter(content, "\n")
	for i, line := range lines {
		lines[i] = HundleMarks(HundlPunctuations(HundlVowel(
			HundlHexAndBin(Hundleflg(line)))))
	}
	return strings.Join(lines, "")
}
