package even

func Even(i int) bool {
	if i % 2 == 0 {
		return true
	}
	return false
}

func uneven(i int) bool {
	if i % 2 != 0 {
		return false
	}
	return true
}
