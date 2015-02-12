package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("KEY")))

var conf = &oauth2.Config{
	ClientID:     os.Getenv("OAUTH_CLIENT_ID"),
	ClientSecret: os.Getenv("OAUTH_CLIENT_SECRET"),
	Scopes:       []string{"full", "refresh_token"},
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
	//Redirect user to consent page to ask for permission for the scopes specified above.
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	log.Println("****Auth URL: " + url)

	body := `<a href="` + url + `">Sign in with Salesforce</a>`
	w.Write([]byte(body))
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	//Get the authorization code.
	code := r.FormValue("code")
	log.Println("****Returned code: " + code)

	//Setting up parameters to do the refresh token request
	data := url.Values{}
	data.Set("code", code)
	data.Set("grant_type", "authorization_code")
	data.Set("client_id", conf.ClientID)
	data.Set("client_secret", conf.ClientSecret)
	data.Set("redirect_uri", conf.RedirectURL)

	//Requesting the refresh token
	client := &http.Client{}
	tokenResp, err := client.Post("https://login.salesforce.com/services/oauth2/token", "application/x-www-form-urlencoded",
		bytes.NewBufferString(data.Encode())) // <-- URL-encoded payload
	showError(w, err)
	//displayOnPage(w, tokenResp)

	type RefreshTokenData struct {
		Id            string `json:"id"`
		Issued_at     string `json:"issued_at"`
		Scope         string `json:"scope"`
		Instance_url  string `json:"instance_url"`
		Token_type    string `json:"token_type"`
		Refresh_token string `json:"refresh_token"`
		Id_token      string `json:"id_token"`
		Signature     string `json:"signature"`
		Access_token  string `json:"access_token"`
	}

	var tokenData RefreshTokenData
	decoder := json.NewDecoder(tokenResp.Body)
	if jsonerr := decoder.Decode(&tokenData); jsonerr != nil {
		log.Println("****Failed to decode json")
	} else {
		log.Println("****Refresh token: " + tokenData.Refresh_token)
		url := fmt.Sprintf(tokenData.Instance_url+"/services/data/32/query?q=%s", url.QueryEscape("select Id, name from account"))
		log.Println("****Query url: " + url)
		queryResp, err := client.Get(url)
		showError(w, err)

		displayOnPage(w, queryResp)
	}
}

func displayOnPage(w http.ResponseWriter, resp *http.Response) {
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	showError(w, err)
	w.Write([]byte(respBody))
}

func showError(w http.ResponseWriter, err error) {
	if err != nil {
		panic(err)
	}
}
