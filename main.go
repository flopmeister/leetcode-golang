package main

import "sort"

func main() {

}

// 题目如果给的不是set， 需要加一个移除duplicates
// 第一反应，sort 然后back track

// reverse the 2D array, then flip [i,j] to [j,i]
// reverse the 2D array, then flip [i,j] to [j,i]

/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */

// 解法1 ： 根据merge interval 思路， 先插入，再merge
func insert(intervals []Interval, newInterval Interval) []Interval {
	target := sort.Search(len(intervals), func(i int) bool {
		return intervals[i].Start >= newInterval.Start
	})
	intervals = append(intervals[:target], append([]Interval{newInterval}, intervals[target:]...)...)
	i := target
	for ; i < len(intervals); i++ {
		if i == 0 {
			continue
		}
		// when to merge
		if intervals[i].Start <= intervals[i-1].End {
			if intervals[i].End >= intervals[i-1].End {
				intervals[i-1].End = intervals[i].End
			}
			intervals = append(intervals[:i], intervals[i+1:]...)
			i--
		}
	}
	return intervals
}

func insert(intervals []Interval, newInterval Interval) []Interval {
	if len(intervals) == 0 {
		return append(intervals, newInterval)
	}

	for i, cur := range len(intervals) {
		if cur.End < newInterval.Start {
			continue
		}
		if cur.Start > newInterval.End {
			intervals = append(intervals[:i], append([]Interval{newInterval}, intervals[i:]...)...)
			break
		}
		// overlap cases
		if newInterval.Start > cur.Start {
			newInterval.Start = cur.Start
		}
		if newInterval.End < cur.End {
			newInterval.End = cur.End
		}
		if i == len(intervals)-1 {
			intervals = append(intervals, newInterval)
		}
	}
	return intervals
}
