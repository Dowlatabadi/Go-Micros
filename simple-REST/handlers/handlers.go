package handlers
import( 
	"net/http"
	"log"
	"encoding/json"
	"io"
)
type post_data struct {
	 link string
	 user_id string

}

type response_data struct {
	Shortened string `json:"Shortened"`

}

type Hello struct {
l *log.Logger
}
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}
func (h1 *post_data) ProcessRequest(  r io.Reader){ 

e:= json.NewDecoder(r)

e.Decode(h1)
}
func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request){
	
	
	
	
	if r.Method == http.MethodGet {

	h.l.Println("Hello World from Handler1")
	
	
	}
	
	if r.Method == http.MethodPost {
		h.l.Println("sadsad")
post:=&post_data{"sad","u1"}

h.l.Println(post.link)
post.ProcessRequest(r.Body)
h.l.Println(post.link)
//preapring result
en:=json.NewEncoder(rw) 
rd:= &response_data{"SRT"}
en.Encode(rd)
//fmt.Fprintf(rw,"Hello %s",post.link)
	}

}

