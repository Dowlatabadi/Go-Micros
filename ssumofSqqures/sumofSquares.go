package main
import(
	"fmt"
	"bufio"
	"strconv"
	"strings"
	"io"
	"sync"
	"math/rand"
	"time"
)
func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}
func Convert2intArray(stringArray []string) []int {
	result:=[]int{}
	if len(stringArray)==0{
		return result
	} else{
		element,err:=strconv.Atoi(stringArray[0] )
		checkError(err)
		tTemp := []int {element}
		fmt.Println(tTemp)
		second_part:=Convert2intArray(stringArray[1:])
		return append(tTemp,second_part...)
	}
}
func sum_squares(inputArray []int) int{
	l:=len(inputArray)
	if l==1{


		return inputArray[0]*inputArray[0]
	} else{
		mid:=l/2
		//		fmt.Println("length",l,"mid",mid)
		left:=inputArray[:mid]
		right:=inputArray[mid:]
		return sum_squares(left)+sum_squares(right)
	}

}
func sum_squares_concurrent(inputArray []int,input_wg *sync.WaitGroup,input_channel chan int) {
	defer (*input_wg).Done()
	l:=len(inputArray)
	//	fmt.Println(l)
	if l==0{
		input_channel<- 0
	} else if l<=1000{
		temp:=sum_squares(inputArray)

		input_channel <-temp
		//	fmt.Println(inputArray[0],"array")
	} else{
		result:= make(chan int,2)
		mid:=l/2
		//		fmt.Println("length",l,"mid",mid)
		left:=inputArray[:mid]
		right:=inputArray[mid:]
		wg :=sync.WaitGroup{}
		wg.Add(2)
		go sum_squares_concurrent(left,&wg,result)
		go sum_squares_concurrent(right,&wg,result)
		wg.Wait()
		l_result:=<-result
		r_result:=<-result
		close(result)
		input_channel<-l_result+r_result 
	}

}
func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(str), "\r\n")
}
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
func consume( channel1 *chan []int,result_channel chan int, wg *sync.WaitGroup,done chan bool,counter int) {
if counter==0{
	fmt.Println("recursion")
	done <- true
return
}
	msg1 := <-*channel1

go		consume(channel1,result_channel,wg,done,counter-1)
		result_channel <- sum_squares(msg1)
		fmt.Println(msg1)
		(*wg).Done()
}
func produce(channel_name chan []int){

	for i:=1;i<2000000;i++{
		x :=[]int{}
		for y:=0;y<=i;y++{
			x=append(x,y)
		}
		channel_name <- x 
		fmt.Printf("is in ch %v \n",i)
	}
	close(channel_name)
}
func main(){
	wg:=sync.WaitGroup{}
	arrays:=500
	wg.Add(arrays)
	ch:=make(chan int,arrays)
	done:=make(chan bool)
	ch_arrays:=make(chan []int)
	go func(){
		for i:=1;i<=arrays;i++{
			rand_num:=rand.Intn(100)
			arr:=makeRange(1,rand_num)
			ch_arrays <- arr

		}
	}()
	time.Sleep(10 * time.Millisecond)
	go consume(&ch_arrays,ch,&wg,done,arrays) 
	<- done
		close(ch_arrays)
	close(ch)
		fmt.Println("final results----------------------")
	for sdfs:= range ch{
		fmt.Println("final",sdfs)
	}
	fmt.Println("%v",len(ch))	
}
