package main

import (
	"io"
	"log"
	"net/http"

	"go.rotbotbox.com/helloweb/web"
)

func helloHandle(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "first go web")
}

func main1() {
	http.HandleFunc("/hello", helloHandle)
	//listen, _ := url.Parse("127.0.0.1:8080")
	err := http.ListenAndServe("0.0.0.0:8311", nil)

	if err != nil {
		log.Fatal("server err: {0}", err.Error())
	}
}

func main() {
	web.Start()
}
