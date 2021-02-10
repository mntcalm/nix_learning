package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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
	PPC := make(chan int)
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
		go readpost(PPP, PPC)
		PPP <- pst[i].Id
	}
	for i := 0; i < len(pst); i++ {
		_ = <-PPC
	}
}

func readpost(PPP, PPC chan int) int {
	Ccompost := make(chan CommIt)
	CCC := make(chan int)
	postnom := <-PPP
	//fmt.Println(postnom)

	cmt := []CommIt{}
	url := "https://jsonplaceholder.typicode.com/comments?postId=" + strconv.Itoa(postnom)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(body, &cmt)
	for i := 0; i < len(cmt); i++ {

		go readcomm(Ccompost, CCC)
		Ccompost <- cmt[i]

	}
	for i := 0; i < len(cmt); i++ {
		_ = <-CCC
	}
	PPC <- 0
	return 0

}
func readcomm(Ccompost chan CommIt, CCC chan int) int {
	compost := <-Ccompost

	text := "ID:" + strconv.Itoa(compost.Id) + " POSTID:" + strconv.Itoa(compost.PostId) + " NAME:" + compost.Name + " EMAIL:" + compost.Email + " BODY:" + compost.Body
	fmt.Println(text, "\n")
	CCC <- 1
	return 0
}
