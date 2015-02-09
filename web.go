package main

import (
	"code.google.com/p/goauth2/oauth"
	"encoding/json"
	"fmt"
	"github.com/gorilla/sessions"
	"html"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", hello)
	fmt.Println("listening...")

	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "hello, heroku")
	//adding a reference to herokugoauth so it doesn't complain we are not using it
	fmt.Fprintln(res, "Authentication URL: "+herokugoauth.Endpoint.AuthURL)
	fmt.Fprintln(res, "Token URL: "+herokugoauth.Endpoint.TokenURL)
}
