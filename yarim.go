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
	fmt.Printf(a)
}

func main() {
	url := "http://www.ikd.sadearge.com/Firma/tablo.php"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	str := string(body)
	r, _ := regexp.Compile(`row12_satis(.*?)>(.*?)<\/td>`)
	arr := r.FindStringSubmatch(str)
	p(strings.TrimSpace(arr[2]))
	bufio.NewReader(os.Stdin).ReadString('\n')
}

// go run yarim.go
// go build yarim.go
