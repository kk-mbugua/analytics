package utils

func CreateSerialNumber(prefix string, uuid string) (string, error) {
	serialNumber := prefix + uuid[24:]
	return serialNumber, nil
}

func GetAverage(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0.0
	}

	sum := 0.0
	for _, num := range numbers {
		sum += num
	}

	average := sum / float64(len(numbers))
	return average
}

func GetPercentage(numerator int64, denominator int64) float32 {
	if denominator == 0 {
		return 0.0
	}
	return float32(numerator) / float32(denominator) * 100.0
}

func ContainsString(arr []string, target string) bool {
	for _, s := range arr {
		if s == target {
			return true
		}
	}
	return false
}
