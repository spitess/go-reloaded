package functions

import (
	// "fmt"
	"strings"
)

func HundleMarks(s string) string {
	output := ""
	if !strings.Contains(s, "'") {
		return s
	} else {

		res := strings.Split(s, "'")
		if len(res) > 1 {
			res[1] = strings.TrimSpace(res[1])

			output = res[0] + "'" + res[1] + "'" + res[2]
		}

		return output

	}
}
