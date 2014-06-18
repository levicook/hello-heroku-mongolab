hello-heroku-mongolab
=====================

```bash
$ heroku create -b https://github.com/kr/heroku-buildpack-go.git
Creating immense-forest-2514... done, stack is cedar
BUILDPACK_URL=https://github.com/kr/heroku-buildpack-go.git
http://immense-forest-2514.herokuapp.com/ | git@heroku.com:immense-forest-2514.git
Git remote heroku added

$ heroku addons:add mongolab
Adding mongolab on immense-forest-2514... done, v4 (free)
Welcome to MongoLab.  Your new subscription is being created and will be available shortly.  Please consult the MongoLab Add-on Admin UI to check on its progress.
Use `heroku addons:docs mongolab` to view documentation.

$ git push heroku master
Fetching repository, done.
Counting objects: 7, done.
Delta compression using up to 8 threads.
Compressing objects: 100% (4/4), done.
Writing objects: 100% (4/4), 435 bytes | 0 bytes/s, done.
Total 4 (delta 3), reused 0 (delta 0)

-----> Fetching custom git buildpack... done
-----> Go app detected
-----> Using go1.2
-----> Running: godep go install -tags heroku ./...
-----> Discovering process types
       Procfile declares types -> web

-----> Compressing... done, 2.0MB
-----> Launching... done, v7
       http://immense-forest-2514.herokuapp.com/ deployed to Heroku

$ curl -i http://immense-forest-2514.herokuapp.com/
HTTP/1.1 200 OK
Content-Type: text/plain; charset=utf-8
Date: Wed, 18 Jun 2014 18:06:52 GMT
Content-Length: 28
Connection: keep-alive

ds033317.mongolab.com:33317
```
