package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

/*
	The program test two kinds of url path:
	1. http://localhost:3500/foo   OR  http://localhost:3500/hello
	2. http://localhost:3500/
	3. http://localhost:3500/bar
	4. http://localhost:3500/hi/simon
*/
func main() {
	// Cannot listen to both
	// test0()
	// test1()
	// test2()
	// log.Fatal(http.ListenAndServe(":3500", nil))

	//comment all above to test the following
	router := test3()
	log.Fatal(http.ListenAndServe(":3500", router))

}

//Teat0: using handler
type FooHandler struct{}

func (h *FooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, foo!")
}
func test0() {
	foo := FooHandler{}
	http.Handle("/foo", &foo)
	http.Handle("/hello", &foo)
}

//Test1: using handleFunc for write
func test1() {

	http.HandleFunc("/", HelloHandler)
	fmt.Println("Server started at port 3500")
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, there\n")
}

//Test2: using handleFunc for read and write
func test2() {
	//http.Handle("/foo", fooHandler)

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

}

//Test3: LIsten to http Router httprouter after specyfing router
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func test3() *httprouter.Router {
	router := httprouter.New()
	router.GET("/wow/", Index)
	router.GET("/hi/:name", Hello)
	return router
}
