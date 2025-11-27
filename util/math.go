package util

func GetAvg(datas []float64) float64 {
	var mean float64 = 0
	for i, data := range datas {
		mean += (data - mean) / float64(i+1)
	}
	return mean
}
