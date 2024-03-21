package servicelayer

import (
	"strings"
)

func checkPasswordIsValid(Password, RePassword string) bool {	//probably should redo and make one for loop iteration (^_^)
	if Password != RePassword || len(Password) < 8 {
		return false
	}
	if containsInvalidChar(Password) && containsNumeric(Password) && containsSpecChar(Password) && containsLowerCase(Password) && containsUpperCase(Password) {
		return true
	}
	return true
}

func checkNameEmailIsValid(name, email string) bool {
	if len(name) < 4 || len(name) > 18 {
		return false
	} else if containsInvalidChar(name) || containsSpecChar(name) {
		return false
	} else if len(email) < 4 || len(email) > 254 || containsInvalidChar(email) {
		return false
	}
	return true
}

func containsNumeric(pass string) bool {
	return strings.ContainsAny(pass, "0123456789")
}

func containsSpecChar(pass string) bool {
	specChar := "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~ "
	for _, v := range pass {
		for _, spV := range specChar {
			if v == spV {
				return true
			}
		}
	}
	return false
}

func containsUpperCase(pass string) bool {
	for _, v := range pass {
		if v > 64 && v < 91 {
			return true
		}
	}
	return false
}

func containsLowerCase(pass string) bool {
	for _, v := range pass {
		if v > 96 && v < 123 {
			return true
		}
	}
	return false
}

func containsInvalidChar(pass string) bool {
	for _, v := range pass {
		if v <= 32 || v >= 127 {
			return true
		}
	}
	return false
}
