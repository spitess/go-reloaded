package functions

import (
	"strings"
)

func AllFunc(content string) string {
	lines := strings.SplitAfter(content, "\n")
	for i, line := range lines {
		lines[i] = HundlPunctuations(HundleMarks(HundlVowel(
			HundlHexAndBin(Hundleflg(line)))))
	}
	return strings.Join(lines, "")
}
