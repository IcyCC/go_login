package go_login

import (
	"net/http"
	"log"
	"strconv"
)

type LoginManager struct {
	UserMap map[string] BaseUser
	config *config
	userNum int
}

func NewLoginManager(config *config)*LoginManager{
	return &LoginManager{
		UserMap:make(map[string] BaseUser),
		config:config,
		userNum:0,
	}
}


func (manager *LoginManager)Auth(request *http.Request) (BaseUser, bool){
	currentUser, ok := manager.Current(request)
	if currentUser==nil&&ok == false{
		return nil,false
	}

	if currentUser.getIsLogin() == false{
		return  currentUser,false
	} // user not login

	token, _:= GetToken(request)
	if token != currentUser.getToken(){
		return nil,false
	} //token auth fail

	return  currentUser,true
}

func (manager *LoginManager)Current(request *http.Request) (BaseUser, bool){


	sessionId, err := GetSessionId(request)
	if err != nil{
		return nil,false
	}

	user := manager.UserMap[sessionId]

	if user == nil{
		return nil,false
	}

	return  user,user.getIsLogin()
}

func (manager *LoginManager) Login(user BaseUser,w *http.ResponseWriter)  {
	if user.getIsLogin() == true{
		log.Println("User: ",user.getIdentity()," Already Login")
		return
	}
	user.setIsLogin(true)
	if _,ok := manager.UserMap[user.getIdentity()];!ok{
		user.setIdentity(GenSessionId(manager.config.secret+strconv.Itoa(manager.userNum)))
	}
	manager.userNum++
	user.setToken(GenToken(user.getIdentity()))
	http.SetCookie(*w,&http.Cookie{Name:"session_id",Value:user.getIdentity()})
	http.SetCookie(*w,&http.Cookie{Name:"token",Value:user.getToken()})
	manager.UserMap[user.getIdentity()] = user
	log.Println("User: ",user.getIdentity()," Login Success")
}

func (manager *LoginManager) Logout(user BaseUser, r *http.Request, w *http.ResponseWriter)  {
	if user == nil || user.getIsLogin() == false {
		return
	}
	user.setIsLogin(false)
	sessionId, _ := GetSessionId(r)
	delete(manager.UserMap, sessionId)
	log.Println("User: ",user.getIdentity()," Logout")

}

