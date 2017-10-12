package go_login

type UserMixin struct{
	isLogin bool
	token string
	identity string
}

type BaseUser interface {
	setIsLogin(bool)
	getIsLogin()(bool)

	setToken(token string)
	getToken()(string)

	setIdentity(identity string)
	getIdentity()(string)
}

func (user*UserMixin)setIsLogin(b bool)  {
	user.isLogin = b
}

func (user*UserMixin)getIsLogin()bool{
	return user.isLogin
}

func (user*UserMixin)setToken(token string){
	user.token = token
}

func (user*UserMixin)getToken()string{
	return user.token
}

func (user *UserMixin)setIdentity(identity string){
	user.identity = identity
}

func (user *UserMixin)getIdentity()string{
	return user.identity
}
