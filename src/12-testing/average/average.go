package average

func Media(xi ...int) float64 {
	var resultado int
	for _, v := range xi {
		resultado += v
	}
	return float64(resultado) / float64(len(xi))
}
