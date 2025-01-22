package functions

func RemoveEmptyStrings(slice []string) []string {
	str := []string{}
	for _, v := range slice {
		if v != "" {
			str = append(str, v)
		}
	}
	return str
}
