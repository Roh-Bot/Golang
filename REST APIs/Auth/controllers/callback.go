package controllers

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

func Callback(w http.ResponseWriter, r *http.Request){
	state:=r.FormValue("state")
	code:= r.FormValue("code")
	data,err:=getUserData(state,code)
	if err!=nil{
		log.Fatal("error")
	}
	fmt.Fprint(w,"Data :%s",data)
}

func getUserData(state,code string)([]byte,error){
	if state!=RandomString{
		return nil,errors.New("Invalid user state")
	}
	token,err:=oauthGoogle.Exchange(context.Background(),code)
	if err!=nil{
		return nil,err
	}
	respone,err:=http.Get("https://googleapis.com/oauth2/v2/userinfo?access_token="+token.AccessToken)
	if err!=nil{
		return nil,err
	}
	defer respone.Body.Close()
	data,err:=io.ReadAll(respone.Body)
	if err!=nil{
		return nil,err
	}
	return data,nil
}