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
	return str + ""
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

func generate() string {
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

	out := ""
	out += "\n  Tarih: " + pars(str, `tarih(.*?)>(.*?)<\/span>`, 2, "Son Güncellenme Tarihi : ")
	out += "\n   Gram: " + pars(str, `row6_satis(.*?)>(.*?)<\/td>`, 2, "") + " " + a + " " + b + " Oy"
	out += "\n Çeyrek: " + pars(str, `row11_satis(.*?)>(.*?)<\/td>`, 2, "") + " " + c + " " + d + " Oy"
	out += "\n  Yarım: " + pars(str, `row12_satis(.*?)>(.*?)<\/td>`, 2, "") + " " + e + " " + f + " Oy"
	return out
}

func main() {
	p(generate())
	http.HandleFunc("/", func (w http.ResponseWriter, req *http.Request) {
		out := generate()
		p(out)
		fmt.Fprintln(w, out)
	})

	http.ListenAndServe(":6010", nil)

	paus()
}

// go run gold.go
// go build gold.go
