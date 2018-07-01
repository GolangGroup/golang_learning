package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)


type  Spider struct{
	ListUrl string
	DownBaseUrl string
	header map[string]string
}

func (keyword Spider) get_html_header() string{
	client := &http.Client{}
	fmt.Println(keyword.ListUrl)
	req, err := http.NewRequest("GET", keyword.ListUrl, nil)
	if err != nil {
		
	}
	for key, value := range keyword.header {
		req.Header.Add(key, value)
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	
	}
	return string(body)
}

func parse(){
	header := map[string]string{

		":authority": "www.dy2018.com",
		":method": "GET",
		":path": "/html/gndy/dyzz/index.html",
		":scheme": "https",
		"accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8",
		"accept-encoding": "gzip, deflate, br",
		"accept-language": "en-US,en;q=0.9",
		"cache-control": "max-age=0",
		"upgrade-insecure-requests": "1",
		"user-agent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/66.0.3359.181 Chrome/66.0.3359.181 Safari/537.36",
	}
		
		f, err := os.OpenFile("./result.txt",os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		ListUrl := "https://www.dy2018.com/html/gndy/dyzz/index_"
		DownBaseUrl := "https://www.dy2018.com/"
		spider := &Spider{ListUrl,DownBaseUrl,header}
		for i := 2; i<3; i++{
			posind := string(i) + ".html"
			spider.ListUrl = spider.ListUrl + posind
			f.WriteString(spider.ListUrl)
			html := spider.get_html_header()
			f.WriteString(html)
		}
}

func main(){
	parse()
}