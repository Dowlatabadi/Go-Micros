package main
import( 
	"fmt"
	"math/rand"
	"time"
	"os"
	"strconv"
)
func main(){
	number_of_testcases,_:=  strconv.Atoi(os.Args[1])
	sum_duration :=time.Duration(0)
	sums:=[]int{}
	fmt.Println("cxvc")
	fmt.Println(number_of_testcases)
	for i:=0;i<number_of_testcases;i++{
		x:=rand.Intn(100)+1
		if x>0{

			fmt.Println(x)
		}
		sum:=0
		for l:=1;l<=x;l++{
			num:=rand.Intn(100)-50
			fmt.Print(num," ")
			if (num>0){
				start := time.Now()
				sum+=num*num

				elapsed := time.Since(start)
				sum_duration+=elapsed

			}

		}
		if x>0{
			sums=append(sums,sum)
		}

		fmt.Print("\n")

	}
	fmt.Println("")
	fmt.Println("---------------")
	//	for i:=0;i<len(sums);i++{
	//	fmt.Println(sums[i])
	//}
	fmt.Println("time of sum squaring",sum_duration)
}
