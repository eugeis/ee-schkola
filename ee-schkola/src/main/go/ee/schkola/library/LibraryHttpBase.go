package library

import (
    "github.com/gorilla/mux"
    "net/http"
)

type BookHttpHandler struct {
    Router  *mux.Router
    PathPrefix  string
}

func (o *BookHttpHandler) Setup() (ret error) {
            
    return
    
}

func (o *BookHttpHandler) BooksGet(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}

func (o *BookHttpHandler) BooksPost(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}

func (o *BookHttpHandler) BooksPut(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}

func (o *BookHttpHandler) BooksDelete(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}



