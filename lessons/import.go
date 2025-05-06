package main

import (
	"fmt"          // default way for import library
	foo "net/http" // import library as custom name
) // this way for import multiple library

func import_lesson() {

	fmt.Println("hello go standard library")
	resp, err := foo.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer resp.Body.Close()
	
	fmt.Println("http response ststus: ", resp.Status)
}
