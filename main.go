package main

import (
	"awesomeProject2/src/novel"
	"net/http"

)


func e2searchbook(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	w.Header().Set("Content-Type","text/html;charset=utf-8")
	xsname:=r.Form["xsname"][0]
	novel.Searchnovel(xsname)
}

func main() {
	http.HandleFunc("/e2searchbook",e2searchbook)
	http.ListenAndServe("localhost:8080",nil)
}
