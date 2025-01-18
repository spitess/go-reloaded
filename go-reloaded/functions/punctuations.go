package functions

func HundlPunctuations(s string) string {
	slice := []rune(s)
	checkpunc := ".,!?:;"
	elment := []rune{' '}

	for i := 0; i < len(slice)-1; i++ {
		for _, v := range checkpunc {
			if slice[i] == ' ' && slice[i+1] == v {
				slice = append(slice[:i], slice[i+1:]...)
				i--
			}
			if slice[i] == v && slice[i+1] != v {
				slice = append(slice[:i+1], append(elment, slice[i+1:]...)...)
			}
		}
	}

	return string(slice)
}
