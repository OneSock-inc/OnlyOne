package utils

func Contains(arr []string, obj string) bool {
	for _, elem := range arr {
		if elem == obj {
			return true
		}
	}
	return false
}
