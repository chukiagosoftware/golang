package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"time"
)

// chapter 2 composite types
type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func processOne(res *http.Response) {
	var data Todo
	err := json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", data)
}

func processMany(res *http.Response) {
	var data []Todo

	err := json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		fmt.Println(res.Body)
		panic(err)
	}
	for i, _ := range data {
		fmt.Printf("%+v\n", data[i])
	}
}

func main() {

	many := flag.Bool("many", false, "get one or many items")
	flag.Parse()
	c := &http.Client{
		Timeout: 30 * time.Second,
	}

	var url string
	var processor func(*http.Response)
	switch *many {
	case true:
		url = "https://jsonplaceholder.typicode.com/todos"
		processor = processMany
	case false:
		url = "https://jsonplaceholder.typicode.com/todos/1"
		processor = processOne
	}
	req, err := http.NewRequestWithContext(context.Background(), "GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("X-My-Client", "Go")
	res, err := c.Do(req)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		panic(fmt.Sprintf("status code error: %d %s", res.StatusCode, res.Status))

	}
	fmt.Println(res.Header.Get("Content-Type"))
	processor(res)
}
