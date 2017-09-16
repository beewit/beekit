package main

import (
	"compress/flate"
	"compress/gzip"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"

	"golang.org/x/net/html/charset"
)

var (
	links   []string
	db      *sql.DB
	urlHost = "http://www.atobo.com.cn" //抓取网站域名
)

func init() {
	db, _ = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/yl?charset=utf8&timeout=3s")
	//fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&collation=utf8_general_ci&parseTime=True&loc=Asia%%2FShanghai&timeout=30s", username, password, ip, db)
	db.SetMaxOpenConns(0)
	db.SetMaxIdleConns(16)
	db.SetConnMaxLifetime(time.Hour * 2)
	db.Ping()
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var ip = flag.String("ip", "123.146.16.115", "ip address")
	var port = flag.Int("port", 1234, "port")
	//flag.Parse()
	fmt.Printf("ip:%s, port:%d\n", *ip, *port)
	/*urlMap = make(map[string]bool, 1000000)
	fetch("http://www.hao123.com/", 0) 54730
	writeValues("urls.dat")*/
	/*fmt.Println("开始抓人...")
	for i := 1; i <= 100; i++ {
		getLinks(fmt.Sprintf("http://www.atobo.com.cn/GongShang/s-p4-s873-y%d/", i), "#filterArea a")
		time.Sleep(time.Duration(rand.Intn(6)) * time.Second)
	}
	count := len(links)
	for i := 0; i < count; i++ {
		customer := getBoss(links[i])
		//customer["district"], customer["firm"], customer["fullname"], customer["tel"], customer["mobile"]
		if customer != nil && customer["fullname"] != "" && (customer["tel"] != "" || customer["mobile"] != "") {
			err := Add(customer)
			CheckError(err)
			time.Sleep(time.Duration(rand.Intn(6)) * time.Second)
		}
	}
	fmt.Println("抓人完成...")*/

	getLinks("http://www.atobo.com.cn/GongShang/s-p4-s873/", "a")

	//getBoss("http://www.atobo.com.cn/GongShang/696/E1B5348AA3474A37802513C1905D2553.html")
	//getBossTemp("http://www.atobo.com.cn/GongShang/696/E1B5348AA3474A37802513C1905D2553.html")
	resp, err := Get("http://www.baidu.com")
	CheckErr(err)
	defer resp.Body.Close()
	/*byt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(byt))*/

	/*body, _ := httpGet("http://shop.yqfa.cn")
	_body, err = charset.NewReader(resp.Body, "text/html; charset=gb2312")
	fmt.Println(body)*/
}

func batchScrape(urls []string) {

}

func linkScrape() {
	doc, err := goquery.NewDocument("http://jonathanmh.com")
	if err != nil {
		log.Fatal(err)
	}

	// use CSS selector found with the browser inspector
	// for each, use index and item
	doc.Find("body a").Each(func(index int, item *goquery.Selection) {
		linkTag := item
		link, _ := linkTag.Attr("href")
		linkText := linkTag.Text()
		fmt.Printf("Link #%d: '%s' - '%s'\n", index, linkText, link)
	})
}

func GetBetterBody(resp *http.Response, contentType string) (io.Reader, error) {
	if strings.Contains(resp.Header.Get("Content-Type"), "utf-8") {
		return resp.Body, nil
	} else {
		return charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
	}
}

func Scrape(url string) {
	fmt.Println("i am here:" + url)
	doc, err := goquery.NewDocument("http://www.oschina.net")
	if err != nil {
		fmt.Println(err)
		return
	}

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		//band := s.Find("span").Text()
		title := strings.TrimSpace(s.Text())
		href, _ := s.Attr("href")
		fmt.Printf("%s - %s\n", title, href)
	})
}

func getBoss(url string) map[string]string {
	resp, err := Get(url)
	CheckErr(err)
	defer resp.Body.Close()

	fmt.Printf("Method=%s, URL=%s, StatusCode=%d, Status=%s\n", resp.Request.Method, resp.Request.URL, resp.StatusCode, resp.Status)
	fmt.Printf("%v", resp.Header.Get("Content-Type"))

	body, err := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))

	CheckErr(err)
	//doc, _ := goquery.NewDocumentFromReader(resp.Body)
	doc, _ := goquery.NewDocumentFromReader(body)

	boss := make(map[string]string)
	boss["title"] = doc.Find("title").Text()
	boss["district"] = doc.Find("div.cur_post").Text()
	boss["firm"] = doc.Find("tbody > tr:nth-child(1) > td.rowcontact.rowcontact_long").Text()
	boss["fullname"] = doc.Find("tbody > tr:nth-child(3) > td:nth-child(2) > a").Text()
	boss["tel"] = doc.Find("li.llist > div > ul > li > div > ul > li.tuo-tel > span").Text()
	boss["mobile"] = doc.Find("li.llist > div > ul > li > div > ul > li.tuo-mobile > span").Text()
	fmt.Printf("\n%s\n%v\n", url, boss)
	return boss
}

func getLinks(url, css string) {
	//fmt.Printf("开始抓URL:%s\n", url)
	resp, err := Get(url)
	CheckErr(err)
	defer resp.Body.Close()

	body, err := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
	doc, _ := goquery.NewDocumentFromReader(body)
	//fmt.Println(doc.Html())
	fmt.Println(doc.Find("title").Text())

	doc.Find(css).Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Html())
		//title := s.Text()
		href, _ := s.Attr("href")
		fmt.Printf("http://www.atobo.com.cn%s\n", href)
		//fmt.Printf("拿到待抓URL[%s:http://www.atobo.com.cn%s]\n", title, href)
		/*if strings.Contains(href, ".html") && !strings.HasPrefix(href, "http") {
			fmt.Printf("拿到待抓URL[%s: http://www.atobo.com.cn%s]\n", title, href)
			links = append(links, fmt.Sprintf("%s%s", urlHost, href))
		}*/
	})
	//fmt.Printf("抓完成URL:%s\n", url)
}

/*func getLinks() {
	resp, err := Get("http://www.atobo.com.cn/GongShang/s-p4-s873/")
	CheckError(err)
	defer resp.Body.Close()

	doc, _ := goquery.NewDocumentFromReader(resp.Body)

	conver, _ := iconv.Open("utf-8", "gbk")
	defer conver.Close()

	doc.Find("#filterArea > ul > li a").Each(func(i int, s *goquery.Selection) {
		title := conver.ConvString(s.Text())
		href, _ := s.Attr("href")
		fmt.Printf("%s: http://www.atobo.com.cn%s\n", title, href)
		links = append(links, fmt.Sprintf("http://www.atobo.com.cn%s", href))
	})
}*/

func Get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	//req.Header.Set("User-Agent", useragent.GetRandomUserAgent())
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	//req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.8")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cache-Control", "max-age=0")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Printf("%v\n", resp.Header.Get("Content-Type"))
	return resp, nil
}

func httpGet(url string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	//req.Header.Set("User-Agent", useragent.GetRandomUserAgent())
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.8")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cache-Control", "max-age=0")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	contentEncoding := strings.Trim(strings.ToLower(resp.Header.Get("Content-Encoding")), " ")
	var bytes []byte
	if contentEncoding == "gzip" {
		x, err := gzip.NewReader(resp.Body)
		if err != nil {
			return "", err
		}
		bytes, err = ioutil.ReadAll(x)
		if err != nil {
			return "", err
		}
	} else if contentEncoding == "deflate" {
		x := flate.NewReader(resp.Body)
		bytes, err = ioutil.ReadAll(x)
		if err != nil {
			return "", err
		}
	} else {
		bytes, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
	}

	//srcStr := string(bytes)
	/*pageCodes := pageCodeReg.FindStringSubmatch(srcStr)
	if len(pageCodes) >= 2 {
		curCode := strings.ToLower(pageCodes[1])
		if curCode == "gb2312" || curCode == "gbk" {
			decoder := mahonia.NewDecoder("gbk")
			srcStr = decoder.ConvertString(srcStr)
		}
	}*/
	return string(bytes), nil
}

/*func DecodeToGBK(text string) (string, error) {

	dst := make([]byte, len(text)*2)
	tr := simplifiedchinese.GB18030.NewDecoder()
	nDst, _, err := tr.Transform(dst, []byte(text), true)
	if err != nil {
		return text, err
	}

	return string(dst[:nDst]), nil
}*/

//添加一个新用户
func Add(customer map[string]string) error {
	if customer == nil || len(customer) == 0 {
		return errors.New("插入的内容不能为空")
	}
	stmt, err := db.Prepare("REPLACE INTO customer(district, firm, fullname, tel, mobile) VALUES(?, ?, ?, ?, ?)")
	// 一定要关闭，很重要
	defer stmt.Close()

	if err != nil {
		log.Println(err)
		return err
	}
	_, err = stmt.Exec(customer["district"], customer["firm"], customer["fullname"], customer["tel"], customer["mobile"])

	return err
}

func updateUser() error {
	sql := "update test.table1 set _create_time=now() where id=?"
	result, err := db.Exec(sql, 7988161482)
	if err != nil {
		fmt.Println("exec error ", err.Error())
		return err
	}
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())
	return nil
}

//查询demo
func query() {
	rows, err := db.Query("SELECT * FROM user")
	CheckErr(err)

	//字典类型,构造scanArgs、values两个数组，scanArgs的每个值指向values相应值的地址
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		//将行数据保存到record字典
		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		fmt.Println(record)
	}
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

// -----------------------------
func IsUrl(str string) bool {
	if strings.HasPrefix(str, "#") || strings.HasPrefix(str, "//") || strings.HasSuffix(str, ".exe") || strings.HasSuffix(str, ":void(0);") {
		return false
	} else if strings.HasPrefix(str, "{") && strings.HasSuffix(str, "}") {
		return false
	} else if strings.EqualFold(str, "javascript:;") {
		return false
	} else {
		return true
	}
	return true
}

func SamePathUrl(preUrl string, url string, mark int) (newUrl string) {
	last := strings.LastIndex(preUrl, "/")
	if last == 6 {
		newUrl = preUrl + url
	} else {
		if mark == 1 {
			newUrl = preUrl[:last] + url
		} else {
			newPreUrl := preUrl[:last]
			newLast := strings.LastIndex(newPreUrl, "/")
			newUrl = newPreUrl[:newLast] + url
		}
	}
	return newUrl
}

// --------------------------judgeUrl-------------------------------------
var urlMap map[string]bool //防止重复链接进入死循环，不过获取链接太多可能会内存溢出

func fetch(url string, count int) {
	if count > 1 { //设定爬取深度是1页
		return
	}

	body, err := goquery.NewDocument(url)
	if err != nil {
		return
	}
	body.Find("a").Each(func(i int, aa *goquery.Selection) {
		href, IsExist := aa.Attr("href")
		if IsExist == true {
			href = strings.TrimSpace(href)
			if len(href) > 2 && IsUrl(href) {
				if _, ok := urlMap[href]; ok == false {
					fmt.Println("之前的url：", href)
					if strings.HasPrefix(href, "/") || strings.HasPrefix(href, "./") {
						href = SamePathUrl(url, href, 1)
					} else if strings.HasPrefix(href, "../") {
						href = SamePathUrl(url, href, 2)
					}
					fmt.Println("修改之后的url：", href)
					urlMap[href] = true
					fetch(href, count+1)
				}
			}
		}
	})
}

func writeValues(outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Printf("创建%s文件失败！", outfile)
		return err
	}
	defer file.Close()
	for k, _ := range urlMap {
		file.WriteString(k + "\n")
	}
	return nil
}

// -----------------------------
