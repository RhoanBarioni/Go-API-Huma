package serviceapi

// Esse cara é um helper, organizar ele.
func CalcMedia(notas []float32) float32 {
	if len(notas) == 0 {
		return 0
	}
	var nota float32
	for _, n := range notas {
		nota += n
	}

	return nota / float32(len(notas))
}
