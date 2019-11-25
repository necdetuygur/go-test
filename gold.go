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

func paus() {
	bufio.NewReader(os.Stdin).ReadString('\n')
}

func pars(str string, rgx string, key int, clr string) string {
	r, _ := regexp.Compile(rgx)
	arr := r.FindStringSubmatch(str)
	if len(arr) == 0 {
		return "0"
	}
	pri := strings.ReplaceAll(strings.TrimSpace(arr[key]), clr, "")
	return pri
}

func generate() string {
	str := get("https://finanswebde.com/altin/gram-altin")
	gram_status := pars(str, `<div class="col-md-6"><span class="detail-change(.*?)>(.*?)<!--(.*?)<\/span>(.*?)<span(.*?)class=\"detail-title-sm\">(.*?)<span>(.*?)<\/span>(.*?)<\/span>(.*?)`, 2, "")
	gram_value := pars(str, `<div class="col-md-6"><span class="detail-change(.*?)>(.*?)<!--(.*?)<\/span>(.*?)<span(.*?)class=\"detail-title-sm\">(.*?)<span>(.*?)<\/span>(.*?)<\/span>(.*?)`, 7, "")

	str = get("https://finanswebde.com/altin/ceyrek-altin")
	ceyrek_status := pars(str, `<div class="col-md-6"><span class="detail-change(.*?)>(.*?)<!--(.*?)<\/span>(.*?)<span(.*?)class=\"detail-title-sm\">(.*?)<span>(.*?)<\/span>(.*?)<\/span>(.*?)`, 2, "")
	ceyrek_value := pars(str, `<div class="col-md-6"><span class="detail-change(.*?)>(.*?)<!--(.*?)<\/span>(.*?)<span(.*?)class=\"detail-title-sm\">(.*?)<span>(.*?)<\/span>(.*?)<\/span>(.*?)`, 7, "")

	str = get("https://finanswebde.com/altin/yarim-altin")
	yarim_status := pars(str, `<div class="col-md-6"><span class="detail-change(.*?)>(.*?)<!--(.*?)<\/span>(.*?)<span(.*?)class=\"detail-title-sm\">(.*?)<span>(.*?)<\/span>(.*?)<\/span>(.*?)`, 2, "")
	yarim_value := pars(str, `<div class="col-md-6"><span class="detail-change(.*?)>(.*?)<!--(.*?)<\/span>(.*?)<span(.*?)class=\"detail-title-sm\">(.*?)<span>(.*?)<\/span>(.*?)<\/span>(.*?)`, 7, "")

	str = get("http://www.ikd.sadearge.com/Firma/tablo.php")

	out := ""
	out += "\n  Tarih: " + pars(str, `tarih(.*?)>(.*?)<\/span>`, 2, "Son Güncellenme Tarihi : ")
	out += "\n   Gram: " + pars(str, `row6_satis(.*?)>(.*?)<\/td>`, 2, "") + " " + gram_status + " " + gram_value + " Oy"
	out += "\n Çeyrek: " + pars(str, `row11_satis(.*?)>(.*?)<\/td>`, 2, "") + " " + ceyrek_status + " " + ceyrek_value + " Oy"
	out += "\n  Yarım: " + pars(str, `row12_satis(.*?)>(.*?)<\/td>`, 2, "") + " " + yarim_status + " " + yarim_value + " Oy"
	return out
}

func main() {
	p(generate())
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		out := generate()
		p(out)
    w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		fmt.Fprintln(w, "<meta name=\"viewport\" content=\"width=device-width, initial-scale=1\"><pre>" + out + "</pre>")
	})

	http.ListenAndServe(":6010", nil)

	paus()
}

// go run gold.go
// go build gold.go
