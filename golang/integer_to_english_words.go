// 首先明白几个词：one, two, three, four, five, six, seven, eight, nine, ten, 
//              eleven, twelve, thirteen, Fourteen, Fifteen, sixteen, seventeen, eighteen, Nineteen, 
//              twenty, thirty, forty, fifty, sixty, seventy, eighty, ninety,
//              hundred,
//              thousand,
//              million
//              billion
var (
    list1 = []string{"Zero", "One", "Two", "Three", "Four", "Five","Six","Seven","Eight","Nine"}
    list2 = []string{"Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen","Eighteen", "Nineteen"}
    list3 = []string{"", "","Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
)
// below code would be much clean if golang supports ternary operation 
func numberToWords(i int) string {
    if i>= 1000000000{
            if i%1000000000 == 0 {
                return numberToWords(i/1000000000) + " " + "Billion"
            }
            return numberToWords(i/1000000000) + " " + "Billion" + " " +  numberToWords(i%1000000000)
        }else if i >= 1000000{
            if i%1000000 == 0{
                return numberToWords(i/1000000) + " " + "Million"
            }
            return numberToWords(i/1000000) + " " + "Million" + " " +  numberToWords(i%1000000)
        }else if i >= 1000{
            if i%1000 == 0{
                 return numberToWords(i/1000) + " " + "Thousand"
            }
            return numberToWords(i/1000) + " " + "Thousand" + " " + numberToWords(i%1000)
        }else if i >= 100{
            if i%100==0{
                return list1[i/100] + " "+ "Hundred"
            }
            return list1[i/100] + " "+ "Hundred" + " " + numberToWords(i%100)
        }else if i >= 20{
            if i%10 == 0{
                return list3[i/10]
            }else{
                return list3[i/10] + " " + numberToWords(i%10)
            }
        }else if i >= 10{
            return list2[i%10]
        }else {
            return list1[i%10]
        }
}
