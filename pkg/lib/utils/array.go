package utils

//FindMatchString finds target string in the array
func FindMatchString(target string, array []string) bool {
	for _, val := range array {
		if val == target {
			return true
		}
	}
	return false
}
