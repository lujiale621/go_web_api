package novel

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Bookdetailinfos struct {
	Title string
	Author string
	Cover string
	Url string
	Introduce string
	Nvtype string
	Updatatime string
	Lastchapter string
	Lastchapterlink string
}

func Searchnovel(xsname string) []Bookdetailinfos {
	var bookinfos = make([]Bookdetailinfos,0)
	var p = url.Values{}
	p.Add("keyword", xsname)
	client := &http.Client{} // 建立client
	value:=p.Encode()
	url :="https://www.xsbiquge.com/search.php?"+value
	req,err := http.NewRequest("GET",url,nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("User-Agent","Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.131 Safari/537.36")
	resp,err := client.Do(req)
	//LogContent(resp)
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
		bookinfos=append(bookinfos, Bookdetailinfos{title,author,cover,url,introduce,nvtype,updatatime,lastart,lastartlink})
	})
	if resp.StatusCode != 200 {
		panic("error")
		fmt.Println("status code error: %d %s", resp.StatusCode, resp.Status)
	}
	defer resp.Body.Close()
	defer func() {
		recover()
	}()
	return bookinfos
}
func LogContent(reader *http.Response)  {
	body, err:= ioutil.ReadAll(reader.Body) // 读取内容
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", body) // 以字符串形式打印
}