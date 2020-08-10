package perfect

import "errors"

// Classification is the type of the natural number
type Classification int

// ErrOnlyPositive is the error when a non positive number is inputed
var ErrOnlyPositive = errors.New("Must be a positive number")

// Three types of natural numeber
const (
	ClassificationDeficient Classification = iota
	ClassificationPerfect
	ClassificationAbundant
)

// Classify classifies a natural number
func Classify(num int64) (Classification, error) {
	if num <= 0 {
		return -1, ErrOnlyPositive
	}

	var aliquot int64 = 0
	for i := int64(1); i < num; i++ {
		if num%i == 0 {
			aliquot += i
		}
	}

	var class Classification
	if aliquot == num {
		class = ClassificationPerfect
	} else if aliquot > num {
		class = ClassificationAbundant
	} else {
		class = ClassificationDeficient
	}
	return class, nil
}
