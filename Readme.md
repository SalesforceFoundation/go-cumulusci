This application performs Oauth from Heroku to Salesforce.

First you need to create a Connected app in any production of developer org you own (Setup > Apps > Connected App). Enable Oauth for it, give it "Full" and "Perform requests on your behalf at any time" permissions, and enter a callback like https://MYAPPNAME.herokuapp.com/auth/heroku/callback. Save it, and enter the provided ID and Secret in your app configuration in Heroku, as below. (Or deploy the app and enter them through the web UI.)

To deploy to Heroku:

```
$ git clone https://github.com/SalesforceFoundation/go-cumulusci
$ cd go-cumulusci/
$ heroku create MYAPPNAME -b https://github.com/kr/heroku-buildpack-go.git
$ heroku config:add OAUTH_CLIENT_ID= #Salesforce-provided ID
$ heroku config:add OAUTH_CLIENT_SECRET= #Salesforce-provided Secret
$ heroku config:add REDIRECT_URL=https://MYAPPNAME.herokuapp.com/auth/heroku/callback
$ git push heroku master
```

To develop you'll also need to know these commands:

```
$ go get 					 	#installs locally any dependencies defined in the import declarations, links and compiles the project
$ go get github.com/kr/godep 	#installs godeps
$ godep save 					#adds dependencies to the Godeps folder (they need to have been installed locally)
```

Note: currently, when you run godeps on mac, it adds darwin/386 after the go version in Godeps.json. You need to remove that, or the push to Heroku will fail. (https://github.com/tools/godep/issues/181.)