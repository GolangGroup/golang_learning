package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"regexp"
	"time"
	"os"
	"strconv"
)

type Spider struct {
	url string
	header map[string]string
}

func (keyword Spider) get_html_header() string{
	client := &http.Client{}
	fmt.Println(keyword.url)
	req, err := http.NewRequest("GET", keyword.url, nil)
	if err != nil {
		
	}
	for key, value := range keyword.header {
		req.Header.Add(key, value)
	}
	resp, err := client.Do(req)
	if err != nil {
		
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	
	}
	return string(body)
}

func parse() {
	header := map[string]string{
	"Host":"movie.douban.com",
	"Connection":"keep-alive",
	"Cache-Control":"max-age=0",
	"Upgrade-Insecure-Requests":"1",
	"User-Agent":"Mozilla/5.0(Windows NT 6.1; WOW64)AppleWebKit/537.36 (KHTML, link Gecko) Chrome/53.0.2785.143 Safari/537.36",
	"Accept":"test/html,application/xhtml+xml;q=0.9,image/webp,*/*;q=0.8",
	"Referer":"https://movie.douban.com/top250",
	}
	
	f, err := os.OpenFile("./result.txt",os.O_CREATE|os.O_APPEND|os.O_RDWR, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	
	f.WriteString("电影名称"+"\t"+"评分"+"\t"+"评价人数"+"\t"+"\r\n")
	
	for i := 0; i < 2; i++ {
		fmt.Println("正在抓取第"+strconv.Itoa(i)+"页....")
		url := "https://movie.douban.com/top250?start="+strconv.Itoa(i*25)+"&filter="
		spider := &Spider{url, header}
		html := spider.get_html_header()
		
		pattern2 := `<span>(.*?)评价</span>`
		rp2 := regexp.MustCompile(pattern2)
		find_txt2 := rp2.FindAllStringSubmatch(html, -1)
		
		pattern3 := `property="v:average">(.*?)</span>`
		rp3 := regexp.MustCompile(pattern3)
		find_txt3 := rp3.FindAllStringSubmatch(html, -1)		
		
		pattern4 := `img width="100" alt="(.*?)" src=`
		rp4 := regexp.MustCompile(pattern4)
		find_txt4 := rp4.FindAllStringSubmatch(html, -1)
		
		f.WriteString("\xEF\xBB\xBF")
		for i := 0; i < len(find_txt2); i++ {
			fmt.Printf("%s %s %s\n", find_txt4[i][1],find_txt3[i][1],find_txt2[i][1])
			f.WriteString(find_txt4[i][1]+"\t"+find_txt3[i][1]+"\t"+find_txt2[i][1]+"\t"+"\r\n")
		}		
	}
	
	
	
	
	
	
}

func main() {
	t1 := time.Now()
	parse()
	elapsed := time.Since(t1)
	fmt.Println("爬虫结束，总共耗时：", elapsed)
}