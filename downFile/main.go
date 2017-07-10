package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func downLoad(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	file, _ := os.Open("test.txt")
	defer file.Close()
	w.Header().Set("Content-Disposition", "attachment; filename=savefile.txt")
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	io.Copy(w, file)
}

func main() {
	http.HandleFunc("/", downLoad)
	if err := http.ListenAndServe(":9000", nil); err != nil {
		fmt.Println(err)
	}
}
