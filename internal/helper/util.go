package helper

func IsInStringList(target string, list []string) bool {
	found := false
	for i := 0; i < len(list); i++ {
		if list[i] == target {
			found = true
			break
		}
	}
	return found
}

func IsInBoolList(target bool, list []bool) bool {
	found := false
	for i := 0; i < len(list); i++ {
		if list[i] == target {
			found = true
			break
		}
	}
	return found
}
