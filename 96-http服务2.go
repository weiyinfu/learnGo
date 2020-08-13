package main
import (
"log"
"net/http"
"time"
)

func main() {
	mux := http.NewServeMux()

	th := &timeHandler{format: time.RFC1123}
	mux.Handle("/time", th)

	log.Println("Listening...")
	http.ListenAndServe(":3000", mux)
}

type timeHandler struct {
	format string
}

func (th *timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(th.format)
	log.Println(r.Body)
	w.Write([]byte("time : " + tm))
	w.Header()["content-type"] =[]string{"application/json"}
}