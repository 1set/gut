package yrand

import (
	"errors"
	"fmt"
	"sort"
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

func WeightedShuffle(weights []float64, yield func(idx int) (err error)) (err error) {
	var (
		count   = len(weights)
		cumSum  = make([]float64, 0, count)
		sum     = 0.0
		randNum float64
	)
	if count <= 0 {
		err = fmt.Errorf("empty weight list")
		return
	}
	for _, weight := range weights {
		// add check here for weights like [1e-6, 1e30, 1e-3],
		if (weight <= 0) || (sum+weight == sum) || (sum-weight == sum) {
			err = fmt.Errorf("invalid weight 1&2")
			return
		}
		sum += weight
		cumSum = append(cumSum, sum)
	}

	for range weights {
		// get random pos
		if randNum, err = Float64(); err != nil {
			break
		}
		randNum *= cumSum[count-1]

		// binary search for the pos and yield it
		j := sort.Search(count, func(i int) bool { return cumSum[i] > randNum })
		if !((0 <= j) && (j < count)) {
			err = fmt.Errorf("invalid weight 3")
			break
		}
		if err = yield(j); err != nil {
			break
		}

		// remove weight from rest of the sum list
		for p := j; p < count; p++ {
			cumSum[p] -= weights[j]
		}
	}
	return
}
