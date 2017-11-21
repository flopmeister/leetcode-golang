/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */

 type IntervalSlice []Interval
 func (it IntervalSlice)Swap(i,j int){
	 it[i], it[j] = it[j], it[i]
 }
 
 func (it IntervalSlice)Less(i, j int) bool{
	 return it[i].Start < it[j].Start
 }
 func (it IntervalSlice)Len() int{
	 return len(it)
 }
 
 func merge(intervals []Interval) []Interval {
	 if len(intervals) <=1 {
		 return intervals
	 }
	 it := IntervalSlice(intervals)
	 sort.Sort(it)
	 i := 1
	 for i<len(it) {
		 // when to merge
		 if it[i].Start <= it[i-1].End{
			 if it[i].End >= it[i-1].End{
				 it[i-1].End = it[i].End
			 }
			 it = append(it[:i], it[i+1:]...)
		 }else{
			 i++
		 }
	 }
	 return []Interval(it)
 }
 