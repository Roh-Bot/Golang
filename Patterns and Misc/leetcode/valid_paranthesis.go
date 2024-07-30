package leetcode

func ValidParanthesis(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	paranthesisOpen := map[string]string{
		"{": "}",
		"(": ")",
		"[": "]",
	}

	var openers, closures string
	if v, ok := paranthesisOpen[string(s[0])]; !ok {
		return false
	} else if v == string(s[1]) {
		openers, closures = Continuos(paranthesisOpen, s)
	} else {
		openers, closures = NonContinuos(paranthesisOpen, s)
	}

	if (len(openers) != len(closures)) || (len(openers) == 0 || len(closures) == 0) {
		return false
	}

	for i := 0; i < len(openers); i++ {
		if !(paranthesisOpen[string(openers[i])] == string(closures[i])) {
			return false
		}
	}
	return true
}

func Continuos(paranthesisOpen map[string]string, s string) (openers, closures string) {
	for i := 0; i < len(s); i++ {
		if _, ok := paranthesisOpen[string(s[i])]; ok {
			openers += string(s[i])
			continue
		} else {
			closures += string(s[i])
		}
	}
	return
}

func NonContinuos(paranthesisOpen map[string]string, s string) (openers, closures string) {
	for i := 0; i < len(s); i++ {
		if _, ok := paranthesisOpen[string(s[i])]; ok {
			openers += string(s[i])
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if _, ok := paranthesisOpen[string(s[i])]; !ok {
			closures += string(s[i])
		}
	}
	return
}
