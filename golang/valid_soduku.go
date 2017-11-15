
func isValidSudoku(board [][]byte) bool {
    // 遍历每一个row
    // 遍历每一个column
    // 遍历每一个subbox
    // sudoku puzzle must be length = width
    rows := [9][9]int{} // 9 rows 
    columns := [9][9]int{} // 9 columns
    boxes := [9][9]int{}// 9 boxes
    
    for i:=0; i<9;i++{
        for j:=0; j<9; j++{
            if board[i][j]!='.'{
                num := board[i][j] -'0' -1
                k := i / 3 * 3 + j / 3
                if rows[i][num]==1 || columns[j][num]==1 || boxes[k][num]==1{
                    return false
                }else{
                    rows[i][num] = 1
                    columns[j][num] = 1
                    boxes[k][num] = 1
                }
            }
        }
    }
    return true
}