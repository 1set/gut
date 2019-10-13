package yrand

import (
	"sort"
	"strconv"
	"strings"
	"testing"
)

func TestShuffle(t *testing.T) {
	num := []string{"a", "b", "c", "d", "e"}
	swap := func(i, j int) {
		num[i], num[j] = num[j], num[i]
	}
	t.Log(num)
	ShuffleV1(len(num), swap)
	t.Log(num)
	return
}

func BenchmarkShuffle(b *testing.B) {
	const count = 1000
	//const count = 10
	num, _ := rangeInt(count)
	for i := 0; i < count; i++ {
		num[i] *= 10
	}
	swapFunc := func(i, j int) {
		num[i], num[j] = num[j], num[i]
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ShuffleV6(count, swapFunc)
	}
}

func TestShuffleEdgeCase(t *testing.T) {
	var max int = 9223372036854775807
	cnt := 0
	cntFunc := func(i, j int) {
		cnt++
		t.Logf("[%d] %d <-> %d, 0 ≤ %d (j) ≤ %d (i), %v", cnt, i, j, j, cnt, j <= cnt)
	}
	ShuffleV6(max, cntFunc)
}

//	for i from n−1 downto 1 do
//		j ← random integer such that 0 ≤ j ≤ i
//		exchange a[j] and a[i]
func TestDemoShuffle(t *testing.T) {
	n := 3
	cnt := n - 1
	ShuffleV6(n, func(i, j int) {
		t.Logf("[%d] %d <-> %d, 0 ≤ %d (j) ≤ %d (i), %v", cnt, i, j, j, cnt, j <= cnt)
		cnt--
	})
}

func TestShuffleVerify(t *testing.T) {
	const count = 8
	const times = 1000000

	//const count = 9
	//const times = 5000000

	num, _ := rangeInt(count)
	swapFunc := func(i, j int) {
		num[i], num[j] = num[j], num[i]
	}

	expected := uint64(1)
	for i := uint64(2); i <= count; i++ {
		expected *= i
	}

	counters := map[string]int{}
	for i := 0; i < times; i++ {
		num, _ = rangeInt(count)
		ShuffleV6(count, swapFunc)
		str := numSlice2String(num)
		counters[str] += 1
	}

	total := len(counters)
	type mPair struct {
		key   string
		value int
	}
	pairs := make([]mPair, 0, total)
	for k, v := range counters {
		pairs = append(pairs, mPair{key: k, value: v})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].key < pairs[j].key
	})
	for _, p := range pairs {
		t.Logf("%q: %6d - %10.6f", p.key, p.value, float64(p.value)/float64(total))
	}

	t.Log("expect", expected)
	t.Log("actual", total)
}

func rangeInt(max int) (sl []int, err error) {
	if max < 0 {
		// TODO: use serious error later
		return nil, errIterateCallback
	}

	sl = make([]int, 0, max)
	for i := 0; i < max; i++ {
		sl = append(sl, i)
	}
	return
}

func numSlice2String(num []int) (s string) {
	if len(num) == 0 {
		return
	}

	str := make([]string, 0, len(num))
	for _, v := range num {
		str = append(str, strconv.Itoa(v))
	}
	return strings.Join(str, "|")
}

// Extract test to confirm that all cases will be covered evenly
