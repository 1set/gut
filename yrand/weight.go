package yrand

import (
	"errors"
	"sort"
)

var (
	errInvalidWeights = errors.New("invalid weight list")
	errInvalidIndex   = errors.New("invalid index")
	tolerance         = 1e-14
)

var (
	// QuitShuffle is used as a return value from ShuffleIndexFunc to indicate that the execution of WeightedShuffle should be terminated immediately.
	// It is not returned as an error by any function.
	QuitShuffle = errors.New("quit this shuffle")
)

type (
	// ShuffleIndexFunc is the type of the function called for each random index selected by WeightedShuffle.
	ShuffleIndexFunc func(idx int) (err error)
	// ShuffleSwapFunc is the type of the function called by Shuffle to swap the elements with indexes i and j.
	ShuffleSwapFunc func(i, j int)
)

// WeightedChoice selects a random index according to the associated weights (or probabilities).
//
// Indexes with zero or negative weight value will be ignored.
//
// The slice of associated weights must contain at least one positive value.
func WeightedChoice(weights []float64) (idx int, err error) {
	var (
		sum     = 0.0
		randNum float64
	)
	// get sum of weights
	for _, w := range weights {
		if w > 0 {
			sum += w
		}
	}
	if sum <= 0 {
		err = errInvalidWeights
		return
	}

	// get random value
	if randNum, err = Float64(); err != nil {
		return
	}
	sum *= randNum

	// find the random pos
	for i, w := range weights {
		if w > 0 {
			sum -= w
			if sum < 0 {
				idx = i
				return
			}
		}
	}
	idx = len(weights) - 1
	return
}

// WeightedShuffle shuffles the sequence of values according to the associated weights (or probabilities).
//
// All values in the slice of associated weights must be positive.
func WeightedShuffle(weights []float64, indexFunc ShuffleIndexFunc) (err error) {
	var (
		count   = len(weights)
		cumSum  = make([]float64, 0, count)
		sum     = 0.0
		nextSum = 0.0
		randNum float64
	)

	for _, weight := range weights {
		// check non-positive weight, and weights like [1e30, 1e-6, 1e30],
		if nextSum = sum + weight; (weight <= 0) || !isFloatEqual(nextSum-weight, sum, tolerance) || !isFloatEqual(nextSum-sum, weight, tolerance) {
			err = errInvalidWeights
			break
		}
		sum = nextSum
		cumSum = append(cumSum, sum)
	}
	if err != nil || sum <= 0 {
		err = errInvalidWeights
		return
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
			err = errInvalidIndex
			break
		}
		if err = indexFunc(j); err != nil {
			break
		}

		// remove weight from rest of the sum list
		for p := j; p < count; p++ {
			cumSum[p] -= weights[j]
		}
	}

	if err == QuitShuffle {
		err = nil
	}
	return
}
