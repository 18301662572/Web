package main

import (
	"log"
	"net/http"
	"os"
	"time"
)
//优雅的middeware

type middleware func(http.Handler) http.Handler

type Router struct {
	middlewareChain [] func(http.Handler) http.Handler
	mux map[string] http.Handler
}

func NewRouter() *Router{
	return &Router{}
}

func (r *Router) Use(m middleware) {
	r.middlewareChain = append(r.middlewareChain, m)
}

func (r *Router) Add(route string, h http.Handler) {
	var mergedHandler = h
	for i := len(r.middlewareChain) - 1; i >= 0; i-- {
		mergedHandler = r.middlewareChain[i](mergedHandler)
	}
	r.mux[route] = mergedHandler
}

var logger=log.New(os.Stdout,"",0)

//获取耗时的中间件
func timeout(next http.Handler) http.Handler {

	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()

		// next handler 实现Handler interface的方法
		next.ServeHTTP(wr, r)

		timeElapsed := time.Since(timeStart)
		logger.Println(timeElapsed)
	})
}


func helloHandler(h http.Handler)( http.Handler ) {

	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()
		wr.Write([]byte("hello"))
		h.ServeHTTP(wr, r)
		timeElapsed := time.Since(timeStart)
		logger.Println(timeElapsed)
	})
}

func main(){
	r := NewRouter()
	r.Use(timeout())
	r.Add("/", helloHandler)

}