package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

func writeFile(filename string, content string) {
	mydata := []byte(content)
	err := ioutil.WriteFile(filename, mydata, 0777)
	if err != nil {
		fmt.Println(err)
	}
}

func getDayName() string {
	return ([7]string{"PZR", "PTS", "SLI", "CRS", "PRS", "CMA", "CTS"})[time.Now().Weekday()]
}

func main() {
	t := time.Now()
	time := fmt.Sprintf("%d%02d%02d_%s", t.Year(), t.Month(), t.Day(), getDayName())
	file := time + ".txt"
	writeFile(file, "")
	fmt.Println(file + " generated.")
}
