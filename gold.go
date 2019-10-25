package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func p(a string) {
	fmt.Fprintln(os.Stdout, a)
}

func get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	str := string(body)
	return str
}

func paus(){
	bufio.NewReader(os.Stdin).ReadString('\n')
}

func pars(str string, rgx string, key int, clr string) string {
	r, _ := regexp.Compile(rgx)
	arr := r.FindStringSubmatch(str)
	pri := strings.ReplaceAll(strings.TrimSpace(arr[key]), clr, "")
	return pri
}

func main() {
	str := get("https://finanswebde.com/altin/gram-altin")
	a := pars(str, `<div class="col-md-6"><span class="detail-change(.*?)>(.*?)<!--(.*?)<\/span>(.*?)<span(.*?)class=\"detail-title-sm\">(.*?)<span>(.*?)<\/span>(.*?)<\/span>(.*?)`, 2, "")
	b := pars(str, `<div class="col-md-6"><span class="detail-change(.*?)>(.*?)<!--(.*?)<\/span>(.*?)<span(.*?)class=\"detail-title-sm\">(.*?)<span>(.*?)<\/span>(.*?)<\/span>(.*?)`, 7, "")

	str = get("https://finanswebde.com/altin/ceyrek-altin")
	c := pars(str, `<div class="col-md-6"><span class="detail-change(.*?)>(.*?)<!--(.*?)<\/span>(.*?)<span(.*?)class=\"detail-title-sm\">(.*?)<span>(.*?)<\/span>(.*?)<\/span>(.*?)`, 2, "")
	d := pars(str, `<div class="col-md-6"><span class="detail-change(.*?)>(.*?)<!--(.*?)<\/span>(.*?)<span(.*?)class=\"detail-title-sm\">(.*?)<span>(.*?)<\/span>(.*?)<\/span>(.*?)`, 7, "")

	str = get("https://finanswebde.com/altin/yarim-altin")
	e := pars(str, `<div class="col-md-6"><span class="detail-change(.*?)>(.*?)<!--(.*?)<\/span>(.*?)<span(.*?)class=\"detail-title-sm\">(.*?)<span>(.*?)<\/span>(.*?)<\/span>(.*?)`, 2, "")
	f := pars(str, `<div class="col-md-6"><span class="detail-change(.*?)>(.*?)<!--(.*?)<\/span>(.*?)<span(.*?)class=\"detail-title-sm\">(.*?)<span>(.*?)<\/span>(.*?)<\/span>(.*?)`, 7, "")

	str = get("http://www.ikd.sadearge.com/Firma/tablo.php")
	p("  Tarih: " + pars(str, `tarih(.*?)>(.*?)<\/span>`, 2, "Son Güncellenme Tarihi : "))
	p("   Gram: " + pars(str, `row6_satis(.*?)>(.*?)<\/td>`, 2, "") + " " + a + " " + b + " Oy")
	p(" Çeyrek: " + pars(str, `row11_satis(.*?)>(.*?)<\/td>`, 2, "") + " " + c + " " + d + " Oy")
	p("  Yarım: " + pars(str, `row12_satis(.*?)>(.*?)<\/td>`, 2, "") + " " + e + " " + f + " Oy")

	paus()
}

// go run gold.go
// go build gold.go
