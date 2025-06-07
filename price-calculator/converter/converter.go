package converter

import (
	"errors"
	"fmt"
	"strconv"
)

func StringToFloats(str []string) ([]float64, error) {
	var floats []float64

	for _, stingVal := range str {
		floatPrice, err := strconv.ParseFloat(stingVal, 64)
		if err != nil {
			fmt.Println("Error parsing string to float64:", err)
			return nil, errors.New("error parsing string to float64")
		}
		floats = append(floats, floatPrice)
	}
	return floats, nil
}
