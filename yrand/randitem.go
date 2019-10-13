package yrand

import (
	"crypto/rand"
	"math/big"
)

// 10000	    226971 ns/op	   55944 B/op	    3996 allocs/op
func ShuffleV1(n int, swap func(i, j int)) {
	if n < 0 {
		panic("invalid argument to ShuffleV1")
	}

	// Fisher-Yates shuffle: https://en.wikipedia.org/wiki/Fisher%E2%80%93Yates_shuffle
	// ShuffleV1 really ought not be called with n that doesn't fit in 32 bits.
	// Not only will it take a very long time, but with 2³¹! possible permutations,
	// there's no way that any PRNG can have a big enough internal state to
	// generate even a minuscule percentage of the possible permutations.
	// Nevertheless, the right API signature accepts an int n, so handle it as best we can.
	i := n - 1
	for ; i > 1<<31-1-1; i-- {
		n, _ := Int64Range(0, int64(i+1))
		j := int(n)
		swap(i, j)
	}
	for ; i > 0; i-- {
		n, _ := Int32Range(0, int32(i+1))
		j := int(n)
		swap(i, j)
	}
}

// 48792	     49189 ns/op	    8238 B/op	       6 allocs/op
func ShuffleV2(n int, swap func(i, j int)) {
	//-- To shuffle an array a of n elements (indices 0..n-1):
	//for i from n−1 downto 1 do
	//j ← random integer such that 0 ≤ j ≤ i
	//exchange a[j] and a[i]

	randNum := make([]int, n)

	var stop error
	for i := uint64(n - 1); stop == nil; {
		stop = iterateRandomNumbers(n, i+1, func(r uint64) error {
			if r <= i {
				randNum[i] = int(r)
				i--
			}
			if i <= 0 {
				return errIterateCallback
			}
			return nil
		})
	}

	for i := n - 1; i > 0; i-- {
		j := randNum[i]
		swap(i, j)
	}
}

// 112395	     21663 ns/op	    8208 B/op	       3 allocs/op
func ShuffleV3(n int, swap func(i, j int)) {
	randBig := new(big.Int)
	randBytes := make([]byte, 8)

	randNum := make([]int, n)
	for upper := uint64(n - 1); upper > 0; {
		if _, err := rand.Read(randBytes); err != nil {
			return
		}

		randBig.SetBytes(randBytes)
		for num := randBig.Uint64(); num > 0 && upper > 0; upper-- {
			max := upper + 1
			randNum[upper] = int(num % max)
			num /= max
		}
	}

	for i := n - 1; i > 0; i-- {
		j := randNum[i]
		swap(i, j)
	}
}

// 115790	     19984 ns/op	      16 B/op	       2 allocs/op
func ShuffleV4(n int, swap func(i, j int)) {
	//-- To shuffle an array a of n elements (indices 0..n-1):
	//	for i from n−1 downto 1 do
	//		j ← random integer such that 0 ≤ j ≤ i
	//		exchange a[j] and a[i]

	randBig := new(big.Int)
	randBytes := make([]byte, 8)

	i := n - 1
done:
	for upper := uint64(n - 1); upper > 0; {
		if _, err := rand.Read(randBytes); err != nil {
			return
		}

		randBig.SetBytes(randBytes)
		for num := randBig.Uint64(); num > 0 && upper > 0; upper-- {
			max := upper + 1
			j := int(num % max)
			num /= max

			swap(i, j)
			if i > 0 {
				i--
			} else {
				goto done
			}
		}
	}
}

// 122638	     19532 ns/op	      16 B/op	       2 allocs/op
func ShuffleV5(n int, swap func(i, j int)) {
	//-- To shuffle an array a of n elements (indices 0..n-1):
	//	for i from n−1 downto 1 do
	//		j ← random integer such that 0 ≤ j ≤ i
	//		exchange a[j] and a[i]

	randBig := new(big.Int)
	randBytes := make([]byte, 8)

	for i := n - 1; i > 0; {
		if _, err := rand.Read(randBytes); err != nil {
			return
		}

		randBig.SetBytes(randBytes)
		for num := randBig.Uint64(); num > 0 && i > 0; i-- {
			max := uint64(i + 1)
			j := int(num % max)
			num /= max
			swap(i, j)
		}
	}
}

// 118530	     21787 ns/op	      16 B/op	       2 allocs/op
func ShuffleV6(n int, swap func(i, j int)) {
	//-- To shuffle an array a of n elements (indices 0..n-1):
	//	for i from n−1 downto 1 do
	//		j ← random integer such that 0 ≤ j ≤ i
	//		exchange a[j] and a[i]

	randBig := new(big.Int)
	randBytes := make([]byte, 8)

	for i := uint64(n - 1); i > 0; {
		if _, err := rand.Read(randBytes); err != nil {
			return
		}

		randBig.SetBytes(randBytes)
		for num := randBig.Uint64(); num > i && i > 0; i-- {
			max := i + 1
			j := int(num % max)
			num /= max
			swap(int(i), j)
		}
	}
}
