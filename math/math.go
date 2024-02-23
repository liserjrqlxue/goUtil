package math

import (
	"math"
	"sort"
)

func SumFloat(x []float64) (sum float64) {
	for _, v := range x {
		sum += v
	}
	return
}

func Mean(x []float64) (mean float64) {
	return Sum(x) / float64(len(x))
}

func MeanVar(x []float64) (mean, variance float64) {
	mean = Mean(x)
	for _, v := range x {
		variance += (v - mean) * (v - mean)
	}
	return mean, variance / float64(len(x)-1)
}

// MeanStdDev
func MeanStdDev(x []float64) (mean, stdDev float64) {
	mean, stdDev = MeanVar(x)
	return mean, math.Sqrt(stdDev)
}

// DivisionInt return float64(x)/float64(y)
func DivisionInt(x, y int) float64 {
	return float64(x) / float64(y)
}

func Sum[G int | float64](x []G) (sum G) {
	for _, v := range x {
		sum += v
	}
	return
}

// SumInt sum int
func SumInt(x []int) (sum int) {
	for _, v := range x {
		sum += v
	}
	return
}

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Integer interface {
	Signed | Unsigned
}

type Float interface {
	~float32 | ~float64
}

type Ordered interface {
	Integer | Float | ~string
}

type Pair[T Ordered, U any] struct {
	sortable   T
	unsortable U
}

// SortSlice sorts the elements of the sortable slice and reorders the unsortable slice accordingly.
func SortSlice[T Ordered, U any](sortable []T, unsortable []U) {
	// 创建一个辅助结构体，用于存储 sortable 和 unsortable 中的元素对

	// 创建一个 Pair 的 slice，用于存储 sortable 和 unsortable 中的元素对
	pairs := make([]Pair[T, U], len(sortable))
	for i := range sortable {
		pairs[i] = Pair[T, U]{sortable[i], unsortable[i]}
	}

	// 按照 sortable 的顺序对 pairs 进行排序
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].sortable < pairs[j].sortable
	})

	// 将排序后的 unsortable 的值赋回给原始的 unsortable slice
	for i, pair := range pairs {
		unsortable[i] = pair.unsortable
	}
}
