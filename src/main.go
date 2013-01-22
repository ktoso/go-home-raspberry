package main

import (
        "fmt"
        "http"
)

func main(){
        fmt.Printf("Starting http Server ... ")
        http.Handle("/", http.HandlerFunc(sayHello))
        err := http.ListenAndServe("0.0.0.0:8080", nil)
        if err != nil {
                fmt.Printf("ListenAndServe Error :" + err.String())
        }

func sayHello(c *http.Conn, req *http.Request) {
        fmt.Printf("New Request\n")
        c.Write([]byte("<h1>Go Say's Hello</h1><h2>(Via http)</h2>"))
        c.Flush()
}
