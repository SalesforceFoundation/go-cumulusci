##### Heroku OAuth HTTP handler

Lovingly based on <https://github.com/kr/githubauth>

See <http://godoc.org/github.com/heroku/herokugoauth> for documentation.

###### Example

```shell
# Set the random key
$ heroku config:set KEY=$(openssl rand -hex 32)

# install the oauth client
$ heroku plugins:install https://github.com/heroku/heroku-oauth
$ heroku clients:create goauthtest https://goauthtestapp.herokuapp.com/auth/heroku/callback
$ heroku config:set OAUTH_CLIENT_ID=<id> OAUTH_CLIENT_SECRET=<secret>
```
then:

```go
h := &herokugoauth.Handler{
	RequireDomain:   "heroku.com",
	// e.g. faba0c08be7474a785b272c4f4154c998c0943b51e662637be11b1a0ecda43b3
	Key:         os.Getenv("KEY"),
	ClientID:     os.Getenv("OAUTH_CLIENT_ID"),
	ClientSecret: os.Getenv("OAUTH_CLIENT_SECRET"),
}
http.ListenAndServe(":8080", h)
```
