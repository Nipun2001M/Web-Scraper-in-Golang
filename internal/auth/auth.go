package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIkey(headers http.Header) (string,error){
	val:=headers.Get("Authorization")
	if val==""{
		return "",errors.New("no auth Info Found")

	}
	vals:=strings.Split(val," ")
	if len(vals) !=2{
		return "",errors.New("malformed Auth Header")
	}
	if vals[0] !="ApiKey"{
		return "",errors.New("malformed First Part Of the Header")
	}
	return vals[1],nil

}