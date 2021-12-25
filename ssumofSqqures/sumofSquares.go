package main
import(
	"fmt"
	"bufio"
	"strconv"
	"strings"
	"io"
	"sync"
	"time"
	"os"
)
func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}
func Read_Array_Elements(reader *bufio.Reader,Elements int) []int{
	if Elements==0{
return []int{}
}
line_array:=Convert2intArray(strings.Split(strings.TrimSpace(readLine(reader)), " "))
length:=len(line_array)
return append(Read_Array_Elements(reader,Elements-length),line_array...)
}
func Read_Array_Block(reader *bufio.Reader,block_num int,channel chan []int) {
	if (block_num==0) {
		return 
	}
	length,_:= strconv.Atoi(readLine(reader))
	res:= Read_Array_Elements(reader,length)
channel	<- res
	fmt.Println("received : ",res)
	Read_Array_Block(reader,block_num-1,channel)
	
}
func Convert2intArray(stringArray []string) []int {
	result:=[]int{}
	if len(stringArray)==0{
		return result
	} else{
		element,err:=strconv.Atoi(stringArray[0] )
		checkError(err)
		tTemp := []int {element}
//		fmt.Println(tTemp)
		second_part:=Convert2intArray(stringArray[1:])
		return append(tTemp,second_part...)
	}
}
func sum_squares(inputArray []int) int{
	l:=len(inputArray)
	if l==0{
		return 0
	} else if l==1{


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
	//	defer (*input_wg).Done()
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
func consume( channel1 *chan []int,result_channel chan int, wg *sync.WaitGroup,counter int) {
	if counter<=0{
		fmt.Println("recursion")
		return
	} else{
		defer (*wg).Done()
	}
	msg1 := <-*channel1
	go consume(channel1,result_channel,wg,counter-1)
	result_channel <- sum_squares(msg1)
	//		fmt.Println(msg1)
}
func read_input(channel chan []int){
	reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)




	length , _:= strconv.Atoi(readLine(reader))


	Read_Array_Block(reader,length,channel )


}
func main(){
	arrays:=10000
	channel:=make(chan []int,arrays)
	read_input(channel)
	//---------------------------sequential part

	//concurrent --------------------------------------------------------
	wg:=sync.WaitGroup{}
	wg2:=sync.WaitGroup{}
	wg.Add(arrays)
	ch:=make(chan int,arrays)
	wg2.Add(1)
	go func(){
		for i:=0;i<arrays;i++{
		}
		wg2.Done()
	}()
	wg2.Wait()
	start:=time.Now()
	go consume(&channel,ch,&wg,arrays) 
	wg.Wait()
	fmt.Println("time og con ",time.Since(start))
	close(channel)
	close(ch)
	fmt.Println("final results----------------------")
	//	for sdfs:= range ch{
	//		fmt.Println("final",sdfs)
	//	}
	fmt.Println("%v",len(ch))	
}
