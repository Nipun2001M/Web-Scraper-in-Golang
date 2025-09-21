package main

import (
	"net/http"
)

func handlerRediness(w http.ResponseWriter,r *http.Request){
	resondWithJson(w,200,struct{}{})

}