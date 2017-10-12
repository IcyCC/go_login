package go_login

import (
	"math/rand"
	"time"
	"net/http"
	"strconv"
	"crypto/md5"
	"fmt"
)

func GenSessionId(identity string) string{

	sessionId := identity
	return sessionId
}

func GenToken(identity string) string{
	rand.Seed(time.Now().UnixNano())
	x := strconv.Itoa(rand.Intn(10000))
	s := md5.Sum([]byte(identity+x))
	res := fmt.Sprintf("%x",s)
	return res
}

func GetSessionId(cookies []*http.Cookie) (sessionId string, index int){
	sessionId = ""
	index = -1
	for i, cookie := range cookies{
		if cookie.Name == "session_id"{
			index = i
			sessionId = cookie.Value
			break
		}
	}
	if index == -1{
		return  "",-1
	}
	return sessionId, index
}

func GetToken(cookies []*http.Cookie) (token string, index int){
	index = -1
	for i, cookie := range cookies{
		if cookie.Name == "token"{
			index = i
			token = cookie.Value
			break
		}
	}
	if index == -1{
		return  "",-1
	}

	return token, index
}
