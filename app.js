const DATA_URL = "json";
let DATA = {};
let TO_DARK = ls("TO_DARK") === "true";

function Main() {
    Checker();
    setInterval(GetJson, 1e4);
    DarkTheme(TO_DARK);
}

window.addEventListener("load", Main);

function GetJson() {
    var url = DATA_URL;
    url = ArrayShuffle([
        "https://gold-1.herokuapp.com/json",
        "https://gold-2.herokuapp.com/json",
        "https://gold-3.herokuapp.com/json",
        "https://gold-4.herokuapp.com/json",
        "https://gold-5.herokuapp.com/json",
        "https://gold-6.herokuapp.com/json",
        "https://gold-7.herokuapp.com/json",
        "https://gold-8.herokuapp.com/json",
        "https://gold-9.herokuapp.com/json",
        "https://gold-10.herokuapp.com/json",
        DATA_URL
    ])[0];

    var xmlhttp = new XMLHttpRequest();
    xmlhttp.onreadystatechange = function () {
        if (this.readyState == 4 && this.status == 200) {
            DATA = JSON.parse(this.responseText);
        }
    };
    xmlhttp.open("GET", url, true);
    xmlhttp.send();
}

function Checker() {
    if (
        typeof DATA["GramStatus"] === "undefined"
    ) {
        setTimeout(Checker, 1e3);
        GetJson();
        Writer();
    } else {
        Writer();
    }
}

function Writer() {
    document.getElementById("content").innerHTML = "";

    let Tarih = DATA.Tarih === undefined ? "00-00-0000 00:00:00" : DATA.Tarih;

    let GramTl = DATA.GramTl === undefined ? "000,00" : DATA.GramTl;
    let GramValue = DATA.GramValue === undefined ? "000" : DATA.GramValue;
    let GramStatus = DATA.GramStatus === undefined ? "%00 XXXXX" : DATA.GramStatus;

    let CeyrekTl = DATA.CeyrekTl === undefined ? "000,00" : DATA.CeyrekTl;
    let CeyrekValue = DATA.CeyrekValue === undefined ? "000" : DATA.CeyrekValue;
    let CeyrekStatus = DATA.CeyrekStatus === undefined ? "%00 XXXXX" : DATA.CeyrekStatus;

    let YarimTl = DATA.YarimTl === undefined ? "000,00" : DATA.YarimTl;
    let YarimValue = DATA.YarimValue === undefined ? "000" : DATA.YarimValue;
    let YarimStatus = DATA.YarimStatus === undefined ? "%00 XXXXX" : DATA.YarimStatus;

    let TamTl = DATA.TamTl === undefined ? "000,00" : DATA.TamTl;
    let TamValue = DATA.TamValue === undefined ? "000" : DATA.TamValue;
    let TamStatus = DATA.TamStatus === undefined ? "%00 XXXXX" : DATA.TamStatus;

    let UsdTlPay = DATA.UsdTlPay === undefined ? "000,00" : DATA.UsdTlPay;
    let UsdTlSell = DATA.UsdTlSell === undefined ? "000,00" : DATA.UsdTlSell;
    let UsdValue = DATA.UsdValue === undefined ? "000" : DATA.UsdValue;
    let UsdStatus = DATA.UsdStatus === undefined ? "%00 XXXXX" : DATA.UsdStatus;

    let EurTlPay = DATA.EurTlPay === undefined ? "000,00" : DATA.EurTlPay;
    let EurTlSell = DATA.EurTlSell === undefined ? "000,00" : DATA.EurTlSell;
    let EurValue = DATA.EurValue === undefined ? "000" : DATA.EurValue;
    let EurStatus = DATA.EurStatus === undefined ? "%00 XXXXX" : DATA.EurStatus;

    let GbpTlPay = DATA.GbpTlPay === undefined ? "000,00" : DATA.GbpTlPay;
    let GbpTlSell = DATA.GbpTlSell === undefined ? "000,00" : DATA.GbpTlSell;
    let GbpValue = DATA.GbpValue === undefined ? "000" : DATA.GbpValue;
    let GbpStatus = DATA.GbpStatus === undefined ? "%00 XXXXX" : DATA.GbpStatus;

    let str = `
        <tr>
            <th>Gram</th>
            <td>${GramTl}TL</td>
            <td>${GramValue}</td>
            <td>${GramStatus}</td>
        </tr>
        <tr>
            <th>Çeyrek</th>
            <td>${CeyrekTl}TL</td>
            <td>${CeyrekValue}</td>
            <td>${CeyrekStatus}</td>
        </tr>
        <tr>
            <th>Yarım</th>
            <td>${YarimTl}TL</td>
            <td>${YarimValue}</td>
            <td>${YarimStatus}</td>
        </tr>
        <tr>
            <th>Tam</th>
            <td>${TamTl}TL</td>
            <td>${TamValue}</td>
            <td>${TamStatus}</td>
        </tr>
        <tr>
            <th>USD</th>
            <td>
                Alış: ${UsdTlPay}TL<br>
                Satış: ${UsdTlSell}TL
            </td>
            <td>${UsdValue}</td>
            <td>${UsdStatus}</td>
        </tr>
        <tr>
            <th>EUR</th>
            <td>
                Alış: ${EurTlPay}TL<br>
                Satış: ${EurTlSell}TL
            </td>
            <td>${EurValue}</td>
            <td>${EurStatus}</td>
        </tr>
        <tr>
            <th>GBP</th>
            <td>
                Alış: ${GbpTlPay}TL<br>
                Satış: ${GbpTlSell}TL
            </td>
            <td>${GbpValue}</td>
            <td>${GbpStatus}</td>
        </tr>
    `;
    document.getElementById("datetime").innerHTML = Tarih;
    document.getElementById("content").innerHTML = str;
}

function DarkTheme(a) {
    TO_DARK = !TO_DARK;
    if (a != null) TO_DARK = a;
    ls("TO_DARK", TO_DARK);
    if (TO_DARK) {
        document.getElementById("datetime").classList.add("bg-dark");
        document.querySelector("body > nav").classList.remove("bg-primary");
        document.querySelector("body > nav").classList.add("bg-dark");
        document.querySelector("body").classList.remove("bg-primary");
        document.querySelector("body").classList.add("bg-dark");
        document.querySelector("body > table").classList.remove("table-hover");
        document.querySelector("body > table").classList.add("table-dark");
    } else {
        document.getElementById("datetime").classList.remove("bg-dark");
        document.querySelector("body > nav").classList.remove("bg-dark");
        document.querySelector("body > nav").classList.add("bg-primary");
        document.querySelector("body").classList.remove("bg-dark");
        document.querySelector("body > table").classList.remove("table-dark");
        document.querySelector("body > table").classList.add("table-hover");
    }
}

function ls(a, b) {
    if (b != null) {
        localStorage.setItem(a, b);
    } else {
        if (localStorage.getItem(a) != null) {
            return localStorage.getItem(a);
        }
    }
}

function ArrayShuffle(a) {
    return a.sort(() => Math.random() - 0.5);
}
