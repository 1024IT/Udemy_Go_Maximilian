package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloats(strings []string) ([]float64, error) {
	var floats []float64

	for _, stringVal := range strings {
		//strconv.ParseFloat()...指定したString型の値をFlaot型に変換している。
		floatVal, err := strconv.ParseFloat(stringVal, 64)

		if err != nil {
			return nil, errors.New("Failed to convert string to float.")
		}

		floats = append(floats, floatVal)
	}

	return floats, nil
}
