import "fmt"
func maxSumOfThreeSubarrays(nums []int, k int) []int {
    sum := make([]int, len(nums)+1)
    posLeft := make([]int, len(nums))
    posRight := make([]int, len(nums))
    result := make([]int, 3)
    if len(nums) == 3*k{
        result[0], result[1], result[2] = 0, k, 2*k
        return result
    }
    // Iterate over the array and calculate the sum of previous items for each item in the array
    for i:=0; i< len(nums); i++{ 
        sum[i+1] = sum[i] + nums[i]
    }
    // Calculate the left window, posLeft[i] is the starting index of the max sliding window till i
    mol:= sum[k] - sum[0]
    for i:=k; i<len(nums)-2*k; i ++ {
        if sum[i+1] - sum[i+1-k] > mol {
            posLeft[i] = i-k+1
            mol = sum[i+1] - sum[i+1-k]
        }else{
            posLeft[i] = posLeft[i-1]
        }
    }
    
    // Calculate the right window
    posRight[len(nums)-k] = len(nums)-k
    mor:=sum[len(nums)] - sum[len(nums)-k]
    for j:=len(nums)-1-k; j>=2*k; j--{
        if sum[j+k] - sum[j] >= mor {
            posRight[j] = j  
            mor = sum[j+k] - sum[j]
        }else{
            posRight[j] = posRight[j+1]
        }
    }
    max := 0
    for mid:=k; mid<=len(nums) - 2*k; mid ++{
        // sumThree := leftSum + midSum + rightSum
        leftStart:=posLeft[mid-1]
        rightStart:= posRight[mid+k]
        leftSum:= sum[leftStart+k] - sum[leftStart]
        rightSum:= sum[rightStart+k] - sum[rightStart]
        current := leftSum + rightSum + sum[mid+k] - sum[mid]  
        if max < current {
            max = current
            result[0], result[1], result[2] = leftStart, mid, rightStart
        }
    }
    return result
}