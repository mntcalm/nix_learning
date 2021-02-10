package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	//      "os"
)

type CommIt struct {
	PostId int    `json:"postId"`
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

type PostIt struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	PPP := make(chan int)
	CCC := make(chan int)
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts?userId=7")
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	pst := []PostIt{}
	err = json.Unmarshal(body, &pst)
	for i := 0; i < len(pst); i++ {
		go readpost(PPP)
		PPP <- pst[i].Id
	}

}

func readpost(PPP chan int) int {
	postnom := <-PPP
	fmt.Println(postnom)
	
	cmt := []CommIt{}
    url := "https://jsonplaceholder.typicode.com/comments?postId=" + strconv.Itoa(postnom)
	cmn, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
		err = json.Unmarshal(body, &cmt)
		for i := 0; i < len(cmt); i++ {
			go readcomm(CCC)
			CCC <- cmt[i].Id
		}
	
	

	return 0
}

func readcomm(CCC chan int) int {
 commnom := <-CCC
 url=

}
