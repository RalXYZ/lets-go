package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func output(response http.ResponseWriter, content *content) {
	fmt.Fprintf(response, "%v %v %v\n", content.uid, content.id, content.age)
}

func main() {
	http.HandleFunc("/create", func(response http.ResponseWriter, request *http.Request) {
		temp, _ := url.ParseQuery(request.URL.RawQuery)
		id := temp.Get("id")
		age, _ := strconv.Atoi(temp.Get("age")) // FIXME: Add error handler
		content := dbCreate(&bar{id, age})
		output(response, content)
		/*
			for key, value := range temp {
				fmt.Println(key, value)
			}
		*/
	})

	http.HandleFunc("/retrieve", func(response http.ResponseWriter, request *http.Request) {
		temp, _ := url.ParseQuery(request.URL.RawQuery)
		uid, _ := strconv.ParseInt(temp.Get("uid"), 10, 64) // FIXME: Add error handler
		content := dbRetrieve(uid)
		output(response, content)
	})

	http.HandleFunc("/update", func(response http.ResponseWriter, request *http.Request) {
		temp, _ := url.ParseQuery(request.URL.RawQuery)
		uid, _ := strconv.ParseInt(temp.Get("uid"), 10, 64) // FIXME: Add error handler
		id := temp.Get("id")
		age, _ := strconv.Atoi(temp.Get("age")) // FIXME: Add error handler
		content := dbUpdate(&content{uid, id, age})
		output(response, content)
	})

	http.HandleFunc("/delete", func(response http.ResponseWriter, request *http.Request) {
		temp, _ := url.ParseQuery(request.URL.RawQuery)
		uid, _ := strconv.ParseInt(temp.Get("uid"), 10, 64) // FIXME: Add error handler
		dbDelete(uid)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
