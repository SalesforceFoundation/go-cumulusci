package main

import (
	"fmt"
	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

// Account representation.
type Account struct {
	Email string `json:"email"`
}

var store = sessions.NewCookieStore([]byte(os.Getenv("KEY")))

var conf = &oauth2.Config{
	ClientID:     os.Getenv("OAUTH_CLIENT_ID"),
	ClientSecret: os.Getenv("OAUTH_CLIENT_SECRET"),
	Endpoint: oauth2.Endpoint{
		AuthURL:  "https://login.salesforce.com/services/oauth2/authorize",
		TokenURL: "https://login.salesforce.com/services/oauth2/token",
	},
	RedirectURL: os.Getenv("REDIRECT_URL"),
}

func main() {
	http.HandleFunc("/config", config)
	http.HandleFunc("/", handleAuth)
	http.HandleFunc("/auth/heroku/callback", handleCallback)

	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func config(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Current Configuration:")
	fmt.Fprintln(w, "Authentication URL: "+conf.Endpoint.AuthURL)
	fmt.Fprintln(w, "Token URL: "+conf.Endpoint.TokenURL)
	fmt.Fprintln(w, "Redirect URL: "+conf.RedirectURL)
}

func handleAuth(w http.ResponseWriter, r *http.Request) {
	// Redirect user to consent page to ask for permission
	// for the scopes specified above.
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)

	body := `<a href="` + url + `">Sign in with Salesforce</a>`
	w.Write([]byte(body))
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	// Get the authorization code.
	code := r.FormValue("code")
	/**access_token := r.FormValue("access_token")
	refresh_token := r.FormValue("access_token")
	body := `<p>Access Token: ` + access_token + `</p>`
	body = body + `<p>Refresh Token: ` + refresh_token + `</p>`
	body = body + `<p>Code: ` + code + `</p>`
	w.Write([]byte(body))**/

	//Exchange the received code for a token - this is the line that is failing
	tok, err := conf.Exchange(oauth2.NoContext, code)
	showError(w, err)

	//Client returns an HTTP client using the provided token. The token will auto-refresh as necessary.
	client := conf.Client(oauth2.NoContext, tok)

	//we are hardcoding the server url!!! we need to change this!!!
	url := fmt.Sprintf("https://na15.salesforce.com/services/data/32/query?q=%s", url.QueryEscape("select Id, name from account"))
	resp, err := client.Get(url)
	showError(w, err)

	defer resp.Body.Close()
	bodyArray, err := ioutil.ReadAll(resp.Body)
	showError(w, err)

	w.Write([]byte(bodyArray))
}

func showError(w http.ResponseWriter, err error) {
	//var body string
	if err != nil {
		//body = `<p>There was an error: ` + err.Error() + `</p>`
		//w.Write([]byte(body))
		panic(err)
	}
}
