package main 

import (
"fmt"
"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Halaman Beranda")
}

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8282", nil)
}