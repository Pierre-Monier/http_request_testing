package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleFunc)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

	fmt.Println("Server listening")
}

func handleFunc(w http.ResponseWriter, r *http.Request) {
	switch method := r.Method; method {
	case "GET":
		fmt.Println("return file")
		data, err := ioutil.ReadFile("test.txt")
		if err != nil {
			w.WriteHeader(404)
			return
		}
		w.WriteHeader(200)
		w.Write([]byte(data))
		return
	case "POST":
		fmt.Println("Ecrire fichier")
		ioutil.WriteFile("response/current", convertReq(r))
	case "DELETE":
	default:
	}

	fmt.Println("no file, return 403, method not authorized")
}

func convertReq(r *http.Request) string {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	return b
	fmt.Println(data)
}

// func returnFile(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println(r.Method)
// 	if r.Method == "GET" {

// 		if err != nil {
// 			fmt.Println("no file, return 404")
// 		}
// 		fmt.Println(file)
// 		fmt.Println("return file")
// 	}
// 	fmt.Println("no file, return 403, method not authorized")
// }

// func postFile(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "POST" {
// 		fmt.Println(r.Body)
// 	}
// 	fmt.Println("no file, return 403, method not authorized")
// }

// func delFile(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "DEL" {
// 		file, err := os.Open("./response/current")
// 		if err != nil {
// 			fmt.Println("no file, return 404")
// 		}
// 		fmt.Println(file)
// 		fmt.Println("delete file")
// 	}
// 	fmt.Println("no file, return 403, method not authorized")
// }
