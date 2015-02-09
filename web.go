package main

import (
	"fmt"
	"github.com/heroku/herokugoauth"
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

var h := &herokugoauth.Handler {
    RequireDomain:   "heroku.com",
    Key:         os.Getenv("KEY"),
    ClientID:     os.Getenv("OAUTH_CLIENT_ID"),
    ClientSecret: os.Getenv("OAUTH_CLIENT_SECRET"),
}
