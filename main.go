package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/get", returnFile)
	http.HandleFunc("/post", postFile)
	http.HandleFunc("/del", delFile)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
	fmt.Println("Server listening")
}

func returnFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	if r.Method == "GET" {
		file, err := os.Open("./response/current")
		if err != nil {
			fmt.Println("no file, return 404")
		}
		fmt.Println(file)
		fmt.Println("return file")
	}
	fmt.Println("no file, return 403, method not authorized")
}

func postFile(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Println(r.Body)
	}
	fmt.Println("no file, return 403, method not authorized")
}

func delFile(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DEL" {
		file, err := os.Open("./response/current")
		if err != nil {
			fmt.Println("no file, return 404")
		}
		fmt.Println(file)
		fmt.Println("delete file")
	}
	fmt.Println("no file, return 403, method not authorized")
}
