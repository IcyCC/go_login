package go_login


type UserMixin struct{
	isLogin bool
	token string
	identity string
	gcFlag int
}

type BaseUser interface {
	setIsLogin(bool)
	getIsLogin()(bool)

	setToken(token string)
	getToken()(string)

	setIdentity(identity string)
	getIdentity()(string)

	reduceLife()
	addLife()
	getLife()int
	setLife(int)
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

func (user *UserMixin)reduceLife(){
	user.gcFlag = user.gcFlag - 1
}

func (user *UserMixin)addLife(){
	user.gcFlag = user.gcFlag + 1
}

func (user *UserMixin)getLife()int{
	return user.gcFlag
}

func (user *UserMixin)setLife(life int){
	user.gcFlag = life
}