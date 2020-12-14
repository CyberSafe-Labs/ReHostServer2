package main

import (
    "log"
    "net/http"
    "fmt"
)

func main() {
    fmt.Println("ReHost Server Started on Port 8081")

    http.Handle("/", http.FileServer(http.Dir("./static")))

    log.Fatal(http.ListenAndServeTLS(":8081", "server.crt", "server.key", nil))
}
