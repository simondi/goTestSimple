package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

/*
	The program test two kinds of url path:
	1. http://localhost:3500/foo   OR  http://localhost:3500/hello
	2. http://localhost:3500/
	3. http://localhost:3500/bar
*/
func main() {
	test0()
	test1()
	test2()
	log.Fatal(http.ListenAndServe(":3500", nil))
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
