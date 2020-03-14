package main

import (
	"bufio"
	"fmt"
	"github.com/grokify/html-strip-tags-go" // => strip
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
	_, _ = bufio.NewReader(os.Stdin).ReadString('\n')
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
	str := get("https://www.sabah.com.tr/izmir-namaz-vakitleri")
	out := strip.StripTags(pars(str, `(?ms)(.*?)<div class="vakitler boxShadowSet">(.*?)<\/div>(.*?)`, 2, " "))
	out = regexp.MustCompile(`\r?\n\r?\n`).ReplaceAllString(strings.TrimSpace(out), "\n")
	return out
}

func main() {
	p(generate())
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		out := generate()
		p(out)
		w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		fmt.Fprintln(w, "<meta name=\"viewport\" content=\"width=device-width, initial-scale=1\"><style>*{background: #000; color: #fff;}</style><pre style=\"font-size: 1.5em;\">"+out+"</pre>")
	})

	_ = http.ListenAndServe(":8035", nil)

	paus()
}

// go build iznamaz.go && ./iznamaz*
