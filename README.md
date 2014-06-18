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
```

### GET /
```
BUILDPACK_URL=https://github.com/kr/heroku-buildpack-go.git
DYNO=web.1
PATH=/usr/local/bin:/usr/bin:/bin:/app/bin
PWD=/app
MONGOLAB_URI=mongodb://heroku_app26512694:ceja7ber4iqrsmo5cht73t4ejq@ds033317.mongolab.com:33317/heroku_app26512694
PS1=\[\033[01;34m\]\w\[\033[00m\] \[\033[01;32m\]$ \[\033[00m\]
SHLVL=1
HOME=/app
PORT=12682
_=/app/bin/hello-heroku-mongolab
```
