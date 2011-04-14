package even		|\coderemark{开始自定义的包}|

func Even(i int) bool {	|\coderemark{可导出函数}|
	return i % 2 == 0
}

func odd(i int) bool {	|\coderemark{私有函数}|
	return i % 2 == 1
}
