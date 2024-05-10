package utils

import "strconv"

func IntToString(i int) string {
	return strconv.Itoa(i)
}

func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		// Handle error, return default value, or propagate error as needed
		return 0 // Return 0 as default value in case of error
	}
	return i
}
