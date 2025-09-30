package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"webscraper/internal/database"

	"github.com/google/uuid"
)

func (apiCfg *apiConfig)handlerCreateUser(w http.ResponseWriter,r *http.Request){
	type parameters struct{
		Name string`name`

	}
	params:=parameters{}
	error:=json.NewDecoder(r.Body).Decode(&params)
	if error!=nil{
		respondWithError(w,400,fmt.Sprintf("ERROR Parsing Json:",error))
	}
	user,err:=apiCfg.DB.CreateUser(r.Context(),database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
	})
	if err!=nil{
		respondWithError(w,400,fmt.Sprintf("Couldnt Create User:",err))
		return


	}
	respondWithJson(w,200,user)

}