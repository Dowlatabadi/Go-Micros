package main
import(
	"fmt"
	"bufio"
	"strconv"
	"strings"
	"io"
	"sync"
	"os"
)
//reads lines on array items and ignores negative ones and converts to int array recursively
func Read_Array_Elements(reader *bufio.Reader,Elements int) []int{
	if Elements==0{
		return []int{}
	}
	string_array:=strings.Split(strings.TrimSpace(readLine(reader))," ")
	line_array:=Convert2intArray(string_array)
	length:=len(string_array)
	return append(Read_Array_Elements(reader,Elements-length),line_array...)
}
//reads an array length and next the actual array items
func Read_Array_Block(reader *bufio.Reader,block_num int,channel chan []int) {
	if (block_num==0) {
		return 
	}
	length,err:= strconv.Atoi(readLine(reader))
	checkError(err)
	res:= Read_Array_Elements(reader,length)
	channel	<- res
	Read_Array_Block(reader,block_num-1,channel)

}
//converts an array of strings to an int array recursively
func Convert2intArray(stringArray []string) []int {
	result:=[]int{}
	if len(stringArray)==0{
		return result
	} else{
		element,err:=strconv.Atoi(stringArray[0] )
		checkError(err)
		tTemp:=[]int{}
		if (element>0){

			tTemp = []int {element}
		}
		second_part:=Convert2intArray(stringArray[1:])
		return append(tTemp,second_part...)
	}
}
//calculate sum of squares of items using a recirsive halfing approach
func sum_squares(inputArray []int) int{
	l:=len(inputArray)
	if l==0{
		return 0
	} else if l==1{
		return inputArray[0]*inputArray[0]
	} else{
		mid:=l/2
		left:=inputArray[:mid]
		right:=inputArray[mid:]
		return sum_squares(left)+sum_squares(right)
	}

}
//implementation of concurrent calculation of elements of an array which in less than 100 items is not making a big difference, therefore not using here
//func sum_squares_concurrent(inputArray []int,input_wg *sync.WaitGroup,input_channel chan int) {
//	//	defer (*input_wg).Done()
//	l:=len(inputArray)
//	//	fmt.Println(l)
//	if l==0{
//		input_channel<- 0
//	} else if l<=1000{
//		temp:=sum_squares(inputArray)
//
//		input_channel <-temp
//		//	fmt.Println(inputArray[0],"array")
//	} else{
//		result:= make(chan int,2)
//		mid:=l/2
//		//		fmt.Println("length",l,"mid",mid)
//		left:=inputArray[:mid]
//		right:=inputArray[mid:]
//		wg :=sync.WaitGroup{}
//		wg.Add(2)
//		go sum_squares_concurrent(left,&wg,result)
//		go sum_squares_concurrent(right,&wg,result)
//		wg.Wait()
//		l_result:=<-result
//		r_result:=<-result
//		close(result)
//		input_channel<-l_result+r_result 
//	}
//
//}
//reads a line to the end as a  string
func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(str), "\r\n")
}
//the function to check the error
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
//main consumer to consume and calculate the sum of squares
func consume( channel1 *chan []int,result_channel chan int, wg *sync.WaitGroup,counter int) {
	if counter<=0{
		return
	} else{
		defer (*wg).Done()
	}
	msg1 := <-*channel1
	go consume(channel1,result_channel,wg,counter-1)
	result_channel <- sum_squares(msg1)
}
//function to print the input channel items recursively
func Print_channel(channel chan int){
	fmt.Println(<- channel)
	if len(channel)>0{
		Print_channel(channel)
	}
}
func main(){
	reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)
	arrays , err:= strconv.Atoi(readLine(reader))
	checkError(err)
	channel:=make(chan []int,arrays)
	wg:=sync.WaitGroup{}
	wg.Add(arrays)
	ch:=make(chan int,arrays)
	//fire up the consumer before the actual reading
	go consume(&channel,ch,&wg,arrays) 
	Read_Array_Block(reader,arrays,channel )
	wg.Wait()
	close(channel)
	close(ch)
	Print_channel(ch)
}
