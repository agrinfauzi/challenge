package main 

import (
"fmt"
"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Halaman Beranda")
}

func intro(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w, "Hello World")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/intro", intro)
	http.ListenAndServe(":8282", nil)
}
