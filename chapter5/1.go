package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	//	"regexp"

	"strconv"
)

type PostIt struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

var k PostIt

func main() {
	TTT := make(chan string)
	RRR := make(chan string)
	var urlN, ell string
	//    numPost := make(chan string)
	//    bodyPost := make(chan string)
	for i := 0; i <= 100; i++ {
		urlN = "https://jsonplaceholder.typicode.com/posts/" + strconv.Itoa(i)
		TTT <- urlN
	}

	for i := 0; i <= 100; i++ {
		ell = <-RRR
		if ell != "none" {
			fmt.Println(ell)
		}
	}

}

func oneposter(TTT chan string, RRR chan string) int {
	ell := ""
	url := <-TTT
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	salo := PostIt{}
	err = json.Unmarshal(body, &salo)
	if err != nil {
		fmt.Println(err)

	}
	ell = "{UserId:" + string(salo.UserId) + " Id:" + string(salo.Id) + " Title:" + salo.Title + " Body:" + salo.Body + "}"
	RRR <- ell //fmt.Println(string(salo.Title))
	//	return strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(string(body), "\n", ""), "  ", " "), "{ ", "{")

	return 0
}
