package AMJ

func AllEmpty(arr []string) bool {
	for _, word := range arr {
		if word != "" {
			return false
		}
	}
	return true
}
