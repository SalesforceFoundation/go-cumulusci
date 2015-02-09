To deploy to Heroku:

```
$ git clone https://github.com/SalesforceFoundation/go-cumulusci
$ cd go-cumulusci/
$ heroku create go-cumulusci -b https://github.com/kr/heroku-buildpack-go.git
$ git push heroku master
```

To develop you'll also need to know these commands:

```
$ go get github.com/kr/godep #installs godeps
$ go get #installs locally any dependencies defined in the import declarations
$ godep save #adds dependencies to the Godeps folder (they need to have been installed locally)
```

Currently, when you run godep on mac, it adds darwin/386 after the go version in Godeps.json. You need to remove that, or the push to Heroku will fail. (I have reported this issue: https://github.com/tools/godep/issues/181.)