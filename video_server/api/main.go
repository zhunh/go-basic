package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type middleWarehandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandle(r *httprouter.Router) http.Handler {
	m := middleWarehandler{}
	m.r = r
	return m
}

func (m middleWarehandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//check session
	validateUserSession(r)
	m.r.ServeHTTP(w, r)
}

func RegisterHandler() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", CreateUser)
	router.POST("/user/:user_name", Login)
	return router
}

func main() {
	r := RegisterHandler()
	mh := NewMiddleWareHandle(r)
	http.ListenAndServe(":8000", mh)
}
