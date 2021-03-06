package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	v := r.FormValue("q")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//from requestbody
	io.WriteString(w, `<form method="post"> <input type="text" name="q"> <input type="submit"> </form> <br> `+v)
	//from url
	io.WriteString(w, `<form method="get"> <input type="text" name="q"> <input type="submit"> </form> <br> `+v)
}
