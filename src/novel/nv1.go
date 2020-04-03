package novel

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)
var header="Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.131 Safari/537.36"
type Bookdetailinfos struct {
	Title string `json:"title"`
	Author string `json:"author"`
	Cover string `json:"cover"`
	Url string `json:"url"`
	Introduce string `json:"introduce"`
	Nvtype string `json:"nvtype"`
	Updatatime string `json:"updatatime"`
	Lastchapter string `json:"lastchapter"`
	Lastchapterlink string `json:"lastchapterlink"`
	Status string `json:"status"`
	Booksize string `json:"booksize"`
}
type Chapter struct {
	Chaptername string `json:"num"`
	Chapterlink string `json:"url"`
}
type Detailinfos struct {
	Data Bookdetailinfos `json:"data"`
	List []Chapter `json:"list"`
}
type Bookcontent struct {
	Title string `json:"title"`
	Content []string `json:"content"`
}
func Searchnovel(xsname string) []Bookdetailinfos {
	var bookinfos = make([]Bookdetailinfos,0)
	var p = url.Values{}
	p.Add("keyword", xsname)
	client := &http.Client{} // 建立client
	value:=p.Encode()
	url :="https://www.xsbiquge.com/search.php?"+value
	req,err := http.NewRequest("GET",url,nil)
	defer LogError()
	if err != nil {
		panic(err)
	}
	req.Header.Add("User-Agent",header)
	resp,err := client.Do(req)
	defer resp.Body.Close()
	//LogContent(resp)
	if resp.StatusCode != 200 {
		panic(fmt.Sprintf("status code error: %d %s", resp.StatusCode, resp.Status))
	}
	dom, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}
	items:=dom.Find(".result-item,.result-game-item")
	items.Each(func(i int, selection *goquery.Selection) {
		its:=selection.Find("a[href]")
		ps:=selection.Find("p[class=result-game-item-desc]")
		infos:=selection.Find("p[class=result-game-item-info-tag]")
		url:=its.Get(0).Attr[1].Val
		cover:=its.Get(0).FirstChild.NextSibling.Attr[0].Val
		title:=its.Get(1).Attr[2].Val
		introduce:=ps.Text()
		author:=infos.Get(0).LastChild.PrevSibling.LastChild.Data
		author=strings.Replace(author," ","",-1)
		author=strings.Replace(author,"\n","",-1)
		nvtype:=infos.Get(1).LastChild.PrevSibling.LastChild.Data
		updatatime:=infos.Get(2).LastChild.PrevSibling.LastChild.Data
		lastart:=infos.Get(3).LastChild.PrevSibling.LastChild.Data
		lastartlink:=infos.Get(3).LastChild.PrevSibling.Attr[1].Val
		bookinfos=append(bookinfos, Bookdetailinfos{title,author,cover,url,introduce,nvtype,updatatime,lastart,lastartlink,"",""})
	})

	return bookinfos
}
func Noveldetail(url string)Detailinfos{
	var chapterlist = make([]Chapter,0)
	var bookinfos = Bookdetailinfos{}

	client := &http.Client{} // 建立client
	req,err := http.NewRequest("GET",url,nil)
	defer LogError()
	if err != nil {
		panic(err)
	}
	req.Header.Add("User-Agent",header)
	resp,err := client.Do(req)
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		panic(fmt.Sprintf("status code error: %d %s", resp.StatusCode, resp.Status))
	}
	//LogContent(resp)
	dom, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}
	items:=dom.Find(".box_con").Find("dd").Find("a")
	items.Each(func(i int, selection *goquery.Selection) {
		chaptername:=selection.Get(0).Attr[0].Val
		chapterurl:=selection.Get(0).FirstChild.Data
		chapterlist=append(chapterlist, Chapter{chaptername,chapterurl})
	})
	item2s:=dom.Find(".box_con").First()
	item2s.Each(func(i int, selection *goquery.Selection) {
		maininfo:=selection.Find("#maininfo")
		title:=maininfo.Find("h1").Text()
		author:=maininfo.Find("p").Get(0).FirstChild.Data
		author=cutst(author)
		status:=maininfo.Find("p").Get(1).FirstChild.Data
		status=cutst(status)
		lastupdata:=maininfo.Find("p").Get(2).FirstChild.Data
		lastupdata=cutst(lastupdata)
		lastchapter:=maininfo.Find("p").Get(3).FirstChild.NextSibling.FirstChild.Data
		lastchapterlink:=maininfo.Find("p").Get(3).LastChild.Attr[0].Val
		cover, _ :=selection.Find("#sidebar").Find("img").Attr("src")
		introduce:=selection.Find("#intro").Find("p").First().Text()
		bookinfos=Bookdetailinfos{title,author,cover,url,introduce,"",lastupdata,lastchapter,lastchapterlink,status,""}
	})
	detailinfo:= Detailinfos{}
	detailinfo.Data=bookinfos
	detailinfo.List=chapterlist
	return detailinfo

}
///90_90760/272984.html
func Novelcontent(url string) Bookcontent {
	url="https://www.xsbiquge.com"+url
	contents:=make([]string,0)
	var bookcontent = Bookcontent{}
	client := &http.Client{} // 建立client
	req,err := http.NewRequest("GET",url,nil)
	defer LogError()
	if err != nil {
		panic(err)
	}
	req.Header.Add("User-Agent",header)
	resp,err := client.Do(req)
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		panic(fmt.Sprintf("status code error: %d %s", resp.StatusCode, resp.Status))
	}
	//LogContent(resp)
	dom, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}
	dom.Find("#content").Find("br").Each(func(i int, selection *goquery.Selection) {
		content3:=selection.Get(0)
		fmt.Println(content3)
		content:=selection.Get(0).PrevSibling.Data
		contents=append(contents, content)
	})
	title:=dom.Find(".bookname").Nodes[0].FirstChild.NextSibling.FirstChild.Data
	bookcontent.Content=contents
	bookcontent.Title=title
	return bookcontent
}
func cutst(st string) (returnst string){
	returnst=st
	index := strings.Index(returnst, "：")+ 3
	if index==-1 {

	}else {
		returnst=returnst[index:]
	}
	return
}
func LogError(){
		if err := recover(); err != nil { //产生了panic异常
			fmt.Println(err)
		}
}
func LogContent(reader *http.Response)  {
	body, err:= ioutil.ReadAll(reader.Body) // 读取内容
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", body) // 以字符串形式打印
}