package basics

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{}
	idx := 0
	res = append(res, intervals[0])

	for i := 1; i < len(intervals); i++ {
		if res[idx][1] >= intervals[i][0] {
			if res[idx][1] < intervals[i][1] {
				res[idx][1] = intervals[i][1]
			}
		} else {
			res = append(res, intervals[i])
			idx++
		}
	}

	return res
}

func TestMergeIntervals() {
	arr := [][][]int{
		{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
		{{1, 4}, {4, 5}},
		{{4, 7}, {1, 4}},
	}

	for _, v := range arr {
		fmt.Println("MergeIntervals, before=", v, ", after=", merge(v))
	}
}
