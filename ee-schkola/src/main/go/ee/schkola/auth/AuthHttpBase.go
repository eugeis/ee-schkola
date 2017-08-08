package auth

import (
    "github.com/gorilla/mux"
    "net/http"
)

type AccountHttpHandler struct {
    Router  *mux.Router
    PathPrefix  string
}

func (o *AccountHttpHandler) Setup() (ret error) {
            
    return
    
}

func (o *AccountHttpHandler) AccountsGet(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}

func (o *AccountHttpHandler) AccountsPost(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}

func (o *AccountHttpHandler) AccountsPut(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}

func (o *AccountHttpHandler) AccountsDelete(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}



