# go_login
A go HTTP login package for net/http.

## Usage


### Define a LoginManager

```go

var(
	LoginConfig = go_login.NewConfig("hello")
	LoginManager = go_login.NewLoginManager(LoginConfig)
)

```
### Define user model

mixin go_login.UserMixin

```go

type User struct {
	go_login.UserMixin
	name string
	password string
}

```

### Login and Logout

#### login

```go

func LoginHandle(w http.ResponseWriter, r *http.Request) {
   	if r.Method == "POST" {
   		u, islogin := LoginManager.Current(r)
   		if u != nil{
   			currentUser,_ := u.(*User)
   			if islogin == true {
   				fmt.Fprint(w,"you already login ", currentUser.name)
   				return
   			}
   			if islogin == false{
   				LoginManager.Login(currentUser, &w)
   			}
   			return
   		}
   		username := r.FormValue("username")
   		password := r.FormValue("password")
   		if username == "admin" && password == "123"{
   			user := &User{
   				name:     username,
   				password: password,
   			}
   			LoginManager.Login(user,&w)
   			fmt.Println("user login")
   		}
   		fmt.Fprint(w,"login")
   	}

   }

```
#### logout

```go

func LogoutHandle(w http.ResponseWriter, r *http.Request) {
	user,ok := LoginManager.Current(r)
	if user == nil || ok == false{
		fmt.Fprint(w,"No one logout")
		return
	}
	LoginManager.Logout(user,r, &w)
	fmt.Fprint(w,"logout")
}

```

### Auth and Current User

```go

func TestHandle(w http.ResponseWriter, r *http.Request){
	user,ok := LoginManager.Auth(r)
	if !ok {
		fmt.Fprint(w,"None")
		return
	}
	myUser,_ := user.(*User)
	fmt.Fprint(w,"Hello ",myUser.name)
}

```