package even			|\coderemark{Our own namespace}|

func Even(i int) bool {		|\coderemark{Exported}|
	if i % 2 == 0 {
		return true
	}
	return false
}

func uneven(i int) bool {	|\coderemark{Private}|
	if i % 2 != 0 {
		return false
	}
	return true
}
