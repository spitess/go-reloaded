package functions

import (
	"strings"
)

func Convert(slice []string) string {
	return strings.Join(slice, " ")
}
