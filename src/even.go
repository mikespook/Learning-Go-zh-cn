package even		|\coderemark{Start our own namespace}|

func Even(i int) bool {	|\coderemark{Exported}|
	if i % 2 == 0 {
		return true
	}
	return false
}

func odd(i int) bool {	|\coderemark{Private}|
	if i % 2 != 0 {
		return false
	}
	return true
}
