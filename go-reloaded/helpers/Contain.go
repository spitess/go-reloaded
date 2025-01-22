package functions

func Contain(strs []string, sub string) bool {

	for _, val := range strs {
		if val == sub {
			return true
		}
	}
	return false
}
