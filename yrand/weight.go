package yrand

import (
	"errors"
)

var (
	errInvalidWeights = errors.New("invalid weight list")
)

func WeightedChoice(weights []float64) (idx int, err error) {
	var (
		count   = len(weights)
		sum     = 0.0
		randNum float64
	)
	for _, w := range weights {
		if w > 0 {
			sum += w
		}
	}
	if sum <= 0 {
		err = errInvalidWeights
		return
	}

	if randNum, err = Float64(); err != nil {
		return
	}

	sum *= randNum
	for i, w := range weights {
		if w > 0 {
			sum -= w
			if sum < 0 {
				idx = i
				return
			}
		}
	}

	idx = count - 1
	return
}
