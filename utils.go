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

func GetSessionId(r *http.Request) (string,error){
	sessionId, err := r.Cookie("session_id")
	if err != nil{
		return "",err
	}

	return sessionId.Value, err
}

func GetToken(r *http.Request) (string,error){
	token, err := r.Cookie("token")
	if err != nil{
		return "",err
	}

	return token.Value, err
}
