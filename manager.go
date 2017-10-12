package go_login

import (
	"net/http"
	"log"
)

type LoginManager struct {
	UserMap map[string] *UserMixin
	config *config
}

func NewLoginManager(config *config)*LoginManager{
	return &LoginManager{
		UserMap:make(map[string] *UserMixin),
		config:config,
	}
}

func (manager *LoginManager)Auth(request *http.Request) (*UserMixin, bool){
	cookies := request.Cookies()
	sessionId, index := GetSessionId(cookies)
	if index == -1{
		return nil,false
	}
	user := manager.UserMap[sessionId]
	if user == nil{
		return nil, false
	}

	if user.isLogin == false{
		return  user,false
	}

	token, _:= GetToken(cookies)
	if token != user.token{
		return nil,false
	}

	return  user,true
}

func (manager *LoginManager)Current(request *http.Request) (*UserMixin, error){

	cookies := request.Cookies()
	sessionId, index := GetSessionId(cookies)
	if index == -1{
		return nil,nil
	}

	user := manager.UserMap[sessionId]

	return  user,nil
}

func (manager *LoginManager) Login(user *UserMixin,w *http.ResponseWriter)  {
	if user.isLogin == true{
		log.Println("User: ",user.identity," Already Login")
		return
	}
	user.isLogin = true
	user.identity = GenSessionId(manager.config.secret)
	user.token = GenToken(user.identity)
	http.SetCookie(*w,&http.Cookie{Name:"session_id",Value:user.identity})
	http.SetCookie(*w,&http.Cookie{Name:"token",Value:user.token})
	manager.UserMap[user.identity] = user
	log.Println("User: ",user.identity," Login Success")
}

func (manager *LoginManager) Logout(user *UserMixin, r *http.Request)  {
	user.isLogin = false
	sessionId, _ := GetSessionId(r.Cookies())
	delete(manager.UserMap, sessionId)
	log.Println("User: ",user.identity," Logout")
}

