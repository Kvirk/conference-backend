
# conference-backend

Back end for a conferencing app

## Running Locally

Make sure you have [Go](http://golang.org/doc/install) version 1.16 or newer and the [Heroku Toolbelt](https://toolbelt.heroku.com/) installed.

```sh
$ git clone https://github.com/heroku/conference-backend.git
$ cd conference-backend
$ go build -o bin/conference-backend -v . # or `go build -o bin/conference-backend.exe -v .` in git bash
$ heroku local
```

Your app should now be running on [localhost:5000](http://localhost:5000/).

## Deploying to Heroku

```sh
$ heroku create
$ git push heroku <branch-to-deploy>
$ heroku open
```