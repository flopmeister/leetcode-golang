// result := right_btm - right_top - left_btm + left_top
// sum[1][1] = matrix[0][0]
// sum[1][0] = matrix[]


type NumMatrix struct {
	Sums [][] int
 }
 
 
 func Constructor(matrix [][]int) NumMatrix {
     if len(matrix) == 0 {
         return NumMatrix{}
     }
	 sums:=make([][]int, len(matrix) +1)
	 for i := range sums {
		 sums[i] = make([]int, len(matrix[0]) + 1)
	 }
	 for i:=0; i< len(matrix); i ++{
		 for j:=0; j<len(matrix[0]); j++{
				 sums[i+1][j+1] = sums[i][j+1] + sums[i+1][j] - sums[i][j] + matrix[i][j]
		 }
	 }
	 return NumMatrix{sums}
 }
 
 
 func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	 // top_left = row1, col1
	 // top_right = row1, col2
	 // btm_left = row2, col1
	 // btm_right = row2, col2
	 return this.Sums[row1][col1] + this.Sums[row2+1][col2+1] - this.Sums[row2+1][col1] - this.Sums[row1][col2+1]
 }
 
 
 /**
  * Your NumMatrix object will be instantiated and called as such:
  * obj := Constructor(matrix);
  * param_1 := obj.SumRegion(row1,col1,row2,col2);
  */