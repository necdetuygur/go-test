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
	fmt.Printf(a + "\n")
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
	str := get("http://www.ikd.sadearge.com/Firma/tablo.php")
	p("Tarih  : " + pars(str, `tarih(.*?)>(.*?)<\/span>`, 2, "Son Güncellenme Tarihi : "))
	p("Gram   : " + pars(str, `row6_satis(.*?)>(.*?)<\/td>`, 2, ""))
	p("Çeyrek : " + pars(str, `row11_satis(.*?)>(.*?)<\/td>`, 2, ""))
	p("Yarım  : " + pars(str, `row12_satis(.*?)>(.*?)<\/td>`, 2, ""))
	paus()
}

// go run ikd.go
// go build ikd.go
