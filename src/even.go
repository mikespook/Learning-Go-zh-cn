package even		|\coderemark{Start our own namespace}|

func Even(i int) bool {	|\coderemark{Exported function}|
	return i % 2 == 0
}

func odd(i int) bool {	|\coderemark{Private function}|
	return i % 2 == 1
}
