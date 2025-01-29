package functions

import (
	goreload "goreload/functions"
	"strings"
)

func AllFunc(content string) string {
	lines := strings.SplitAfter(content, "\n")
	for i, line := range lines {
		lines[i] = goreload.HundlPunctuations(goreload.HundleMarks(goreload.HundlVowel(
		goreload.HundlHexAndBin(goreload.Hundleflg(line),),),),)
	}
	return strings.Join(lines, "")
}
