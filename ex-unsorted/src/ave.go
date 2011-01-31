package main

func average(xs []float64) (ave float64) { |\longremark{We use a named return parameter;}|
	sum := 0.0
	switch len(xs) {
	case 0:                 |\longremark{Is the length is zero, we return 0;}|
		ave = 0
	default:                |\longremark{Otherwise we calculate the average;}|
		for _, v := range xs {
			sum += v
		}
		ave = sum / float64(len(xs)) |\longremark{We have to convert the value to a
\key{float64} to make the division work;}|
	}
	return  |\longremark{We have an avarage, return it.}|
}
