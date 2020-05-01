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
	newLineStr := "\n         "
	str := get("https://finanswebde.com/altin/gram-altin")
	gram_status := pars(str, `<div class="col-md-6"><span class="detail-change(.*?)>(.*?)<!--(.*?)<\/span>(.*?)<span(.*?)class=\"detail-title-sm\">(.*?)<span>(.*?)<\/span>(.*?)<\/span>(.*?)`, 2, "")
	gram_value := pars(str, `<div class="col-md-6"><span class="detail-change(.*?)>(.*?)<!--(.*?)<\/span>(.*?)<span(.*?)class=\"detail-title-sm\">(.*?)<span>(.*?)<\/span>(.*?)<\/span>(.*?)`, 7, "")

	str = get("https://finanswebde.com/altin/ceyrek-altin")
	ceyrek_status := pars(str, `<div class="col-md-6"><span class="detail-change(.*?)>(.*?)<!--(.*?)<\/span>(.*?)<span(.*?)class=\"detail-title-sm\">(.*?)<span>(.*?)<\/span>(.*?)<\/span>(.*?)`, 2, "")
	ceyrek_value := pars(str, `<div class="col-md-6"><span class="detail-change(.*?)>(.*?)<!--(.*?)<\/span>(.*?)<span(.*?)class=\"detail-title-sm\">(.*?)<span>(.*?)<\/span>(.*?)<\/span>(.*?)`, 7, "")

	str = get("https://finanswebde.com/altin/yarim-altin")
	yarim_status := pars(str, `<div class="col-md-6"><span class="detail-change(.*?)>(.*?)<!--(.*?)<\/span>(.*?)<span(.*?)class=\"detail-title-sm\">(.*?)<span>(.*?)<\/span>(.*?)<\/span>(.*?)`, 2, "")
	yarim_value := pars(str, `<div class="col-md-6"><span class="detail-change(.*?)>(.*?)<!--(.*?)<\/span>(.*?)<span(.*?)class=\"detail-title-sm\">(.*?)<span>(.*?)<\/span>(.*?)<\/span>(.*?)`, 7, "")

	str = get("https://finanswebde.com/altin/tam-altin")
	tam_status := pars(str, `<div class="col-md-6"><span class="detail-change(.*?)>(.*?)<!--(.*?)<\/span>(.*?)<span(.*?)class=\"detail-title-sm\">(.*?)<span>(.*?)<\/span>(.*?)<\/span>(.*?)`, 2, "")
	tam_value := pars(str, `<div class="col-md-6"><span class="detail-change(.*?)>(.*?)<!--(.*?)<\/span>(.*?)<span(.*?)class=\"detail-title-sm\">(.*?)<span>(.*?)<\/span>(.*?)<\/span>(.*?)`, 7, "")

	str = get("https://finanswebde.com/doviz/USD")
	usd_tl_pay := pars(str, `<div class="d-flex align-self-center justify-content-between my-2 py-1"><div class="detail-main-title"><span class="detail-title-sm">ALIŞ <\/span> <span class="detail-title-md "><nobr>(.*?)<\/nobr><\/span><\/div><div class="detail-main-title"><span class="detail-title-sm">SATIŞ <\/span> <span class="detail-title-md "><nobr>(.*?)<\/nobr><\/span><\/div>(.*?)<\/div>`, 1, "")
	usd_tl_sell := pars(str, `<div class="d-flex align-self-center justify-content-between my-2 py-1"><div class="detail-main-title"><span class="detail-title-sm">ALIŞ <\/span> <span class="detail-title-md "><nobr>(.*?)<\/nobr><\/span><\/div><div class="detail-main-title"><span class="detail-title-sm">SATIŞ <\/span> <span class="detail-title-md "><nobr>(.*?)<\/nobr><\/span><\/div>(.*?)<\/div>`, 2, "")
	usd_status := pars(str, `<div class="col-md-6"><span class="detail-change(.*?)>(.*?)<!--(.*?)<\/span>(.*?)<span(.*?)class=\"detail-title-sm\">(.*?)<span>(.*?)<\/span>(.*?)<\/span>(.*?)`, 2, "")
	usd_value := pars(str, `<div class="col-md-6"><span class="detail-change(.*?)>(.*?)<!--(.*?)<\/span>(.*?)<span(.*?)class=\"detail-title-sm\">(.*?)<span>(.*?)<\/span>(.*?)<\/span>(.*?)`, 7, "")

	str = get("https://finanswebde.com/doviz/EUR")
	eur_tl_pay := pars(str, `<div class="d-flex align-self-center justify-content-between my-2 py-1"><div class="detail-main-title"><span class="detail-title-sm">ALIŞ <\/span> <span class="detail-title-md "><nobr>(.*?)<\/nobr><\/span><\/div><div class="detail-main-title"><span class="detail-title-sm">SATIŞ <\/span> <span class="detail-title-md "><nobr>(.*?)<\/nobr><\/span><\/div>(.*?)<\/div>`, 1, "")
	eur_tl_sell := pars(str, `<div class="d-flex align-self-center justify-content-between my-2 py-1"><div class="detail-main-title"><span class="detail-title-sm">ALIŞ <\/span> <span class="detail-title-md "><nobr>(.*?)<\/nobr><\/span><\/div><div class="detail-main-title"><span class="detail-title-sm">SATIŞ <\/span> <span class="detail-title-md "><nobr>(.*?)<\/nobr><\/span><\/div>(.*?)<\/div>`, 2, "")
	eur_status := pars(str, `<div class="col-md-6"><span class="detail-change(.*?)>(.*?)<!--(.*?)<\/span>(.*?)<span(.*?)class=\"detail-title-sm\">(.*?)<span>(.*?)<\/span>(.*?)<\/span>(.*?)`, 2, "")
	eur_value := pars(str, `<div class="col-md-6"><span class="detail-change(.*?)>(.*?)<!--(.*?)<\/span>(.*?)<span(.*?)class=\"detail-title-sm\">(.*?)<span>(.*?)<\/span>(.*?)<\/span>(.*?)`, 7, "")

	str = get("https://finanswebde.com/doviz/GBP")
	gbp_tl_pay := pars(str, `<div class="d-flex align-self-center justify-content-between my-2 py-1"><div class="detail-main-title"><span class="detail-title-sm">ALIŞ <\/span> <span class="detail-title-md "><nobr>(.*?)<\/nobr><\/span><\/div><div class="detail-main-title"><span class="detail-title-sm">SATIŞ <\/span> <span class="detail-title-md "><nobr>(.*?)<\/nobr><\/span><\/div>(.*?)<\/div>`, 1, "")
	gbp_tl_sell := pars(str, `<div class="d-flex align-self-center justify-content-between my-2 py-1"><div class="detail-main-title"><span class="detail-title-sm">ALIŞ <\/span> <span class="detail-title-md "><nobr>(.*?)<\/nobr><\/span><\/div><div class="detail-main-title"><span class="detail-title-sm">SATIŞ <\/span> <span class="detail-title-md "><nobr>(.*?)<\/nobr><\/span><\/div>(.*?)<\/div>`, 2, "")
	gbp_status := pars(str, `<div class="col-md-6"><span class="detail-change(.*?)>(.*?)<!--(.*?)<\/span>(.*?)<span(.*?)class=\"detail-title-sm\">(.*?)<span>(.*?)<\/span>(.*?)<\/span>(.*?)`, 2, "")
	gbp_value := pars(str, `<div class="col-md-6"><span class="detail-change(.*?)>(.*?)<!--(.*?)<\/span>(.*?)<span(.*?)class=\"detail-title-sm\">(.*?)<span>(.*?)<\/span>(.*?)<\/span>(.*?)`, 7, "")

	str = get("http://www.ikd.sadearge.com/Firma/tablo.php")

	tarih := pars(str, `tarih(.*?)>(.*?)<\/span>`, 2, "Son Güncellenme Tarihi : ")

	gram :=
		pars(str, `row6_satis(.*?)>(.*?)<\/td>`, 2, "") + "TL" +
			newLineStr + gram_value + " Oyla" +
			newLineStr + gram_status

	ceyrek :=
		pars(str, `row11_satis(.*?)>(.*?)<\/td>`, 2, "") + "TL" +
			newLineStr + ceyrek_value + " Oyla" +
			newLineStr + ceyrek_status

	yarim :=
		pars(str, `row12_satis(.*?)>(.*?)<\/td>`, 2, "") + "TL" +
			newLineStr + yarim_value + " Oyla" +
			newLineStr + yarim_status

	tam :=
		pars(str, `row13_satis(.*?)>(.*?)<\/td>`, 2, "") + "TL" +
			newLineStr + tam_value + " Oyla" +
			newLineStr + tam_status

	usd :=
		"Alış: " + usd_tl_pay + "TL" +
		newLineStr + "Satış: " + usd_tl_sell + "TL" +
			newLineStr + usd_value + " Oyla" +
			newLineStr + usd_status

	eur :=
		"Alış: " + eur_tl_pay + "TL" +
		newLineStr + "Satış: " + eur_tl_sell + "TL" +
			newLineStr + eur_value + " Oyla" +
			newLineStr + eur_status

	gbp :=
		"Alış: " + gbp_tl_pay + "TL" +
		newLineStr + "Satış: " + gbp_tl_sell + "TL" +
			newLineStr + gbp_value + " Oyla" +
			newLineStr + gbp_status

	out := ""
	out += "\n\n  Tarih: " + tarih
	out += "\n\n   Gram: " + gram
	out += "\n\n Çeyrek: " + ceyrek
	out += "\n\n  Yarım: " + yarim
	out += "\n\n    Tam: " + tam
	out += "\n\n    USD: " + usd
	out += "\n\n    EUR: " + eur
	out += "\n\n    GBP: " + gbp
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

	http.ListenAndServe(":6010", nil)

	paus()
}

// go build gold.go && ./gold*
