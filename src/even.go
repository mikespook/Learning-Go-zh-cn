package even		|\coderemark{Start our own namespace}|

func Even(i int) bool {	|\coderemark{Exported}|
	return i % 2 == 0
}

func odd(i int) bool {	|\coderemark{Private}|
	return i % 2 == 1
}
