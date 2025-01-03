package functions

func ToUpper(s string) string {
	upper := ""
	for i := range s {
		alex := s[i]
		if s[i] >= 'a' && s[i] <= 'z' {
			alex = (s[i]) - 32
		}
		upper = upper + string(alex)
	}
	return upper
}
