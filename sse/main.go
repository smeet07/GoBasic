package main
import (
	"fmt"
	"net/http"
	"flag"
	"time"
)
func main(){
	flag.Parse()
	http.HandleFunc("/",home)
	http.HandleFunc("/events",events)
	http.ListenAndServe(":8080", nil)

}
func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w,r, "index.html")
}
func events(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","text/event-stream")
	tokens :=[] string{"token1","token2","token3"}
	for _,token:=range tokens{
		content:=fmt.Sprintf("data: %s\n\n",string(token))
		w.Write([]byte(content))
		w.(http.Flusher).Flush()
		time.Sleep(time.Millisecond * 500)
	}

}
