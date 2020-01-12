package main

import(
	"net/http"
	"log"
)

var m = map[string]string{
	"go": "https://golang.org",
	"gci": "https://codein.withgoogle.com",
}

func redirect(w http.ResponseWriter, r *http.Request) {
	source := r.URL.Path[1:];
	if source != "" {
		http.Redirect(w, r, m[source], http.StatusFound)
	}
}

func main() {
    http.HandleFunc("/", redirect)
    
    log.Fatal(http.ListenAndServe("", nil));
}