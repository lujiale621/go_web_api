package main

import (
	"awesomeProject2/src/novel"
	"awesomeProject2/src/util"
	"fmt"
	"net/http"
)


func e2searchbook(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	w.Header().Set("Content-Type","text/html;charset=utf-8")
	xsname:=r.Form["xsname"][0]
	res:=novel.Searchnovel(xsname)
	presendmessage:=""
	if len(res)<1 {
		presendmessage=util.Deal(nil)
	} else {
		presendmessage=util.Deal(res)
	}
	fmt.Fprint(w,presendmessage)
}
func e2bookdetail(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	w.Header().Set("Content-Type","text/html;charset=utf-8")
	xsdetail:=r.Form["xsdetail"][0]
	res:=novel.Noveldetail(xsdetail)
	presendmessage:=""
	if len(res.List)<5 {
		presendmessage=util.Deal(nil)
	}else {
		presendmessage=util.Deal(res)
	}

	fmt.Fprint(w,presendmessage)

}
func e2bookcontent(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	w.Header().Set("Content-Type","text/html;charset=utf-8")
	xscontent:=r.Form["xscontent"][0]
	presendmessage:=""
	res:=novel.Novelcontent(xscontent)
	if len(res.Content)<2 {
		presendmessage=util.Deal(nil)
	}else {
		presendmessage=util.Deal(res)
	}

	fmt.Fprint(w,presendmessage)
}
func main() {
	http.HandleFunc("/e2searchbook",e2searchbook)
	http.HandleFunc("/e2bookdetail",e2bookdetail)
	http.HandleFunc("/e2bookcontent",e2bookcontent)
	http.ListenAndServe("0.0.0.0:8768",nil)
}
