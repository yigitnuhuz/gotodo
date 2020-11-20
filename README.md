
<div align="center">

GoToDo is an open source restfull API for `golang`. 

</div>

<br/>

#### 📃 Summary

I'm creating my first GO application. I am applying new frameworks and libraries according to Restfull API standarts. 
In this project, using local sqlite database for now. In the future, other databases will be add.  
Don't forget check new commits if pull or fork to your machine.


#### ⚙️ Install

```
go get -u github.com/yigitnuhuz/gotodo
```


#### 📦 Packages

- web framework
  - echo (github.com/labstack/echo)
- database
  - sqlite (github.com/mattn/go-sqlite3)
- authentication
  - jwt-go (github.com/dgrijalva/jwt-go)


#### 🏗️ Run

You can run belowing commands in project directory:
```
go run .
```

Project will starts on "localhost:3200" in your machine.

#### 🐳  Docker
If you're not familiar with docker, you can skip this section.

The most important thing is gotodo api runs on `3200 HTTP port`. You must consider this during configure your mapping.

-  With CLI
Build docker image in project root directory
```bash
docker build -t gotodo-api .
```

Run docker image in your local machine
```bash
docker run --name gotodo-api -d -p 80:3200 gotodo-api
```

#### 🔑 Auth
Project has simple JWT authentication. You can get bearer token belowing url and parameters.
http://localhost:80/auth/login (if you dont use docker use: localhost:3200/auth/login )
```
{
    "UserName" : "admin",
    "Password" : "password"
}
```
