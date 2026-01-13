package common

func StringsContain(dst string, lst ...string) bool {
	for _, s := range lst {
		if s == dst {
			return true
		}
	}
	return false
}
