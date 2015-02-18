This application performs Oauth from Heroku to Salesforce.

Sample application running on Heroku: [https://go-cumulusci.herokuapp.com/](https://go-cumulusci.herokuapp.com/)

Oauth config steps:

 1. Create a Connected App in any production of developer org you own (Setup > Apps > Connected App)
 2. Enable Oauth for it
 3. Give it "Full" and "Perform requests on your behalf at any time" permissions
 4. Enter a callback, like https://MYAPPNAME.herokuapp.com/auth/heroku/callback, where MYAPPNAME is the name of your app running on Heroku
 5. Save it
 6. Enter the provided ID and Secret in your app configuration in Heroku, as below (or deploy the app and enter them through the web UI)

 
 If you want to be able to run it locally (very useful for development):

 7. Create another Connected App, with callback http://localhost:5000/auth/heroku/callback
 8. Create a .env file in the root of your project, and put the OAUTH_CLIENT_ID, OAUTH_CLIENT_SECRET, REDIRECT_URL, and KEY (any random key) as key/value pairs there (add this file to .gitignore so you don't commit it by mistake)

To deploy to Heroku:

```
$ git clone https://github.com/SalesforceFoundation/go-cumulusci
$ cd go-cumulusci/
$ heroku create MYAPPNAME -b https://github.com/kr/heroku-buildpack-go.git
$ heroku config:add OAUTH_CLIENT_ID= #Salesforce-provided ID
$ heroku config:add OAUTH_CLIENT_SECRET= #Salesforce-provided Secret
$ heroku config:add REDIRECT_URL=https://MYAPPNAME.herokuapp.com/auth/heroku/callback
$ heroku config:add KEY=somerandomkey
$ git push heroku master
```

To run locally:

```
$ go get #installs locally any dependencies defined in the import declarations, links and compiles the project
$ foreman start #starts foreman (https://devcenter.heroku.com/articles/procfile#developing-locally-with-foreman)
```
Then visit http://localhost:5000

To develop you'll also need to know these commands:

```					 	
$ go get github.com/kr/godep 	#installs godeps
$ godep save 					#adds dependencies to the Godeps folder (they need to have been installed locally)
$ godep go install 				#builds the project using the saved dependencies in Godeps
```

Note: currently, when you run godeps on mac, it adds darwin/386 after the go version in Godeps.json. You need to remove that, or the push to Heroku will fail. (https://github.com/tools/godep/issues/181.)