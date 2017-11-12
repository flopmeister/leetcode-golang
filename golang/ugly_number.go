
func isUgly(num int) bool {
    if num <= 0 {
        return false
    }
    if num == 1{
        return true
    }
    if num%2==0 && isUgly(num/2){
        return true
    } 
    if num%3==0 && isUgly(num/3){
        return true
    }
    if num%5==0 && isUgly(num/5){
        return true
    }
    return false
}

// 精简版
func isUgly(num int) bool {
	for i:=2; i<=5 && num>0; i++ {
		for num % i == 0{
			num /= i			
		}
	}
	return num == 1
}
