package main

func average(xs []float64) (ave float64) { |\longremark{可以使用命名返回值；}|
	sum := 0.0
	switch len(xs) {
	case 0:                 |\longremark{如果长度是零，返回 0；}|
		ave = 0
	default:                |\longremark{否则，计算平均值；}|
		for _, v := range xs {
			sum += v
		}
		ave = sum / float64(len(xs)) |\longremark{为了使除法能正常计算，必须将值转换为 \key{float64}；}|
	}
	return  |\longremark{得到平均值，返回它}|
}
