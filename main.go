package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handleFunc)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
	fmt.Println("Server listening")
}

func handleFunc(w http.ResponseWriter, r *http.Request) {
	file := "./response/res.txt"
	switch method := r.Method; method {
	case "GET":
		data, err := ioutil.ReadFile(file)
		if err != nil {
			w.WriteHeader(404)
			return
		}
		w.WriteHeader(200)
		w.Write([]byte(data))
		return
	case "POST":
		err := ioutil.WriteFile(file, convertReq(r), 0777)
		if err != nil {
			w.WriteHeader(500)
		}
		return
	case "DELETE":
		err := os.Remove(file)
		if err != nil {
			w.WriteHeader(404)
			return
		}
		w.WriteHeader(200)
		return
	default:
		w.WriteHeader(404)
	}
}

func convertReq(r *http.Request) []byte {
	h := convertHeader(r.Header)
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	// maybe add a \n between header and request ?
	return append(h, b...)
}

func convertHeader(h http.Header) []byte {
	res := []byte{}
	for k, v := range h {
		res = append(res, []byte(k)...)
		res = append(res, []byte(arrToString(v))...)
	}
	return res
}

func arrToString(a []string) string {
	res := ""
	for _, v := range a {
		res += v
	}
	return res
}
