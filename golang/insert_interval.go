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

// 解法2 参考了 solution 
func insert(intervals []Interval, newInterval Interval) []Interval {
	if len(intervals) == 0 {
		return append(intervals, newInterval)
	}
    newIntervals := make([]Interval,0)

	for i, cur := range intervals {
		if cur.End < newInterval.Start {
            newIntervals = append(newIntervals, cur)
        }else if cur.Start > newInterval.End {
			newIntervals = append(newIntervals, newInterval)
            newIntervals = append(newIntervals, intervals[i:]...)
            break
        }else{
		    // overlap cases
            if newInterval.Start > cur.Start {
                newInterval.Start = cur.Start
            }
            if newInterval.End < cur.End {
                newInterval.End = cur.End
            }
        }
		if i == len(intervals)-1 {
            newIntervals = append(newIntervals, newInterval)
		}
	}
	return newIntervals
}
