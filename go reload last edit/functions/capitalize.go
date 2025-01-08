package functions

func Capitalize(s string) string {
	strmani := []rune(s)

	for i := 0; i < len(strmani); i++ {
		if strmani[i] >= 'A' && strmani[i] <= 'Z' {
			strmani[i] = strmani[i] + 32
		}
	}
	for i := 0; i < len(strmani); i++ {
		if i == 0 || strmani[i-1] == ' ' {
			if strmani[i] >= 'a' && strmani[i] <= 'z' {
				strmani[i] = strmani[i] - 32
			}
		}
	}
	return string(strmani)
}
