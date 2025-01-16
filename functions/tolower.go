package functions

func ToLower(s string) string {
	lower := ""
	for i := 0; i < len(s); i++ {
		group := s[i]
		if s[i] >= 'A' && s[i] <= 'Z' {
			group = (s[i]) + 32
		}
		lower = lower + string(group)
	}

	return lower
}
