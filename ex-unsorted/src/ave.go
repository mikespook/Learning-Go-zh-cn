package main

func average(xs []float32) (ave float32) {
	sum := 0.0
	switch len(xs) {
	case 0:
		ave = 0
	default:
		for _, v := range xs {
			sum += v
		}
		ave = sum / float32(len(xs))
	}
	return
}

func main() {
	/* ... */
}
