package main
import( 
	"net/http"
	"log"
	"encoding/json"
	"os"
	"myapp/handlers"
)
type s1 struct {
	Name string
	Num int
}
func handler1(rw http.ResponseWriter, r*http.Request){
log.Println("hello world")
s34:=s1{"as",34}
d,_:=json.Marshal(s34)

rw.Write(d)
}
func main(){

	log.Println("hi World")
	l:=log.New(os.Stdout,"product-api",log.LstdFlags)
	handler2:=handlers.NewHello(l)
	sm :=http.NewServeMux()
	sm.Handle("/",handler2)
	http.Handle("/",handler2)
	//http.HandleFunc("/goodbye",func(http.ResponseWriter, *http.Request){
	//log.Println("Goodbye World")
//})
	http.ListenAndServe(":9090",nil)
}