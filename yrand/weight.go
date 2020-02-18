package yrand

import (
	"container/list"
	"errors"
)

var (
	errInvalidWeights = errors.New("invalid weight list")
	errInvalidIndex   = errors.New("invalid index")
	tolerance         = 1e-7
)

// WeightedChoice selects a random index according to the associated weights (or probabilities).
//
// Indexes with zero or negative weight value will be ignored.
//
// The slice of associated weights must contain at least one positive value.
//
// The complexity is O(n) where n = len(weights).
func WeightedChoice(weights []float64) (idx int, err error) {
	var (
		sum = 0.0
		rnd float64
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
	if rnd, err = Float64(); err != nil {
		return
	}
	sum *= rnd

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

// WeightedShuffle randomizes the order of elements according to the associated weights (or probabilities).
//
// All values in the slice of associated weights must be positive, and values of very different magnitudes are unacceptable.
//
// The yieldFunc will be called for each randomly selected index.
//
// The complexity is O(n^2) where n = len(weights).
func WeightedShuffle(weights []float64, yieldFunc ShuffleIndexFunc) (err error) {
	type weightNode struct {
		index  int
		weight float64
	}
	var (
		rnd float64
		sum = 0.0
		el  *list.Element
		wl  = list.New()
	)

	// check if it's an empty or invalid weight list
	if len(weights) == 0 {
		err = errInvalidWeights
		return
	}
	for i, w := range weights {
		if w <= 0 {
			err = errInvalidWeights
			return
		}
		_ = wl.PushBack(&weightNode{index: i, weight: w})
	}

	for range weights {
		sum = 0.0
		for el = wl.Front(); el != nil; el = el.Next() {
			w := el.Value.(*weightNode)
			sum += w.weight
		}

		// get random value
		if rnd, err = Float64(); err != nil {
			break
		}
		sum *= rnd

		// find the random pos
		for el = wl.Front(); el != nil; el = el.Next() {
			wn := el.Value.(*weightNode)
			if sum -= wn.weight; sum < 0 {
				break
			}
		}
		if el == nil {
			el = wl.Back()
		}

		// yield it and remove for next iteration
		wn := wl.Remove(el).(*weightNode)
		if err = yieldFunc(wn.index); err != nil {
			break
		}
	}

	if err == QuitShuffle {
		err = nil
	}
	return
}
