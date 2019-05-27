package main

import (
    "fmt"
    "net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request){
    mt.Fprintf(w, "Hello Kuldeep !!, Its Illluminati here !! Don't be afraid , Everything is gonna be fine")
}

func main() {
    http.HandleFunc("/", helloWorld)
    http.ListenAndServe(":8081", nil)
}
