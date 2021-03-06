package main

import (
    "fmt"
    "net/http"
)

type String string

type Struct struct {
    Greeting string
    Punct string
    Who string
}

func (s String) ServeHTTP(
    w http.ResponseWriter,
    r *http.Request) {
        fmt.Fprintf(w,string(s))
}

func (str Struct) ServeHTTP(
    w http.ResponseWriter,
    r *http.Request) {
        temp := str.Greeting + str.Punct + str.Who
        fmt.Fprintf(w,temp)
}

func main() {
    http.Handle("/string", String("I'm a frayed knot."))
    http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
    http.ListenAndServe("localhost:4000",nil)
}
