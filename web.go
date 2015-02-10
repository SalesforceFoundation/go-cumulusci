package main

import (
	"fmt"
	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
	"net/http"
	"os"
)

// Account representation.
type Account struct {
	Email string `json:"email"`
}

var store = sessions.NewCookieStore([]byte(os.Getenv("COOKIE_SECRET")))

var conf = &oauth2.Config{
	ClientID:     os.Getenv("OAUTH_CLIENT_ID"),
	ClientSecret: os.Getenv("OAUTH_CLIENT_SECRET"),
	Endpoint: oauth2.Endpoint{
		AuthURL:  "https://login.salesforce.com/services/oauth2/authorize",
		TokenURL: "https://login.salesforce.com/services/oauth2/token",
	},
	RedirectURL: "https://go-cumulusci.herokuapp.com/auth/heroku/callback",
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/", handleAuth)
	http.HandleFunc("/auth/heroku/callback", handleCallback)

	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello, heroku")
	//adding a reference to herokugoauth so it doesn't complain we are not using it
	fmt.Fprintln(w, "Authentication URL: "+conf.Endpoint.AuthURL)
	fmt.Fprintln(w, "Token URL: "+conf.Endpoint.TokenURL)
}

func handleAuth(w http.ResponseWriter, r *http.Request) {
	// Redirect user to consent page to ask for permission
	// for the scopes specified above.
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	//fmt.Fprintln(res, "Visit the URL for the auth dialog: "+url)

	body := `<a href="` + url + `">Sign in with Salesforce</a>`
	w.Write([]byte(body))

	// Use the authorization code that is pushed to the redirect URL.
	// NewTransportWithCode will do the handshake to retrieve
	// an access token and initiate a Transport that is
	// authorized and authenticated by the retrieved token.
	/**var code string
	if _, err := fmt.Scan(&code); err != nil {
		panic(err)
	}
	tok, err := conf.Exchange(oauth2.NoContext, code)
	if err != nil {
		panic(err)
	}

	client := conf.Client(oauth2.NoContext, tok)
	client.Get("...")**/
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	body := `<p>You have successfully authenticated!`
	w.Write([]byte(body))
}
